package eawsy

import (
	"github.com/nalbion/go-any-cloud-poc/lib/rest"
	rhttp "github.com/nalbion/go-any-cloud-poc/lib/rest/http"
	"encoding/json"
	"fmt"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
	"github.com/pkg/errors"
	"net/http"
)

// Request wrapper for AWS Lambda https://github.com/eawsy/aws-lambda-go-shim

type InboundRequest struct {
	//body string
	//params map[string]string
	evt *apigatewayproxyevt.Event
	//Error string
	//ResponseStatus int
}

func NewInboundRequest(evt *apigatewayproxyevt.Event) *InboundRequest {
	//return &InboundRequest{evt.Body, evt.params}
	return &InboundRequest{evt: evt}
}

func (ir *InboundRequest) GetMethod() string {
	return ir.evt.HTTPMethod
}

func (ir *InboundRequest) GetBody() (map[string]interface{}, error) {
	var b map[string]interface{}
	if ir.evt.Body == "" {
		return nil, ir.SendError("empty request", http.StatusBadRequest)
	}

	err := json.Unmarshal([]byte(ir.evt.Body), &b)
	return b, err
}

func (ir *InboundRequest) GetParameter(name string) (value string, ok bool) {
	value, ok = ir.evt.QueryStringParameters[name]
	return
}

func (ir *InboundRequest) SendError(message string, statusCode int) error {
	//ir.error = message
	//ir.responseStatus = statusCode
	fmt.Println(message)
	return errors.Errorf("{\"status\": %d, \"message\": \"%s\"}", statusCode, message)
}

/*
func (evt *apigatewayproxyevt.Event) GetBody() (interface{}, error) {
	var b interface{}
	err := json.Unmarshal([]byte(evt.Body), &b)
	return b, err
}
*/
func (ir *InboundRequest) NewClient(header http.Header) rest.Client {
	return InboundRequest(header)
}
