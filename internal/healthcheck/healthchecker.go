package healthcheck

import "net/http"

type Checker struct {
	client http.Client
}

func (c Checker) Check(url string) bool {
	_, err := c.client.Get(url)
	return err != nil
}
