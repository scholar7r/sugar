// SPDX-License-Identifier: GPLv2
// Copyright (c) 2026 scholar7r.

// Package singleton provides a generic, concurrency-safe Singleton type.
//
// Singleton allows lazy initialization of a value of any type, ensuring that
// the provided creation function is executed at most once, even under
// concurrent access.
package singleton

import (
	"sync"
)

// Singleton represents a lazily initialized, concurrency-safe singleton of any type.
//
// The zero value of Singleton is not usable because the creation function is nil.
// Use New to construct a usable Singleton instance.
type Singleton[T any] struct {
	instance *T
	once     sync.Once
	create   func() T
}

// New constructs a new Singleton using the provided creation function.
//
// The create function must not be nil and will be called at most once, even
// when multiple goroutines call Get concurrently. The singleton instance is
// initialized lazily on first use.
func New[T any](create func() T) *Singleton[T] {
	return &Singleton[T]{
		create: create,
	}
}

// Get returns a pointer to the singleton instance, initializing it on first use.
//
// Get is safe for concurrent use. The creation function is guaranteed to be
// executed at most once, and the returned instance will always point to the
// same object of type T.
// is created only once.
func (s *Singleton[T]) Get() *T {
	s.once.Do(func() {
		instance := s.create()
		s.instance = &instance
	})

	return s.instance
}
