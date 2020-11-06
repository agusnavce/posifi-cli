package pkg

import (
	"net/http"
)

// Client facilitates making HTTP requests
type Client struct {
	http *http.Client
}

// NewHTTPClient - Return a new http client
func NewHTTPClient() *http.Client {
	return &http.Client{}
}

// NewClientFromHTTP takes in an http.Client instance
func NewClientFromHTTP(httpClient *http.Client) *Client {
	client := &Client{http: httpClient}
	return client
}
