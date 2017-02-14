package rest

import (
	"io"
	"net/http"
	"net/url"
)

type Client interface {
	// urlfetch1.Client does not allow query parameters in the URL
	SendGet(url string, form url.Values, callback ResponseHandler)

	SendPost(url string, body io.Reader, callback ResponseHandler)
	// form can be converted into ioReader: `strings.NewReader(form.Encode())`
	PostForm(url string, form url.Values, callback ResponseHandler)
}

type AbstractClient struct {
	Client
}

type ResponseHandler func(resp *http.Response, err error)

//func (client *AbstractClient) SendGet(url string, callback ResponseHandler) {
//	client.sendRequest("GET", url, nil, callback)
//}
//
//func (client *AbstractClient) SendPost(url string, body io.Reader, callback ResponseHandler) {
//	client.sendRequest("POST", url, body, callback)
//}

//func (client Client) sendRequest(method string, url string, body io.Reader, callback ResponseHandler) {
//	panic("this is an abstract class")
//}
