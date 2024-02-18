package service

import "net/http"

type Router struct {
}

func (r Router) RouteRequest(req *http.Request) error {
	_ = req
	return nil
}
