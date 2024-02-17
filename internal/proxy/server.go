package proxy

import (
	"fmt"
	"net/http"
)

type Server struct {
}

func (s Server) Start() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received a request: %+v\n", r)
		w.WriteHeader(http.StatusOK)
	})
	server := http.Server{
		Handler: handler,
		Addr:    ":8080",
	}
	go server.ListenAndServe()
	fmt.Printf("Listening on %s\n", server.Addr)
}
