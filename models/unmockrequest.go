package models

import (
	"errors"
	"net/http"
)

// UnMockRequest is the request body for the unMock endpoint
type UnMockRequest struct {
	Tag    string `json:"tag"`
	Method string `json:"method"`
	Path   string `json:"path"`
}

func (r *UnMockRequest) Validate() error {
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
		// empty path - which is not allowed
		return errors.New("empty path provided")
	}
	return nil
}
