package p2p

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"sync"
)

type Peer struct {
	conn net.Conn
}

func (p *Peer) Send(b []byte) error {
	_, err := p.conn.Write(b)
	return err
}

type ServerConfig struct {
	ListenAddr string
	Version    string
}

type Message struct {
	Payload io.Reader
	From    net.Addr
}

type Server struct {
	ServerConfig

	listener net.Listener
	handler  Handler
	mu       sync.Mutex
	peers    map[net.Addr]*Peer
	addPeer  chan *Peer
	delPeer  chan *Peer
	msgCh    chan *Message
}

func NewServer(cfg ServerConfig) *Server {
	return &Server{
		ServerConfig: cfg,
		peers:        make(map[net.Addr]*Peer),
		addPeer:      make(chan *Peer),
		delPeer:      make(chan *Peer),
		handler:      &DefaultHandler{},
		msgCh:        make(chan *Message),
	}
}

func (s *Server) Start() {
	go s.loop()

	if err := s.listen(); err != nil {
		panic(err)
	}

	fmt.Printf("Game server running in port: %s \n", s.ListenAddr)

	s.acceptConnLoop()
}

func (s *Server) Connect(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	peer := &Peer{conn: conn}
	s.addPeer <- peer

	return peer.Send([]byte(s.Version + "\n"))
}

// TODO: Right now we have redundent code in registering new peers to the netwrok
// Maybe construct a new peer and handshake protocol after registering a plain connection!
func (s *Server) acceptConnLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			panic(err)
		}
		peer := &Peer{conn: conn}
		s.addPeer <- peer

		peer.Send([]byte(s.Version + "\n"))

		go s.handleConn(peer)
	}
}

func (s *Server) handleConn(peer *Peer) {
	buf := make([]byte, 1024)
	for {
		n, err := peer.conn.Read(buf)
		if err != nil {
			break
		}

		s.msgCh <- &Message{
			From:    peer.conn.RemoteAddr(),
			Payload: bytes.NewReader(buf[:n]),
		}
		fmt.Println(string(buf[:n]))
	}

	s.delPeer <- peer
}

func (s *Server) listen() error {
	ln, err := net.Listen("tcp", s.ListenAddr)

	if err != nil {
		return err
	}
	s.listener = ln

	return nil
}

func (s *Server) loop() {
	for {
		select {
		case peer := <-s.delPeer:
			delete(s.peers, peer.conn.RemoteAddr())
			fmt.Printf("Player Disconnected %s \n", peer.conn.RemoteAddr())

		case peer := <-s.addPeer:
			s.peers[peer.conn.RemoteAddr()] = peer
			fmt.Printf("New Player Connected %s \n", peer.conn.RemoteAddr())

		case msg := <-s.msgCh:
			if err := s.handler.HandleMessage(msg); err != nil {
				panic(err)
			}
		}
	}
}
