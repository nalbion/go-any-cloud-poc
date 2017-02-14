// REST client for use with AppEngine's urlfetch and OAuth 1.0a (one-legged) - as used by Semantics3
package urlfetch1

import (
	"net/http"

	"appengine"
	"appengine/urlfetch"

	"github.com/garyburd/go-oauth/oauth"
	"github.com/nalbion/go-any-cloud-poc/lib/rest"
	"github.com/pkg/errors"
	"io"
	"net/url"
)

type Client struct {
	rest.AbstractClient
	client      *http.Client
	oauthClient oauth.Client
}

func NewClient(ctx appengine.Context, key string, secret string) *Client {
	return &Client{
		client: urlfetch.Client(ctx),
		oauthClient: oauth.Client{
			Credentials: oauth.Credentials{
				Token:  key,
				Secret: secret,
			},
			// Node impl sets requestUrl and accessUrl to null
			//TemporaryCredentialRequestURI: "",
			//ResourceOwnerAuthorizationURI: "",
			//TokenRequestURI: "",
			// ?
			//Header: http.Header{},
			SignatureMethod: oauth.HMACSHA1, // as per Node impl
		},
	}
}

func (c *Client) SendGet(url string, form url.Values, callback rest.ResponseHandler) {
	resp, err := c.oauthClient.Get(c.client, nil, url, form)
	if resp != nil {
		defer resp.Body.Close()
	}

	if callback != nil {
		callback(resp, err)
	}
}

func (c *Client) SendPost(url string, body io.Reader, callback rest.ResponseHandler) {
	callback(nil, errors.New("Call client.PostForm for urlfetch1"))
}

func (c *Client) PostForm(url string, form url.Values, callback rest.ResponseHandler) {
	resp, err := c.oauthClient.Post(c.client, nil, url, form)
	if resp != nil {
		defer resp.Body.Close()
	}

	if callback != nil {
		callback(resp, err)
	}
}
