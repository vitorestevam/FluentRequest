package fluentrequest

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	expectedBody = "{\n  \"userId\": 1,\n  \"id\": 1,\n  \"title\": \"delectus aut autem\",\n  \"completed\": false\n}"
)

func TestRequest(t *testing.T) {
	resp, err := FluentRequest().
		Method("GET").
		Url("https://jsonplaceholder.typicode.com/todos/1").
		Run()

	body, _ := ioutil.ReadAll(resp.Body)

	assert.NoError(t, err)
	assert.Equal(t, expectedBody, string(body))
	assert.Equal(t, resp.StatusCode, http.StatusOK)
}
