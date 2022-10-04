package fluentrequest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	expectedBody      = "{\n  \"userId\": 1,\n  \"id\": 1,\n  \"title\": \"delectus aut autem\",\n  \"completed\": false\n}"
	expectedPOSTBody  = "{\n  \"completed\": false,\n  \"id\": 101,\n  \"title\": \"delectus aut autem\",\n  \"userId\": 1\n}"
	expectedPUTBody   = "{\n  \"body\": \"bar\",\n  \"id\": 1,\n  \"title\": \"delectus aut autem\",\n  \"userId\": 1\n}"
	expectedPATCHBody = "{\n  \"userId\": 1,\n  \"id\": 1,\n  \"title\": \"foo\",\n  \"body\": \"bar\"\n}"
)

func TestRequest(t *testing.T) {
	resp, err := FluentRequest().
		Method("GET").
		Url("https://jsonplaceholder.typicode.com/todos/1").
		Run()

	body, _ := ioutil.ReadAll(resp.Body)

	assert.NoError(t, err)
	assert.Equal(t, expectedBody, string(body))
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPOST(t *testing.T) {
	data, _ := json.Marshal(map[string]interface{}{
		"userId":    1,
		"id":        1,
		"title":     "delectus aut autem",
		"completed": false,
	})

	header := http.Header{
		"Content-Type": {"application/json; charset=UTF-8"},
	}

	resp, err := FluentRequest().
		Method(http.MethodPost).
		Url("https://jsonplaceholder.typicode.com/posts").
		Body(bytes.NewBuffer(data)).
		Header(header).
		Run()

	body, _ := ioutil.ReadAll(resp.Body)

	assert.NoError(t, err)
	assert.Equal(t, expectedPOSTBody, string(body))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
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
