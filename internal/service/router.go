package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mikolaj-t/baby-shard/internal/repository"
)

type Router struct {
	hash   ConsistentHashing
	kv     repository.KeyValueStore
	client http.Client
}

const keyFormat = "server-%d"

func (r Router) RouteRequest(req *http.Request) error {
	key := req.Header.Get("dupa")
	hash, err := r.hash.Hash(key)
	if err != nil {
		return fmt.Errorf("failed to hash key %s: %w", key, err)
	}
	serverURL, err := r.kv.GetValue(req.Context(), fmt.Sprintf(keyFormat, hash))
	if err != nil {
		return fmt.Errorf("failed to get value of a key %s: %w", *serverURL, err)
	}
	if serverURL == nil {
		return fmt.Errorf("server url for key %d is nil", hash)
	}

	forwardedReq, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, *serverURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request to server %s: %w", *serverURL, err)
	}
	resp, err := r.client.Do(forwardedReq)
	if err != nil {
		return fmt.Errorf("failed to forward the request to server %s: %w", *serverURL, err)
	}
	defer resp.Body.Close()
	return nil
}
