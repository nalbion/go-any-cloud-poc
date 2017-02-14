package rest

import "net/http"

type InboundRequest interface {
	GetMethod() string
	GetBody() (map[string]interface{}, error)
	GetParameter(name string) (string, bool)
	SendError(message string, statusCode int) error
	NewClient(header http.Header) Client
}
