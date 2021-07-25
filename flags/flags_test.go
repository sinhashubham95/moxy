package flags_test

import (
	"github.com/sinhashubham95/moxy/commons"
	"github.com/sinhashubham95/moxy/flags"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCertFilePath(t *testing.T) {
	assert.Equal(t, commons.CertFilePathDefaultValue, flags.CertFilePath())
}

func TestKeyFilePath(t *testing.T) {
	assert.Equal(t, commons.KeyFilePathDefaultValue, flags.KeyFilePath())
}

func TestPersistencePath(t *testing.T) {
	assert.Equal(t, commons.PersistenceDefaultValue, flags.PersistencePath())
}

func TestPort(t *testing.T) {
	assert.Equal(t, commons.PortDefaultValue, flags.Port())
}

func TestTLSEnabled(t *testing.T) {
	assert.Equal(t, commons.TLSEnabledDefaultValue, flags.TLSEnabled())
}
