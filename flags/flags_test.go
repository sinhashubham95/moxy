package flags_test

import (
	"github.com/sinhashubham95/moxy/commons"
	"github.com/sinhashubham95/moxy/flags"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPersistencePath(t *testing.T) {
	assert.Equal(t, commons.PersistenceDefaultValue, flags.PersistencePath())
}

func TestPort(t *testing.T) {
	assert.Equal(t, commons.PortDefaultValue, flags.Port())
}
