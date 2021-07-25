package models

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/sinhashubham95/moxy/commons"
)

// MockRequest is the request body for the mock endpoint
type MockRequest struct {
	Tag            string      `json:"tag"`
	Method         string      `json:"method"`
	Path           string      `json:"path"`
	ResponseStatus int         `json:"responseStatus"`
	ResponseBody   interface{} `json:"responseBody"`
}

// Clean is used to clean the request path
func (r *MockRequest) Clean() {
	if r.Path != "" {
		r.Path = filepath.Clean(r.Path)
	}
}

// Validate is used to validate the request
func (r *MockRequest) Validate() error {
	if r.Tag == "" {
		// empty tag - which is not allowed because it is the context of the mock
		return errors.New("empty tag provided")
	}
	if r.Method != http.MethodGet && r.Method != http.MethodPost && r.Method != http.MethodPut &&
		r.Method != http.MethodDelete {
		// method should be one of GET, POST, PUT or DELETE
		return errors.New("http method should be one of GET, POST, PUT or DELETE")
	}
	if r.Path == "" {
		return errors.New("empty path provided")
	}
	if strings.HasPrefix(r.Path, commons.ActuatorPrefix) {
		return fmt.Errorf("path cannot start with %s", commons.ActuatorPrefix)
	}
	if strings.HasPrefix(r.Path, commons.MoxyPrefix) {
		return fmt.Errorf("path cannot start with %s", commons.MoxyPrefix)
	}
	if r.ResponseStatus != 0 && (r.ResponseStatus < 100 || r.ResponseStatus > 599) {
		return errors.New("response status code should be in the range 100-599")
	}
	return nil
}

// Default is used to set default values for the missing request fields
func (r *MockRequest) Default() {
	if r.ResponseStatus == 0 {
		r.ResponseStatus = http.StatusOK
	}
}
