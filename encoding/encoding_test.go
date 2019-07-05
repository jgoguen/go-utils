package encoding

import (
	"encoding/hex"
	"reflect"
	"testing"
)

type TestObj struct {
	Name  string
	Value uint64
	Size  int32
	Guard float32
}

var (
	object = TestObj{
		Name:  "Hamlet",
		Value: 128,
		Size:  256,
		Guard: 56.4,
	}
	expectedEncoding = "3bff8103010107546573744f626a01ff8200010401044e616d65010c00010556616c7565010600010453697a6501040001054775617264010800000019ff82010648616d6c657401ff8001fe020001fb4033334c4000"
)

func TestMarshal(t *testing.T) {
	data, err := Marshal(object)
	if err != nil {
		t.Errorf("Failed to marshal object: %+v", err)
		t.FailNow()
	}

	encoding := hex.EncodeToString(data)
	if encoding != expectedEncoding {
		t.Errorf("Expected encoding '%s', got '%s'", expectedEncoding, encoding)
	}
}

func TestUnmarshal(t *testing.T) {
	data, err := Marshal(object)
	if err != nil {
		t.Errorf("Failed to marshal object: %+v", err)
		t.FailNow()
	}

	var o TestObj
	err = Unmarshal(data, &o)
	if err != nil {
		t.Errorf("Failed to unmarshal marshaled bytes: %+v", err)
	}

	if !reflect.DeepEqual(object, o) {
		t.Errorf("Unmarshaled object does not match. Expected %+v got %+v", object, o)
	}
}

func TestUnmarshalStoredBytes(t *testing.T) {
	data, err := hex.DecodeString(expectedEncoding)
	if err != nil {
		t.Errorf("Failed decoding stored bytes: %+v", err)
		t.FailNow()
	}

	var o TestObj
	err = Unmarshal(data, &o)
	if err != nil {
		t.Errorf("Failed to unmarshal stored bytes: %+v", err)
	}

	if !reflect.DeepEqual(object, o) {
		t.Errorf("Unmarshaled object does not match. Expected %+v got %+v", object, o)
	}
}
