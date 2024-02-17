package discovery

import (
	"fmt"
	"net/http"
)

type Server struct {
}

func (s *Server) Start() {
	handler := http.ServeMux{}
	handler.HandleFunc("POST /register",
		func(writer http.ResponseWriter, request *http.Request) {

		})
	server := http.Server{
		Addr:    ":9090",
		Handler: &handler,
	}
	go server.ListenAndServe()
	fmt.Printf("Listening on %s\n", server.Addr)
}
