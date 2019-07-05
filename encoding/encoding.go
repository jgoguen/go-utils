package encoding

import (
	"bytes"
	"encoding/gob"
)

// Marshal converts the given struct to a byte slice.
func Marshal(o interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	err := encoder.Encode(o)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Unmarshal reads from the byte slice `data` and decodes the struct into `o`.
func Unmarshal(data []byte, o interface{}) error {
	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)

	err := decoder.Decode(o)
	if err != nil {
		return err
	}

	return nil
}
