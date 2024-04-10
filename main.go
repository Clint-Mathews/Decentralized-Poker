package main

import (
	"github.com/Clint-Mathews/Decentralized-Poker/p2p"
	server "github.com/Clint-Mathews/Decentralized-Poker/p2p"
)

func main() {
	cfg := p2p.ServerConfig{
		Version:    "Decentralized Poker V1 Beta",
		ListenAddr: ":4000",
	}
	server := server.NewServer(cfg)

	server.Start()
	// d := deck.New()
	// fmt.Println("Hello Poker game!, Deck:", d)
}
