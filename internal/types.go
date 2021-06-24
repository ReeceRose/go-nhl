package internal

import "net/http"

type Client struct {
	httpClient *http.Client
}
