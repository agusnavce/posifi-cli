package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Client interfaces that defines how a apicore works
type Client interface {
	Get(endpoint string, headers map[string]string) (*http.Response, error)
	Post(endpoint string, headers map[string]string, body interface{}) (*http.Response, error)
	ParseResponseBody(response *http.Response, out interface{}) error
}

type client struct {
	BaseURL   string
	Client    *http.Client
	AuthToken string
}

// NewHTTPClient - Intantiate the Client
func NewHTTPClient(cfg Config) Client {
	d := time.Duration(5000)
	httpclient := &http.Client{Timeout: d * time.Millisecond}
	return &client{
		BaseURL: cfg.GetHostname(),
		Client:  httpclient,
	}
}

// Get - The Get method of the Client
func (c *client) Get(endpoint string, headers map[string]string) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, endpoint)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for key, val := range headers {
		req.Header.Add(key, val)
	}

	return c.Client.Do(req)
}

// Post - The Post method of the rest Client
func (c *client) Post(endpoint string, headers map[string]string, body interface{}) (*http.Response, error) {

	url := fmt.Sprintf("%s/%s", c.BaseURL, endpoint)

	b, err := json.Marshal(body)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))

	req.Header.Set("Content-Type", "application/json")

	if headers != nil {
		// set headers in request
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}

	if err != nil {
		return nil, err
	}

	return c.Client.Do(req)
}

// ParseResponseBody - Parse the response body from JSON
func (c *client) ParseResponseBody(response *http.Response, out interface{}) error {
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, out)
}
