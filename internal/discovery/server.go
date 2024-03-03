package discovery

import (
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/mikolaj-t/baby-shard/internal/repository"

	"github.com/rs/zerolog/log"
)

type Server struct {
	Repo repository.KeyValueStore
}

func (s *Server) Start() {
	handler := http.ServeMux{}
	handler.HandleFunc("POST /register",
		func(_ http.ResponseWriter, req *http.Request) {
			b, err := io.ReadAll(req.Body)
			if err != nil {
				return
			}
			count, err := s.Repo.Count(req.Context())
			if err != nil {
				return
			}
			err = s.Repo.SetValue(req.Context(), strconv.Itoa(count), string(b))
			if err != nil {
				return
			}
		})
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
