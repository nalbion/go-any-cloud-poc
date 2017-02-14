package controllers

import (
	"net/http"
	"fmt"
	"github.com/nalbion/go-any-cloud-poc/lib/rest"
)

func HandleIndex(req rest.InboundRequest) (interface{}, error) {
	fmt.Printf("%s request: %v", req.GetMethod(), req.GetBody())

	return nil, req.SendError("not implemented", http.StatusNotImplemented)
}
