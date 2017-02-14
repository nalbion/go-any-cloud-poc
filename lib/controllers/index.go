package controllers

import (
	"fmt"
	"github.com/nalbion/go-any-cloud-poc/lib/rest"
	"net/http"
)

func HandleIndex(req rest.InboundRequest) (interface{}, error) {
	body, _ := req.GetBody()
	fmt.Printf("%s request: %v", req.GetMethod(), body)

	return nil, req.SendError("not implemented", http.StatusNotImplemented)
}
