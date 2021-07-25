package controllers_test

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"

	"github.com/sinhashubham95/moxy/commons"
	"github.com/sinhashubham95/moxy/controllers"
	"github.com/sinhashubham95/moxy/persistence"
)

var encodeJSON = commons.EncodeJSON
var decodeJSON = commons.DecodeJSON

var persistenceSave = persistence.Save
var persistenceDelete = persistence.Delete

var portMu sync.Mutex
var port = 1201

func getRandomPortNumber() int {
	portMu.Lock()
	defer portMu.Unlock()
	port += 10
	return port
}

func setupFastHTTPHandlersAndGetResponse(t *testing.T, method, path string, body []byte,
	handler fasthttp.RequestHandler) *http.Response {
	port := getRandomPortNumber()

	go func(port int, method, path string, handler fasthttp.RequestHandler) {
		assert.NoError(
			t,
			fasthttp.ListenAndServe(
				fmt.Sprintf(":%d", port),
				func(ctx *fasthttp.RequestCtx) {
					if string(ctx.Method()) == method && string(ctx.Path()) == path {
						handler(ctx)
						return
					}
					ctx.SetStatusCode(http.StatusNotFound)
				},
			),
		)
	}(port, method, path, handler)

	request, err := http.NewRequest(method, fmt.Sprintf("http://localhost:%d%s", port, path),
		bytes.NewReader(body))
	assert.NoError(t, err)

	response, err := http.DefaultClient.Do(request)
	assert.NoError(t, err)

	return response
}

func testResponseBodyAgainstError(t *testing.T, resBody io.ReadCloser, msg string) {
	body, err := ioutil.ReadAll(resBody)
	assert.NoError(t, err)
	defer func(t *testing.T, body io.ReadCloser) {
		assert.NoError(t, body.Close())
	}(t, resBody)
	assert.Equal(t, msg, string(body))
}

func mockEncodeJSONWithError() {
	encodeJSON = controllers.EncodeJSON
	controllers.EncodeJSON = func(interface{}) ([]byte, error) {
		return nil, errors.New("error")
	}
}

func unMockEncodeJSON() {
	controllers.EncodeJSON = encodeJSON
}

func mockDecodeJSONWithError() {
	decodeJSON = controllers.DecodeJSON
	controllers.DecodeJSON = func([]byte, interface{}) error {
		return errors.New("error")
	}
}

func unMockDecodeJSON() {
	controllers.DecodeJSON = decodeJSON
}

func mockPersistenceSaveWithError() {
	persistenceSave = controllers.PersistenceSave
	controllers.PersistenceSave = func(persistence.Entity) error {
		return errors.New("error")
	}
}

func unMockPersistenceSave() {
	controllers.PersistenceSave = persistenceSave
}

func mockPersistenceDeleteWithError() {
	persistenceDelete = controllers.PersistenceDelete
	controllers.PersistenceDelete = func(persistence.Entity) error {
		return errors.New("error")
	}
}

func unMockPersistenceDelete() {
	controllers.PersistenceDelete = persistenceDelete
}
