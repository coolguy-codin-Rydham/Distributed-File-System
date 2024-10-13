package main

import (
	"fmt"
	"log"

	"github.com/coolguy-codin-Rydham/Distributed-File-System/p2p"
)

func OnPeer(peer p2p.Peer) error {
	peer.Close()
	// fmt.Println("Doing Some Logic With the Peer outside of TCP Transport")
	return nil
}

func main() {

	fmt.Println("Trying")
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		Handshakefunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
	// fmt.Println("Hello from Go")
}
