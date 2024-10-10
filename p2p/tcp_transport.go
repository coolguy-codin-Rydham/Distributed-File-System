package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote nodes over a TCP established connections
type TCPPeer struct {
	// conn is the underlying connection
	conn net.Conn

	// if we dial and retrieve a connection => outbound = true
	// if we accept and retrieve a connection => outbound = false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	shakeHands    HandShakeFunc
	decoder       Decoder

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		shakeHands:    NOPHandshakeFunc,
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil

}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept Error: %s\n", err)
		}
		fmt.Printf("New Incoming Connection %+v\n", conn)
		go t.handleComm(conn)
	}

}

type Temp struct {
}

func (t *TCPTransport) handleComm(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.shakeHands(peer); err != nil {

	}

	// Read Loop
	msg := &Temp{}
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error: %s\n", err)
			continue
		}

	}
	fmt.Printf("New incoming connection %+v\n", peer)
}
