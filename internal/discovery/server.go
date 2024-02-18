package discovery

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type Server struct {
}

func (s *Server) Start() {
	handler := http.ServeMux{}
	handler.HandleFunc("POST /register",
		func(_ http.ResponseWriter, _ *http.Request) {})
	server := http.Server{
		Addr:              ":9090",
		Handler:           &handler,
		ReadHeaderTimeout: time.Second,
	}
	go func() {
		_ = server.ListenAndServe()
	}()
	log.Info().Msgf("Listening on %s", server.Addr)
}
