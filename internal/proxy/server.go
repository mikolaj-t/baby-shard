package proxy

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/mikolaj-t/baby-shard/internal/repository"

	"github.com/rs/zerolog/log"
)

type Server struct {
	Repo   repository.KeyValueStore
	client http.Client
}

func (s Server) Start() {
	handler := http.NewServeMux()
	s.client = http.Client{}
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("n")
		val, err := s.Repo.GetValue(r.Context(), key)
		if val == nil {
			log.Error().Msgf("val was nil for key %s", key)
		}

		if err == nil && val != nil {
			req, errr := http.NewRequestWithContext(context.Background(), r.Method, *val, nil)
			if errr != nil {
				panic(errr)
			}
			log.Info().Msgf("forwarded req to %s", *val)
			resp, errr := s.client.Do(req)
			if errr != nil {
				panic(errr)
			}
			defer resp.Body.Close()
			b, errr := io.ReadAll(resp.Body)
			if errr != nil {
				return
			}
			_, errr = w.Write(b)
			if errr != nil {
				return
			}
		}
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
