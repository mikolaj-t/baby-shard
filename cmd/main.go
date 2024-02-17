package main

import (
	"fmt"
	"github.com/mikolaj-t/baby-shard/internal/discovery"
	"github.com/mikolaj-t/baby-shard/internal/proxy"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := proxy.Server{}
	s.Start()
	ds := discovery.Server{}
	ds.Start()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	fmt.Println("Shutting down the server")
}
