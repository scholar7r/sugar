package singleton

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestSingleton_CreateOnce(t *testing.T) {
	var count int32

	s := New(func() int {
		atomic.AddInt32(&count, 1)
		return 7
	})

	v := s.Get()

	if v != 7 {
		t.Fatalf("unexpected value: %v", v)
	}

	if atomic.LoadInt32(&count) != 1 {
		t.Fatalf("create called %d times, want 1", count)
	}
}

func TestSingleton_Concurrent(t *testing.T) {
	var count int32

	s := New(func() int {
		atomic.AddInt32(&count, 1)
		return 7
	})

	const goroutines = 100
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for range goroutines {
		go func() {
			defer wg.Done()
			_ = s.Get()
		}()
	}

	wg.Wait()

	if atomic.LoadInt32(&count) != 1 {
		t.Fatalf("create called %d times, want 1", count)
	}
}
