// Package singleton provides a generic, concurrency-safe singleton helper.
//
// It allows lazy initialization of a value of any type, ensuring that the
// creation function is executed at most once.
package singleton

import (
	"sync"
	"sync/atomic"
)

// Singleton represents a lazily initialized, concurrency-safe singleton.
//
// The zero value is not usable; use New to construct a Singleton.
type Singleton[T any] struct {
	instance atomic.Pointer[T]
	once     sync.Once
	create   func() T
}

// New creates a new Singleton using the provided creation function.
//
// The create function will be called at most once, even under concurrent access.
func New[T any](create func() T) *Singleton[T] {
	return &Singleton[T]{
		create: create,
	}
}

// Get returns the singleton instance, initializing it on first use.
//
// Get is safe for concurrent use and guarantees that the underlying instance
// is created only once.
func (s *Singleton[T]) Get() T {
	if v := s.instance.Load(); v != nil {
		return *v
	}

	s.once.Do(func() {
		instance := s.create()
		s.instance.Store(&instance)
	})

	return *s.instance.Load()
}
