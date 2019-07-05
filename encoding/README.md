# go-utils/encoding

This provides utility functions relating to encodings. Currently provided
functions are:

- `Marshal(interface{}) ([]byte, error)` /
  `Unmarshal([]byte, interface{}) error`: These functions encapsulate
    converting a struct to a byte slice and back to the original struct. The
    `gob` package is used under the hood.
