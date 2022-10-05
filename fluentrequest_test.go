package fluentrequest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	expectedBody      = "{\n  \"userId\": 1,\n  \"id\": 1,\n  \"title\": \"delectus aut autem\",\n  \"completed\": false\n}"
	expectedPOSTBody  = "{\n  \"completed\": false,\n  \"id\": 101,\n  \"title\": \"delectus aut autem\",\n  \"userId\": 1\n}"
	expectedPUTBody   = "{\n  \"body\": \"bar\",\n  \"id\": 1,\n  \"title\": \"delectus aut autem\",\n  \"userId\": 1\n}"
	expectedPATCHBody = "{\n  \"userId\": 1,\n  \"id\": 1,\n  \"title\": \"foo\",\n  \"body\": \"bar\"\n}"
)

type ExpectedResult struct {
	body   string
	status int
}

func TestRequests(t *testing.T) {
	var tests = []struct {
		name           string
		options        map[string]interface{}
		expectedResult ExpectedResult
	}{
		{
			name: "testGet",
			options: map[string]interface{}{
				"Method": "GET",
				"Url":    "https://jsonplaceholder.typicode.com/todos/1",
			},
			expectedResult: ExpectedResult{
				body:   `{"userId": 1,  "id": 1,  "title": "delectus aut autem",  "completed": false}`,
				status: http.StatusOK,
			},
		},
		{
			name: "testPost",
			options: map[string]interface{}{
				"Method": "POST",
				"Url":    "https://jsonplaceholder.typicode.com/posts",
				"Body":   bytes.NewBuffer([]byte(`{"userId": 1,  "id": 1,  "title": "delectus aut autem",  "completed": false}`)),
			},
			expectedResult: ExpectedResult{
				body:   `{"id": 101}`,
				status: http.StatusCreated,
			},
		},
		{
			name: "testPostwithJsonHeader",
			options: map[string]interface{}{
				"Method": "POST",
				"Url":    "https://jsonplaceholder.typicode.com/posts",
				"Body":   bytes.NewBuffer([]byte(`{"userId": 1,  "id": 1,  "title": "delectus aut autem",  "completed": false}`)),
				"Header": http.Header{"Content-Type": {"application/json; charset=UTF-8"}},
			},
			expectedResult: ExpectedResult{
				body:   `{"completed": false,"id": 101,"title": "delectus aut autem","userId": 1}`,
				status: http.StatusCreated,
			},
		},
	}

	for _, e := range tests {
		t.Run(e.name, func(t *testing.T) {
			f := FluentRequest()
			for method, parameter := range e.options {
				meth := reflect.ValueOf(f).MethodByName(method)
				param := reflect.ValueOf(parameter)
				resp := meth.Call([]reflect.Value{param})

				f = resp[0].Interface().(*fluentRequest)
			}

			resp, err := f.Run()

			body, _ := ioutil.ReadAll(resp.Body)

			assert.NoError(t, err)
			assert.Equal(t, e.expectedResult.status, resp.StatusCode)
			assert.JSONEq(t, e.expectedResult.body, string(body))
		})
	}
}

func TestPUT(t *testing.T) {
	data, _ := json.Marshal(map[string]interface{}{
		"id":     1,
		"title":  "delectus aut autem",
		"body":   "bar",
		"userId": 1,
	})

	header := http.Header{
		"Content-Type": {"application/json; charset=UTF-8"},
	}

	resp, err := FluentRequest().
		Method(http.MethodPut).
		Url("https://jsonplaceholder.typicode.com/posts/1").
		Body(bytes.NewBuffer(data)).
		Header(header).
		Run()

	body, _ := ioutil.ReadAll(resp.Body)

	assert.NoError(t, err)
	assert.Equal(t, expectedPUTBody, string(body))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPATCH(t *testing.T) {
	data, _ := json.Marshal(map[string]interface{}{
		"title": "foo",
		"body":  "bar",
	})

	header := http.Header{
		"Content-Type": {"application/json; charset=UTF-8"},
	}

	resp, err := FluentRequest().
		Method(http.MethodPatch).
		Url("https://jsonplaceholder.typicode.com/posts/1").
		Body(bytes.NewBuffer(data)).
		Header(header).
		Run()

	body, _ := ioutil.ReadAll(resp.Body)

	assert.NoError(t, err)
	assert.Equal(t, expectedPATCHBody, string(body))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
