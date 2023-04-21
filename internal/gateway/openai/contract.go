package openai

import "net/http"

type client interface {
	Do(request *http.Request) (*http.Response, error)
}
