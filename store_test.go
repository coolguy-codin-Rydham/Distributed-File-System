package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformfunc(t *testing.T) {
	key := "momsbestpicture"
	pathKey := CASPathTransformFunc(key)
	// fmt.Println(pathname + "Hulla hula re hulle hulle test pass hojeye balle balle")
	expectedOriginal := "6804429f74181a63c50c3d81d733a12f14a353ff"
	expectedPathName := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff/"

	if pathKey.PathName != expectedPathName {
		t.Error(t, "have &s want %s", pathKey.PathName, expectedPathName)
	}
	if pathKey.FileName != expectedOriginal {
		t.Error(t, "have &s want %s", pathKey.FileName, expectedOriginal)
	}

}

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)

	key := "momsSpecials"
	data := []byte("Some jpg bytes")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}
	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "momsSpecials"
	data := []byte("Some jgp bytes")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	// r, err := s.Read(key)
	// if err != nil {
	// 	t.Error(err)
	// }
	// b, _ := io.ReadAll(r)

	// fmt.Println(string(b))

	// if string(b) != string(data) {
	// 	t.Errorf("want %s have %s", data, b)
	// }

	// s.Delete(key)

	fmt.Println("Test Passed")
}
