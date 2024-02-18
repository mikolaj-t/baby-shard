package healthcheck

import (
	"context"
	"fmt"
	"net/http"
)

type Checker struct {
	client http.Client
}

func (c Checker) Check(url string) (bool, error) {
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, url, nil)
	if err != nil {
		return false, fmt.Errorf("failed to create http request: %w", err)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to send http request: %w", err)
	}

	defer resp.Body.Close()
	return true, nil
}
