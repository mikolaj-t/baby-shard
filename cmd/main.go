package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mikolaj-t/baby-shard/internal/discovery"
	"github.com/mikolaj-t/baby-shard/internal/proxy"
	"github.com/rs/zerolog/log"
)

func main() {
	s := proxy.Server{}
	s.Start()

	ds := discovery.Server{}
	ds.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Info().Msgf("Shutting down the server")
}
