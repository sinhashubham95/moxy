package commons

import (
	"bytes"
	"encoding/json"
)

// EncodeJSON is used to encode any type of data to byte array
func EncodeJSON(v interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	e := json.NewEncoder(&buffer)
	err := e.Encode(v)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// DecodeJSON is used to decode any type of data from byte array
func DecodeJSON(b []byte, v interface{}) error {
	d := json.NewDecoder(bytes.NewReader(b))
	return d.Decode(v)
}
