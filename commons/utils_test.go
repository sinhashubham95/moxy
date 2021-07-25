package commons_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/sinhashubham95/moxy/commons"
)

func TestEncodeJSON(t *testing.T) {
	data := map[string]string{
		"naruto": "rocks",
	}
	bytes, err := commons.EncodeJSON(data)
	assert.NoError(t, err)
	assert.Contains(t, string(bytes), "naruto")
}

func TestDecodeJSON(t *testing.T) {
	var data map[string]interface{}
	err := commons.DecodeJSON([]byte("{\"naruto\": \"rocks\"}"), &data)
	assert.NoError(t, err)
	assert.Equal(t, "rocks", data["naruto"])
}
