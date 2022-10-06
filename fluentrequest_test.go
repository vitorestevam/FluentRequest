package fluentrequest

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	tests := []struct {
		name   string
		method string
		url    string
		header http.Header
		body   string
		want   struct {
			responseBody string
			statusCode   int
		}
	}{
		{
			name:   "Test GET request",
			method: http.MethodGet,
			url:    "https://jsonplaceholder.typicode.com/todos/1",
			header: http.Header{
				"Content-Type": {"application/json; charset=UTF-8"},
			},
			want: struct {
				responseBody string
				statusCode   int
			}{
				responseBody: `{ "id": 1, "userId": 1, "title": "delectus aut autem", "completed": false }`,
				statusCode:   http.StatusOK,
			},
		},
		{
			name:   "Test POST request",
			method: http.MethodPost,
			url:    "https://jsonplaceholder.typicode.com/todos/",
			header: http.Header{
				"Content-Type": {"application/json; charset=UTF-8"},
			},
			body: `{ "id": 201, "userId": 2, "title": "foo", "body": "bar", "completed": true }`,
			want: struct {
				responseBody string
				statusCode   int
			}{
				statusCode:   http.StatusCreated,
				responseBody: `{ "id": 201, "userId": 2, "title": "foo", "body": "bar", "completed": true }`,
			},
		},
		{
			name:   "Test PUT request",
			method: http.MethodPut,
			url:    "https://jsonplaceholder.typicode.com/todos/1",
			header: http.Header{
				"Content-Type": {"application/json; charset=UTF-8"},
			},
			body: `{ "id": 1, "userId": 1, "title": "foo", "body": "bar", "completed": false }`,
			want: struct {
				responseBody string
				statusCode   int
			}{
				statusCode:   http.StatusOK,
				responseBody: `{ "id": 1, "userId": 1, "title": "foo", "body": "bar", "completed": false }`,
			},
		},
		{
			name:   "Test PATCH request",
			method: http.MethodPatch,
			url:    "https://jsonplaceholder.typicode.com/todos/1",
			header: http.Header{
				"Content-Type": {"application/json; charset=UTF-8"},
			},
			body: `{ "id": 1, "userId": 1, "title": "foobar", "body": "", "completed": true }`,
			want: struct {
				responseBody string
				statusCode   int
			}{
				statusCode:   http.StatusOK,
				responseBody: `{ "id": 1, "userId": 1, "title": "foobar", "body": "", "completed": true }`,
			},
		},
		{
			name:   "Test DELETE request",
			method: http.MethodDelete,
			url:    "https://jsonplaceholder.typicode.com/todos/1",
			header: http.Header{
				"Content-Type": {"application/json; charset=UTF-8"},
			},
			want: struct {
				responseBody string
				statusCode   int
			}{
				statusCode:   http.StatusOK,
				responseBody: `{}`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := FluentRequest().
				Method(tt.method).
				Body(strings.NewReader(tt.body)).
				Header(tt.header).
				Url(tt.url).
				Run()

			responseBody, _ := io.ReadAll(resp.Body)

			assert.NoError(t, err)
			assert.Equal(t, tt.want.statusCode, resp.StatusCode)
			assert.JSONEq(t, tt.want.responseBody, string(responseBody))
		})
	}
}
