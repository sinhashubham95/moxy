package models

import (
	"path/filepath"
)

// UnMockRequest is the request body for the unMock endpoint
type UnMockRequest struct {
	Tag    string `json:"tag"`
	Method string `json:"method"`
	Path   string `json:"path"`
}

func (r *UnMockRequest) Clean() {
	if r.Path != "" {
		r.Path = filepath.Clean(r.Path)
	}
}
