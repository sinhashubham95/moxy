package models_test

import (
	"github.com/sinhashubham95/moxy/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnMockRequestClean(t *testing.T) {
	m := &models.UnMockRequest{
		Tag:    "1234",
		Method: "GET",
		Path:   "//naruto",
	}
	m.Clean()
	assert.Equal(t, "/naruto", m.Path)
}
