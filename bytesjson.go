package jabob

import (
	"encoding/json"
)

const emptyJSONAsBunchOfBytes = "null"

var bytesOfNull = []byte(emptyJSONAsBunchOfBytes)

// BytesJSON contain JSON in byte slice.
type BytesJSON struct {
	Bytes []byte
}

// MarshalJSON implements the json.Marshaler interface.
func (v BytesJSON) MarshalJSON() ([]byte, error) {
	if len(v.Bytes) == 0 {
		return bytesOfNull, nil
	}
	return v.Bytes, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (v *BytesJSON) UnmarshalJSON(data []byte) (err error) {
	l := len(data)
	if l == 0 {
		v.Bytes = nil
		return
	}
	v.Bytes = make([]byte, l)
	copy(v.Bytes, data)
	return
}

// MarshalFrom marshal given data into JSON with encoding/json.Marshal function.
func (v *BytesJSON) MarshalFrom(data interface{}) (err error) {
	buf, err := json.Marshal(data)
	if nil != err {
		return
	}
	v.Bytes = buf
	return
}

// UnmarshalInto unmarshal containing JSON into given data reference.
func (v *BytesJSON) UnmarshalInto(ref interface{}) (err error) {
	if len(v.Bytes) == 0 {
		return
	}
	return json.Unmarshal(v.Bytes, ref)
}

// String implement string interface.
func (v BytesJSON) String() string {
	if len(v.Bytes) == 0 {
		return emptyJSONAsBunchOfBytes
	}
	return string(v.Bytes)
}
