package main

import (
	"context"
	"log"

	libp2p "github.com/libp2p/go-libp2p"
)

func main() {
	host, err := libp2p.New(context.Background())
	if err != nil {
		log.Fatalf("error starting host: %s", err)
	}
	defer host.Close()
}
