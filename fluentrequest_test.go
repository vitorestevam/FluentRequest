package fluentrequest

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	name   string
	method string
	url    string
	body   string
	want   ResponseResult
}

type ResponseResult struct {
	responseBody string
	statusCode   int
}

func TestRequest(t *testing.T) {
	tests := []Test{
		{
			name:   "Test GET request",
			method: http.MethodGet,
			url:    "https://jsonplaceholder.typicode.com/todos/1",
			want: ResponseResult{
				statusCode:   http.StatusOK,
				responseBody: `{ "id": 1, "userId": 1, "title": "delectus aut autem", "completed": false }`,
			},
		},
		{
			name:   "Test POST request",
			method: http.MethodPost,
			url:    "https://jsonplaceholder.typicode.com/todos/",
			body:   `{ "id": 201, "userId": 2, "title": "foo", "body": "bar", "completed": true }`,
			want: ResponseResult{
				statusCode:   http.StatusCreated,
				responseBody: `{ "id": 201, "userId": 2, "title": "foo", "body": "bar", "completed": true }`,
			},
		},
		{
			name:   "Test PUT request",
			method: http.MethodPut,
			url:    "https://jsonplaceholder.typicode.com/todos/1",
			body:   `{ "id": 1, "userId": 1, "title": "foo", "body": "bar", "completed": false }`,
			want: ResponseResult{
				statusCode:   http.StatusOK,
				responseBody: `{ "id": 1, "userId": 1, "title": "foo", "body": "bar", "completed": false }`,
			},
		},
		{
			name:   "Test PATCH request",
			method: http.MethodPatch,
			url:    "https://jsonplaceholder.typicode.com/todos/1",
			body:   `{ "id": 1, "userId": 1, "title": "foobar", "body": "", "completed": true }`,
			want: ResponseResult{
				statusCode:   http.StatusOK,
				responseBody: `{ "id": 1, "userId": 1, "title": "foobar", "body": "", "completed": true }`,
			},
		},
		{
			name:   "Test DELETE request",
			method: http.MethodDelete,
			url:    "https://jsonplaceholder.typicode.com/todos/1",
			want: ResponseResult{
				statusCode:   http.StatusOK,
				responseBody: `{}`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createRequest(t, tt.method, tt.url, tt.body, tt.want)
		})
	}
}

func createRequest(t *testing.T, method string, url string, requestBody string, want ResponseResult) {
	bodyBytes := []byte(requestBody)

	header := http.Header{
		"Content-Type": {"application/json; charset=UTF-8"},
	}

	resp, err := FluentRequest().
		Method(method).
		Body(bytes.NewBuffer(bodyBytes)).
		Header(header).
		Url(url).
		Run()

	responseBody, _ := io.ReadAll(resp.Body)

	assert.NoError(t, err)
	assert.JSONEq(t, want.responseBody, string(responseBody))
	assert.Equal(t, want.statusCode, resp.StatusCode)
}
