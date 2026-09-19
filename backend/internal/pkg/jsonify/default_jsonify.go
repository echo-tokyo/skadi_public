package jsonify

import "encoding/json"

// Ensure default jsonify implements interface.
var _ Jsonify = (*DefaultJsonify)(nil)

// DefaultJsonify implements [Jsonify] via standart json library.
type DefaultJsonify struct{}

// New returns a new instance of [DefaultJsonify].
func New() *DefaultJsonify {
	return &DefaultJsonify{}
}

// Marshal serializes struct into JSON-bytes.
func (j DefaultJsonify) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal deserializes JSON from bytes into struct.
func (j DefaultJsonify) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
