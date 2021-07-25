package controllers_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"github.com/sinhashubham95/moxy/commons"
	"github.com/sinhashubham95/moxy/controllers"
	"github.com/sinhashubham95/moxy/models"
	"github.com/sinhashubham95/moxy/persistence/entities"
)

func getMockBytes(t *testing.T, tag, method, path string, statusCode int, body interface{}) []byte {
	request := models.MockRequest{
		Tag:            tag,
		Method:         method,
		Path:           path,
		ResponseStatus: statusCode,
		ResponseBody:   body,
	}
	bytes, err := commons.EncodeJSON(&request)
	assert.NoError(t, err)
	return bytes
}

func TestHandleMockRequestParseError(t *testing.T) {
	mockDecodeJSONWithError()
	defer unMockDecodeJSON()

	res := setupFastHTTPHandlersAndGetResponse(
		t,
		http.MethodPost,
		commons.MockEndpointPath,
		[]byte("sample"),
		controllers.HandleMock,
	)

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	testResponseBodyAgainstError(t, res.Body, "invalid json request body provided")
}

func TestHandleMockEmptyTag(t *testing.T) {
	res := setupFastHTTPHandlersAndGetResponse(
		t,
		http.MethodPost,
		commons.MockEndpointPath,
		getMockBytes(t, "", "GET", "/naruto", 200, "naruto"),
		controllers.HandleMock,
	)

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	testResponseBodyAgainstError(t, res.Body, "empty tag provided")
}

func TestHandleMockInvalidMethod(t *testing.T) {
	res := setupFastHTTPHandlersAndGetResponse(
		t,
		http.MethodPost,
		commons.MockEndpointPath,
		getMockBytes(t, "1234", "", "/naruto", 200, "naruto"),
		controllers.HandleMock,
	)

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	testResponseBodyAgainstError(t, res.Body, "http method should be one of GET, POST, PUT or DELETE")
}

func TestHandleMockEmptyPath(t *testing.T) {
	res := setupFastHTTPHandlersAndGetResponse(
		t,
		http.MethodPost,
		commons.MockEndpointPath,
		getMockBytes(t, "1234", "GET", "", 200, "naruto"),
		controllers.HandleMock,
	)

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	testResponseBodyAgainstError(t, res.Body, "empty path provided")
}

func TestHandleMockActuatorPath(t *testing.T) {
	res := setupFastHTTPHandlersAndGetResponse(
		t,
		http.MethodPost,
		commons.MockEndpointPath,
		getMockBytes(t, "1234", "GET", "/actuator/info", 200, "naruto"),
		controllers.HandleMock,
	)

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	testResponseBodyAgainstError(t, res.Body, fmt.Sprintf("path cannot start with %s", commons.ActuatorPrefix))
}

func TestHandleMockMoxyPrefix(t *testing.T) {
	res := setupFastHTTPHandlersAndGetResponse(
		t,
		http.MethodPost,
		commons.MockEndpointPath,
		getMockBytes(t, "1234", "GET", "/moxy/info", 200, "naruto"),
		controllers.HandleMock,
	)

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	testResponseBodyAgainstError(t, res.Body, fmt.Sprintf("path cannot start with %s", commons.MoxyPrefix))
}

func TestHandleMockInvalidStatus(t *testing.T) {
	res := setupFastHTTPHandlersAndGetResponse(
		t,
		http.MethodPost,
		commons.MockEndpointPath,
		getMockBytes(t, "1234", "GET", "/naruto", 8, "naruto"),
		controllers.HandleMock,
	)

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	testResponseBodyAgainstError(t, res.Body, "response status code should be in the range 100-599")
}

func TestHandleMockPersistenceSaveError(t *testing.T) {
	mockPersistenceSaveWithError()
	defer unMockPersistenceSave()

	res := setupFastHTTPHandlersAndGetResponse(
		t,
		http.MethodPost,
		commons.MockEndpointPath,
		getMockBytes(t, "1234", "GET", "/naruto", 0, "naruto"),
		controllers.HandleMock,
	)

	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	testResponseBodyAgainstError(t, res.Body, "error")
}

func TestHandleMock(t *testing.T) {
	defer func(t *testing.T) {
		assert.NoError(
			t,
			persistenceDelete(&entities.Mock{
				Tag:    "1234",
				Method: "GET",
				Path:   "/naruto",
			}),
		)
	}(t)
	res := setupFastHTTPHandlersAndGetResponse(
		t,
		http.MethodPost,
		commons.MockEndpointPath,
		getMockBytes(t, "1234", "GET", "/naruto", 0, "naruto"),
		controllers.HandleMock,
	)

	assert.Equal(t, http.StatusOK, res.StatusCode)
}
