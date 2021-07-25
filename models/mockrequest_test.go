package models_test

import (
	"fmt"
	"github.com/sinhashubham95/moxy/commons"
	"github.com/sinhashubham95/moxy/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMockRequestClean(t *testing.T) {
	m := &models.MockRequest{
		Tag:            "1234",
		Method:         "GET",
		Path:           "//naruto",
		ResponseStatus: 0,
		ResponseBody:   "naruto",
	}
	m.Clean()
	assert.Equal(t, "/naruto", m.Path)
}

func TestMockRequestDefault(t *testing.T) {
	m := &models.MockRequest{
		Tag:            "1234",
		Method:         "GET",
		Path:           "naruto",
		ResponseStatus: 0,
		ResponseBody:   "naruto",
	}
	m.Default()
	assert.Equal(t, 200, m.ResponseStatus)
}

func TestMockRequestEmptyTag(t *testing.T) {
	m := &models.MockRequest{
		Tag:            "",
		Method:         "GET",
		Path:           "naruto",
		ResponseStatus: 0,
		ResponseBody:   "naruto",
	}
	err := m.Validate()
	assert.Error(t, err)
	assert.Equal(t, "empty tag provided", err.Error())
}

func TestMockRequestInvalidMethod(t *testing.T) {
	m := &models.MockRequest{
		Tag:            "1234",
		Method:         "",
		Path:           "naruto",
		ResponseStatus: 0,
		ResponseBody:   "naruto",
	}
	err := m.Validate()
	assert.Error(t, err)
	assert.Equal(t, "http method should be one of GET, POST, PUT or DELETE", err.Error())
}

func TestMockRequestEmptyPath(t *testing.T) {
	m := &models.MockRequest{
		Tag:          "1234",
		Method:       "GET",
		Path:         "",
		ResponseBody: "naruto",
	}
	err := m.Validate()
	assert.Error(t, err)
	assert.Equal(t, "empty path provided", err.Error())
}

func TestMockRequestActuatorPath(t *testing.T) {
	m := &models.MockRequest{
		Tag:          "1234",
		Method:       "GET",
		Path:         "/actuator/info",
		ResponseBody: "naruto",
	}
	err := m.Validate()
	assert.Error(t, err)
	assert.Equal(t, fmt.Errorf("path cannot start with %s", commons.ActuatorPrefix), err)
}

func TestMockRequestMoxyPath(t *testing.T) {
	m := &models.MockRequest{
		Tag:          "1234",
		Method:       "GET",
		Path:         "/moxy/info",
		ResponseBody: "naruto",
	}
	err := m.Validate()
	assert.Error(t, err)
	assert.Equal(t, fmt.Errorf("path cannot start with %s", commons.MoxyPrefix), err)
}

func TestMockRequestInvalidStatus(t *testing.T) {
	m := &models.MockRequest{
		Tag:            "1234",
		Method:         "GET",
		Path:           "/naruto",
		ResponseStatus: 8,
		ResponseBody:   "naruto",
	}
	err := m.Validate()
	assert.Error(t, err)
	assert.Equal(t, "response status code should be in the range 100-599", err.Error())
}

func TestMockRequestValid(t *testing.T) {
	m := &models.MockRequest{
		Tag:          "1234",
		Method:       "GET",
		Path:         "/naruto",
		ResponseBody: "naruto",
	}
	err := m.Validate()
	assert.NoError(t, err)
}
