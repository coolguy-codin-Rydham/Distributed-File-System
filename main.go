package main

import (
	"fmt"
	"log"

	"github.com/coolguy-codin-Rydham/Distributed-File-System/p2p"
)

func main() {

	fmt.Println("Trying")
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		Handshakefunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
	// fmt.Println("Hello from Go")
}
