package helper

import (
	"bytes"
	"encoding/json"
)

func InterfaceToBytes(obj interface{}) ([]byte, error) {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.Encode(obj)
	return json.Marshal(obj)
}
