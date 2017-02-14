package http

import (
	"github.com/nalbion/go-any-cloud-poc/lib/rest"
	"github.com/nalbion/go-any-cloud-poc/lib/rest/urlfetch"
	"encoding/json"
	"github.com/pkg/errors"
	"appengine"
	"net/http"
	"net/url"
)

type WrappedHandlerFunc func(req rest.InboundRequest) (interface{}, error)

type InboundRequest struct {
	w          http.ResponseWriter
	req        *http.Request
	parameters url.Values
	urlfetch   bool
}

/**
 * Allows http.HandleFunc to be provided with a HandlerFunc which wraps w & r with an InboundRequest
 * Set urlfetch=true for AppEngine
 */
func WrapHandler(urlfetch bool, handler WrappedHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := NewInboundRequest(w, r, urlfetch)
		result, err := handler(req)
		if err == nil {
			req.SendJson(result)
		}
	}
}

func NewInboundRequest(w http.ResponseWriter, req *http.Request, urlfetch bool) *InboundRequest {
	return &InboundRequest{w, req, nil, urlfetch}
}

func (ir *InboundRequest) GetMethod() string {
	return ir.req.Method
}

func (ir *InboundRequest) GetBody() (map[string]interface{}, error) {
	var b map[string]interface{}
	if ir.req.Body == nil {
		return nil, ir.SendError("empty request", http.StatusBadRequest)
	}

	err := json.NewDecoder(ir.req.Body).Decode(&b)
	if err != nil {
		return nil, ir.SendError("could not parse request", http.StatusBadRequest)
	}

	return b, err
}

func (ir *InboundRequest) GetParameter(name string) (value string, ok bool) {
	if ir.parameters == nil {
		ir.parameters = ir.req.URL.Query()
	}
	value = ir.parameters.Get(name)
	return value, (value != "")
}

// example statusCode: http.StatusBadRequest
func (ir *InboundRequest) SendError(message string, statusCode int) error {
	http.Error(ir.w, message, statusCode)
	return errors.New(message)
}

func (ir *InboundRequest) SendJson(data interface{}) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(ir.w, err.Error(), http.StatusInternalServerError)
		return
	}

	ir.w.Header().Set("Content-Type", "application/json")
	ir.w.Write(js)
}

/** header is not used if the InboundRequest was created with `urlfetch=true` */
func (ir *InboundRequest) NewClient(header http.Header) rest.Client {
	if ir.urlfetch {
		ctx := appengine.NewContext(ir.req)
		return urlfetch.NewClient(ctx, header)
	} else {
		return NewClient(header)
	}
}
