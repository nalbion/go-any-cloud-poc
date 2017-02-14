// REST client for use on AWS (and anywhere else outside of Google Cloud)
package http

import (
	"github.com/nalbion/go-any-cloud-poc/lib/rest"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	rest.AbstractClient
	header http.Header
}

func NewClient(header http.Header) *Client {
	return &Client{
		header: header,
	}
}

// form may be nil
func (c *Client) SendGet(url string, form url.Values, callback rest.ResponseHandler) {
	var body io.Reader
	if form != nil {
		body = strings.NewReader(form.Encode())
	} else {
		body = nil
	}
	c.sendRequest("GET", url, body, callback)
}

func (c *Client) SendPost(url string, body io.Reader, callback rest.ResponseHandler) {
	c.sendRequest("POST", url, body, callback)
}

func (c *Client) PostForm(url string, form url.Values, callback rest.ResponseHandler) {
	callback(nil, errors.New("Call client.SendPost for http.Client"))
}

func (c *Client) sendRequest(method string, url string, body io.Reader, callback rest.ResponseHandler) {
	client := http.DefaultClient //http.Client{}
	req, err := http.NewRequest(method, url, body)
	req.Close = true

	req.Header = c.header

	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	if callback != nil {
		callback(resp, err)
	}
}
