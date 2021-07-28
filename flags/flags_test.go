package flags_test

import (
	"github.com/sinhashubham95/moxy/commons"
	"github.com/sinhashubham95/moxy/flags"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
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

func TestPortFromEnv(t *testing.T) {
	assert.NoError(t, os.Setenv(strings.ToUpper(commons.Port), "5050"))
	assert.Equal(t, 5050, flags.Port())
}

func TestTLSEnabled(t *testing.T) {
	assert.Equal(t, commons.TLSEnabledDefaultValue, flags.TLSEnabled())
}
