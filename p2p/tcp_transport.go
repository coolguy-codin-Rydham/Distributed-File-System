package p2p

import (
	"fmt"
	"net"
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

func (p *TCPPeer) Close() error {
	return p.conn.Close()
}

type TCPTransportOpts struct {
	ListenAddr    string
	Handshakefunc HandShakeFunc
	Decoder       Decoder
	OnPeer        func(Peer) error
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	rpcch    chan RPC
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcch:            make(chan RPC),
	}
}

// Consume implements the Transport interface, which will readonly channel
// for reading the incoming messages received from another peer in the network
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcch
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)
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
		go t.handleConn(conn)
	}

}

type Temp struct {
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	var err error
	defer func() {
		fmt.Printf("Dropping Peer Connection due to error: %v\n", err)
		conn.Close()
	}()

	// Initialize peer
	peer := NewTCPPeer(conn, true)

	// Perform handshake
	if err = t.Handshakefunc(peer); err != nil {
		// Log the handshake failure
		fmt.Printf("Handshake failed: %v\n", err)
		return
	}

	// Trigger OnPeer callback if defined
	if t.OnPeer != nil {
		if err = t.OnPeer(peer); err != nil {
			// Log the OnPeer failure
			fmt.Printf("OnPeer callback failed: %v\n", err)
			return
		}
	}

	// Read Loop
	rpc := RPC{}
	for {
		// Try to decode the RPC
		err = t.Decoder.Decode(conn, &rpc)
		if err != nil {
			return
		}

		// Set the RPC source to the remote address
		rpc.From = conn.RemoteAddr()
		t.rpcch <- rpc
	}
}
