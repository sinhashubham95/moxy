package controllers_test

import (
	"github.com/sinhashubham95/moxy/commons"
	"github.com/sinhashubham95/moxy/controllers"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHandleBase(t *testing.T) {
	res := setupFastHTTPHandlersAndGetResponse(
		t,
		http.MethodGet,
		commons.BasePath,
		nil,
		controllers.HandleBase,
	)

	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	defer func(t *testing.T, body io.ReadCloser) {
		assert.NoError(t, body.Close())
	}(t, res.Body)
	assert.Equal(t, "Naruto rocks, and so does Moxy!!", string(body))
}
