package controllers_test

import (
	"github.com/sinhashubham95/moxy/commons"
	"github.com/sinhashubham95/moxy/controllers"
	"github.com/sinhashubham95/moxy/models"
	"github.com/sinhashubham95/moxy/persistence/entities"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func getUnMockBytes(t *testing.T, tag, method, path string) []byte {
	request := models.UnMockRequest{
		Tag:    tag,
		Method: method,
		Path:   path,
	}
	bytes, err := commons.EncodeJSON(&request)
	assert.NoError(t, err)
	return bytes
}

func TestHandleUnMockRequestParseError(t *testing.T) {
	mockDecodeJSONWithError()
	defer unMockDecodeJSON()

	res := setupFastHTTPHandlersAndGetResponse(
		t,
		http.MethodPost,
		commons.UnMockEndpointPath,
		[]byte("sample"),
		controllers.HandleUnMock,
	)

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	testResponseBodyAgainstError(t, res.Body, "invalid json request body provided")
}

func TestHandleUnMockPersistenceDeleteError(t *testing.T) {
	mockPersistenceDeleteWithError()
	defer unMockPersistenceDelete()

	res := setupFastHTTPHandlersAndGetResponse(
		t,
		http.MethodPost,
		commons.UnMockEndpointPath,
		getUnMockBytes(t, "1234", "GET", "naruto"),
		controllers.HandleUnMock,
	)

	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	testResponseBodyAgainstError(t, res.Body, "error")
}

func TestHandleUnMock(t *testing.T) {
	err := persistenceSave(&entities.Mock{
		Tag:    "1234",
		Method: "GET",
		Path:   "naruto",
		Status: 200,
		Body:   "naruto",
	})
	assert.NoError(t, err)

	res := setupFastHTTPHandlersAndGetResponse(
		t,
		http.MethodPost,
		commons.UnMockEndpointPath,
		getUnMockBytes(t, "1234", "GET", "naruto"),
		controllers.HandleUnMock,
	)

	assert.Equal(t, http.StatusOK, res.StatusCode)
}
