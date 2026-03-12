// Package response provides RESTful response methods
package response

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

// New is the response constructor.
func New[T any](code int, message string, data T) *Response[T] {
	return &Response[T]{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
