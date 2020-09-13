package helper

import (
	"bytes"
	"encoding/json"
)

// InterfaceToBytes encodes an object and returns its as a byte array
func InterfaceToBytes(obj interface{}) ([]byte, error) {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.Encode(obj)
	return json.Marshal(obj)
}
