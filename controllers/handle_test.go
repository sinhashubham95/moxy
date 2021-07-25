package controllers_test

import (
	"fmt"
	"github.com/sinhashubham95/moxy/commons"
	"github.com/sinhashubham95/moxy/controllers"
	"github.com/sinhashubham95/moxy/persistence/entities"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"net/http"
	"testing"
)

func setupFastHTTPHandlerForHandleAndGetResponse(t *testing.T, m *entities.Mock, url string) *http.Response {
	port := getRandomPortNumber()

	go func(port int, handler fasthttp.RequestHandler) {
		assert.NoError(
			t,
			fasthttp.ListenAndServe(
				fmt.Sprintf(":%d", port),
				func(ctx *fasthttp.RequestCtx) {
					handler(ctx)
				},
			),
		)
	}(port, controllers.Handle)

	request, err := http.NewRequest(m.Method, fmt.Sprintf("http://localhost:%d%s", port, m.Path), nil)
	assert.NoError(t, err)
	request.Header.Set(commons.TagHeader, m.Tag)
	request.Header.Set(commons.URLHeader, url)

	response, err := http.DefaultClient.Do(request)
	assert.NoError(t, err)

	return response
}

func TestHandleWithMockStringBody(t *testing.T) {
	m := &entities.Mock{
		Tag:    "1234",
		Method: http.MethodGet,
		Path:   "/naruto",
		Status: http.StatusOK,
		Body:   "naruto",
	}
	assert.NoError(t, persistenceSave(m))
	defer func(t *testing.T) {
		assert.NoError(t, persistenceDelete(m))
	}(t)

	res := setupFastHTTPHandlerForHandleAndGetResponse(t, m, "")
	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, m.Body, string(body))
}

func TestHandleWithMockJSONBody(t *testing.T) {
	m := &entities.Mock{
		Tag:    "1234",
		Method: http.MethodGet,
		Path:   "/naruto",
		Status: http.StatusOK,
		Body: map[string]interface{}{
			"naruto": "rocks",
		},
	}
	assert.NoError(t, persistenceSave(m))
	defer func(t *testing.T) {
		assert.NoError(t, persistenceDelete(m))
	}(t)

	res := setupFastHTTPHandlerForHandleAndGetResponse(t, m, "")
	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, "{\"naruto\":\"rocks\"}\n", string(body))
}

func TestHandleWithMockJSONBodyEncodeFail(t *testing.T) {
	mockEncodeJSONWithError()
	defer unMockEncodeJSON()

	m := &entities.Mock{
		Tag:    "1234",
		Method: http.MethodGet,
		Path:   "/naruto",
		Status: http.StatusOK,
		Body: map[string]interface{}{
			"naruto": "rocks",
		},
	}
	assert.NoError(t, persistenceSave(m))
	defer func(t *testing.T) {
		assert.NoError(t, persistenceDelete(m))
	}(t)

	res := setupFastHTTPHandlerForHandleAndGetResponse(t, m, "")
	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, "", string(body))
}

func TestHandleWithoutMock(t *testing.T) {
	m := &entities.Mock{
		Tag:    "1234",
		Method: http.MethodGet,
		Path:   "/imghp",
	}

	res := setupFastHTTPHandlerForHandleAndGetResponse(t, m, "https://google.co.in")
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestHandleWithoutMockError(t *testing.T) {
	m := &entities.Mock{
		Tag:    "1234",
		Method: http.MethodGet,
		Path:   "/imghp",
	}

	res := setupFastHTTPHandlerForHandleAndGetResponse(t, m, "https://naruto-rocks-test.com")
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)

	testResponseBodyAgainstError(t, res.Body, "lookup naruto-rocks-test.com: no such host")
}
