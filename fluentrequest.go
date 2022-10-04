package fluentrequest

import (
	"net/http"
)

func FluentRequest() *fluentRequest {
	return &fluentRequest{
	client: http.Client{},
	}
}

type fluentRequest struct {
	url    string
	method string
	client http.Client
}

func (r *fluentRequest) Url(url string) *fluentRequest {
	r.url = url

	return r
}

func (r *fluentRequest) Method(method string) *fluentRequest {
	r.method = method

	return r
}

func (r *fluentRequest) Run() (*http.Response, error) {

	req, _ := http.NewRequest(r.method, r.url, nil)
	return r.client.Do(req)
}
