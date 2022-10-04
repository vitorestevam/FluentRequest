package fluentrequest

import (
	"io"
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
	body   io.Reader
	header http.Header
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

func (r *fluentRequest) Body(body io.Reader) *fluentRequest {
	r.body = body

	return r
}

func (r *fluentRequest) Header(header map[string][]string) *fluentRequest {
	r.header = header

	return r
}

func (r *fluentRequest) Run() (*http.Response, error) {

	req, _ := http.NewRequest(r.method, r.url, r.body)

	req.Header = r.header

	return r.client.Do(req)
}
