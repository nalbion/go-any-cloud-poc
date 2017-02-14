// Serves json files named after the request path ("/" replaced by "-")
package stub

import (
	"appengine/file"
	"github.com/nalbion/go-any-cloud-poc/lib/rest"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

type Client struct {
	BasePath string
	t        *testing.T
}

func NewClient(basePath string, t *testing.T) *Client {
	return &Client{
		BasePath: basePath,
		t:        t,
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
	body := strings.NewReader(form.Encode())
	c.sendRequest("POST", url, body, callback)
}

func (c *Client) sendRequest(method string, urlStr string, body io.Reader, callback rest.ResponseHandler) {
	parsedUrl, err := url.Parse(urlStr)

	filePath := parsedUrl.Path
	//filePath = strings.Replace(strings.TrimLeft(filePath, "/"), "/", "-", -1) + ".json"
	filePath = strings.TrimLeft(filePath, "/") + ".json"

	if c.t != nil {
		c.t.Logf("Stubbing %s %s \nwith %s", method, urlStr, c.BasePath+"/"+filePath)
	}

	file, err := os.Open(c.BasePath + "/" + filePath)
	if err != nil {
		callback(nil, err)
	} else {
		defer file.Close()

		response := &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Body:       file,
		}
		callback(response, nil)
	}
}
