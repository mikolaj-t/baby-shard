package proxy

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type Server struct {
}

func (s Server) Start() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msgf("Received a request: %+v", r)
		w.WriteHeader(http.StatusOK)
	})
	server := http.Server{
		Handler:           handler,
		Addr:              ":8080",
		ReadHeaderTimeout: time.Second,
	}
	go func() {
		_ = server.ListenAndServe()
	}()
	log.Info().Msgf("Listening on %s", server.Addr)
}
