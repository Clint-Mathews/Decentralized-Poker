package main

import (
	"fmt"
	"time"

	"github.com/Clint-Mathews/Decentralized-Poker/deck"
	"github.com/Clint-Mathews/Decentralized-Poker/p2p"
)

func main() {
	cfg := p2p.ServerConfig{
		Version:    "Decentralized Poker V1 Beta",
		ListenAddr: ":4000",
	}
	server := p2p.NewServer(cfg)
	go server.Start()

	time.Sleep(1 * time.Second)

	RemoteServerCfg := p2p.ServerConfig{
		Version:    "Decentralized Poker V1 Beta",
		ListenAddr: ":4001",
	}
	RemoteServer := p2p.NewServer(RemoteServerCfg)
	go RemoteServer.Start()
	if err := RemoteServer.Connect(":4000"); err != nil {
		panic(err)
	}

	d := deck.New()
	fmt.Println("Hello Poker game!, Deck:", d)
}
