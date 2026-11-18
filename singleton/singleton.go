package singleton

import "sync"

type Singleton[T any] struct {
	mu     sync.RWMutex
	once   sync.Once
	values map[string]T
}

func (s *Singleton[T]) Get(key string) (T, bool) {
	s.once.Do(func() { s.values = make(map[string]T) })
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.values[key]
	return v, ok
}

func (s *Singleton[T]) Set(key string, value T) {
	s.once.Do(func() { s.values = make(map[string]T) })
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[key] = value
}
