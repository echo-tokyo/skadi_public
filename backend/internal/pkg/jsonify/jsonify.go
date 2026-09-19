// Package jsonify provides Jsonify interface to (de)serialize any struct.
// It contains default implementation of Jsonify via standart json library.
package jsonify

// Jsonify describes tool to (de)serialize any struct.
type Jsonify interface {
	Marshal(v any) ([]byte, error)
	Unmarshal(data []byte, v any) error
}
