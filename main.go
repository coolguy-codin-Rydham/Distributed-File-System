package main

import (
	"fmt"
	"log"

	"github.com/coolguy-codin-Rydham/Distributed-File-System/p2p"
)

func main() {
	fmt.Println("Trying")
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
	// fmt.Println("Hello from Go")
}
