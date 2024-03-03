package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
)

const argCount = 2

func main() {
	args := os.Args
	if len(args) != argCount {
		log.Fatal().Msg("only argument must be the port")
	}
	port, err := strconv.Atoi(args[1])
	if err != nil || port <= 0 {
		log.Fatal().Msg("port must be a number > 0")
	}
	count := 5
	wg := sync.WaitGroup{}
	wg.Add(count)
	for i := range 5 {
		go func() {
			server(port + i)
			wg.Done()
		}()
	}
	wg.Wait()
}

func server(port int) {
	client := http.Client{}
	addr := fmt.Sprintf("http://localhost:%d", port)
	r := bytes.NewReader([]byte(addr))
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "http://localhost:9090/register", r)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()

	log.Info().Msgf("Starting ping server on port %d", port)
	handler := &http.ServeMux{}
	handler.HandleFunc("/", func(writer http.ResponseWriter, _ *http.Request) {
		_, _ = writer.Write([]byte(fmt.Sprintf("pong %d", port)))
	})
	server := http.Server{
		Handler:           handler,
		Addr:              fmt.Sprintf(":%d", port),
		ReadHeaderTimeout: time.Second,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal().Msgf("%v", err)
	}
}
