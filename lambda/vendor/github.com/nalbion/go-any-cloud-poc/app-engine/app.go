package app

import (
	"github.com/nalbion/go-any-cloud-poc/lib/controllers"
	rhttp "github.com/nalbion/go-any-cloud-poc/lib/rest/http"
	"net/http"
)

func init() {
	http.HandleFunc("/", rhttp.WrapHandler(true, rhttp.WrappedHandlerFunc(controllers.HandleIndex)))
}
