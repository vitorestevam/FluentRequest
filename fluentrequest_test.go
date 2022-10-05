package fluentrequest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestBody struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	UserId    int    `json:"userId"`
	Completed bool   `json:"completed"`
}

type ResponseResult struct {
	responseBody TestBody
	statusCode   int
}

type Test struct {
	name   string
	method string
	url    string
	body   TestBody
	want   ResponseResult
}

func TestRequest(t *testing.T) {
	tests := []Test{
		{
			name:   "Test GET request",
			method: http.MethodGet,
			url:    "https://jsonplaceholder.typicode.com/todos/1",
			want: ResponseResult{
				statusCode: http.StatusOK,
				responseBody: TestBody{
					Id:        1,
					UserId:    1,
					Title:     "delectus aut autem",
					Completed: false,
				},
			},
		},
		{
			name:   "Test POST request",
			method: http.MethodPost,
			url:    "https://jsonplaceholder.typicode.com/todos/",
			body: TestBody{
				Id:        201,
				UserId:    2,
				Title:     "foo",
				Body:      "bar",
				Completed: true,
			},
			want: ResponseResult{
				statusCode: http.StatusCreated,
				responseBody: TestBody{
					Id:        201,
					UserId:    2,
					Title:     "foo",
					Body:      "bar",
					Completed: true,
				},
			},
		},
		{
			name:   "Test PUT request",
			method: http.MethodPut,
			url:    "https://jsonplaceholder.typicode.com/todos/1",
			body: TestBody{
				Id:        1,
				UserId:    1,
				Title:     "foo",
				Body:      "bar",
				Completed: false,
			},
			want: ResponseResult{
				statusCode: http.StatusOK,
				responseBody: TestBody{
					Id:        1,
					UserId:    1,
					Title:     "foo",
					Body:      "bar",
					Completed: false,
				},
			},
		},
		{
			name:   "Test PATCH request",
			method: http.MethodPatch,
			url:    "https://jsonplaceholder.typicode.com/todos/1",
			body: TestBody{
				Id:        1,
				UserId:    1,
				Title:     "foobar",
				Body:      "",
				Completed: true,
			},
			want: ResponseResult{
				statusCode: http.StatusOK,
				responseBody: TestBody{
					Id:        1,
					UserId:    1,
					Title:     "foobar",
					Body:      "",
					Completed: true,
				},
			},
		},
		{
			name:   "Test DELETE request",
			method: http.MethodDelete,
			url:    "https://jsonplaceholder.typicode.com/todos/1",
			want: ResponseResult{
				statusCode: http.StatusOK,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createRequest(t, tt.method, tt.url, tt.body, tt.want)
		})
	}
}

func createRequest(t *testing.T, method string, url string, requestBody interface{}, want ResponseResult) {
	bodyBytes, _ := json.Marshal(requestBody)

	header := http.Header{
		"Content-Type": {"application/json; charset=UTF-8"},
	}

	resp, err := FluentRequest().
		Method(method).
		Body(bytes.NewBuffer(bodyBytes)).
		Header(header).
		Url(url).
		Run()

	responseBody, _ := ioutil.ReadAll(resp.Body)

	var deserializedBody TestBody

	json.Unmarshal(responseBody, &deserializedBody)

	assert.NoError(t, err)
	assert.Equal(t, want.responseBody, deserializedBody)
	assert.Equal(t, want.statusCode, resp.StatusCode)
}
