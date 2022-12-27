package lib

import "net/http"

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	client HTTPClient
)

func init() {
	client = &http.Client{}
}

func GetHttpClient() HTTPClient {
	return client
}
