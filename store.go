package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
)

const defaultRootFolderName = "ggnetwork"

func CASPathTransformFunc(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])

	blockSize := 5
	sliceLen := len(hashStr) / blockSize

	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := i*blockSize, (i*blockSize)+blockSize
		paths[i] = hashStr[from:to]

	}

	pathnam := strings.Join(paths, "/")
	pathnam = pathnam + "/"

	return PathKey{
		PathName: pathnam,
		FileName: hashStr,
	}

}

type PathTransformFunc func(string) PathKey

type PathKey struct {
	PathName string
	FileName string
}

func (p PathKey) FirstPathName() string {
	paths := strings.Split(p.PathName, "/")
	if len(paths) == 0 {
		return ""
	}
	return paths[0]
}

func (p PathKey) FullPath() string {
	fmt.Println("I am in function")
	return fmt.Sprintf("%s%s", p.PathName, p.FileName)
}

type StoreOpts struct {
	// Root is the folder name of the root, containing all the files/folders of the system
	Root              string
	PathTransformFunc PathTransformFunc
}

var DefaultPathTransformFunc = func(key string) PathKey {
	return PathKey{
		PathName: key,
		FileName: key,
	}
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	if opts.PathTransformFunc == nil {
		opts.PathTransformFunc = DefaultPathTransformFunc
	}
	if len(opts.Root) == 0 {
		opts.Root = defaultRootFolderName
	}
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) Has(key string) bool {
	PathKey := s.PathTransformFunc(key)

	_, err := os.Stat(PathKey.FullPath())

	return err != fs.ErrNotExist

}

func (s *Store) Delete(key string) error {
	pathKey := s.PathTransformFunc(key)

	defer func() {
		log.Printf("Deleted [%s] from disk", pathKey.FileName)
	}()

	// if err := os.RemoveAll(pathKey.FullPath()); err != nil {
	// 	return err
	// }

	return os.RemoveAll(pathKey.FirstPathName())
}

func (s *Store) Read(key string) (io.Reader, error) {
	f, err := s.readStream(key)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, f)
	return buf, err
}

func (s *Store) readStream(key string) (io.ReadCloser, error) {
	pathKey := s.PathTransformFunc(key)

	return os.Open(pathKey.FullPath())

}

func (s *Store) writeStream(key string, r io.Reader) error {

	pathKey := s.PathTransformFunc(key)
	pathNameWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.PathName)

	if err := os.MkdirAll(pathNameWithRoot, os.ModePerm); err != nil {
		return err
	}

	// fullPath := pathKey.FullPath()
	fullPathWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.FullPath())

	f, err := os.Create(fullPathWithRoot)

	if err != nil {
		return err
	}

	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}

	log.Printf("Written (%d) bytes to disk: %s", n, fullPathWithRoot)
	return nil
}
