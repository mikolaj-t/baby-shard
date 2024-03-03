package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mikolaj-t/baby-shard/internal/repository"

	"github.com/mikolaj-t/baby-shard/internal/discovery"
	"github.com/mikolaj-t/baby-shard/internal/proxy"
	"github.com/rs/zerolog/log"
)

func main() {
	repo := &repository.BasicKV{}
	err := repo.Connect()
	if err != nil {
		return
	}

	s := proxy.Server{
		Repo: repo,
	}
	s.Start()

	ds := discovery.Server{
		Repo: repo,
	}
	ds.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Info().Msgf("Shutting down the server")
}
