package entities

import (
	"github.com/sinhashubham95/moxy/commons"
)

// MockKey is the primary key for mock entity
type MockKey struct {
	Tag    string `json:"tag"`
	Method string `json:"method"`
	Path   string `json:"path"`
}

// Mock is the mock entity
type Mock struct {
	Tag    string      `json:"tag"`
	Method string      `json:"method"`
	Path   string      `json:"path"`
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
}

// Name is the mock entity name
func (m *Mock) Name() ([]byte, error) {
	return []byte(commons.MockEntityName), nil
}

// Key is the key of the mock entity
func (m *Mock) Key() ([]byte, error) {
	return commons.EncodeJSON(MockKey{
		Tag:    m.Tag,
		Method: m.Method,
		Path:   m.Path,
	})
}

func (m *Mock) Encode() ([]byte, error) {
	return commons.EncodeJSON(m)
}

func (m *Mock) Decode(b []byte) error {
	return commons.DecodeJSON(b, m)
}
