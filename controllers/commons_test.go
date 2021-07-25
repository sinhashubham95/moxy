package controllers_test

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/sinhashubham95/moxy/commons"
	"github.com/sinhashubham95/moxy/controllers"
	"github.com/sinhashubham95/moxy/persistence"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"net/http"
	"sync"
	"testing"
)

var encodeJSON = commons.EncodeJSON
var decodeJSON = commons.DecodeJSON

var persistenceSave = persistence.Save
var persistenceView = persistence.View
var persistenceDelete = persistence.Delete

var portMu sync.Mutex
var port = 1001

func getRandomPortNumber() int {
	portMu.Lock()
	defer portMu.Unlock()
	port += 10
	return port
}

func setupFastHTTPHandlersAndGetResponse(t *testing.T, method, path string, body []byte,
	handler fasthttp.RequestHandler) *http.Response {
	port := getRandomPortNumber()

	go func(method, path string, handler fasthttp.RequestHandler) {
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
	}(method, path, handler)

	request, err := http.NewRequest(method, fmt.Sprintf("http://localhost:%d%s", port, path),
		bytes.NewReader(body))
	assert.NoError(t, err)

	response, err := http.DefaultClient.Do(request)
	assert.NoError(t, err)

	return response
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

func mockPersistenceViewWithError() {
	persistenceView = controllers.PersistenceView
	controllers.PersistenceView = func(persistence.Entity) error {
		return errors.New("error")
	}
}

func unMockPersistenceView() {
	controllers.PersistenceView = persistenceView
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
