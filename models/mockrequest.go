package models

import (
	"errors"
	"github.com/sinhashubham95/moxy/commons"
	"net/http"
	"path/filepath"
	"strings"
)

// MockRequest is the request body for the mock endpoint
type MockRequest struct {
	Tag            string      `json:"tag"`
	Method         string      `json:"method"`
	Path           string      `json:"path"`
	ResponseStatus int         `json:"responseStatus"`
	ResponseBody   interface{} `json:"responseBody"`
}

func (r *MockRequest) Clean() {
	r.Path = filepath.Clean(r.Path)
}

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
	if r.Path == "" || strings.HasPrefix(r.Path, commons.MoxyPrefix) {
		// empty path - which is not allowed
		return errors.New("empty path provided")
	}
	return nil
}

func (r *MockRequest) Default() {
	if r.ResponseStatus == 0 {
		r.ResponseStatus = http.StatusOK
	}
}
