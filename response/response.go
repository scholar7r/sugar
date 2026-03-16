// SPDX-License-Identifier: GPLv2
// Copyright (c) 2026 scholar7r.

// Package response provides RESTful response methods
package response

// Response represents a standard RESTful response.
//
// T is the type of the data payload. It includes a status code, a message,
// and optional data.
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
