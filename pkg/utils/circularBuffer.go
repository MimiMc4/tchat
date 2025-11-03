package utils

import (
	"sync"
)

type CircularBuffer[T any] struct {
	mu sync.RWMutex

	data   []T
	size   int
	start  int
	count  int
	eqFunc func(a, b T) bool // function for comparing elements
}

func NewCircularBuffer[T any](size int, eqFunc func(a, b T) bool) *CircularBuffer[T] {
	return &CircularBuffer[T]{
		data:   make([]T, size),
		size:   size,
		start:  0,
		count:  0,
		eqFunc: eqFunc,
	}
}

func (b *CircularBuffer[T]) Add(element T) {
	b.mu.Lock()
	defer b.mu.Unlock()

	pos := (b.start + b.count) % b.size
	b.data[pos] = element

	if b.count < b.size {
		// if buffer not full -> increment count
		b.count++
	} else {
		// if buffer full -> element in b.start overwritten
		// -> increment start
		b.start = (b.start + 1) % b.size
	}
}

// Get all elements in order
func (b *CircularBuffer[T]) GetAll() []T {
	b.mu.RLock()
	defer b.mu.RUnlock()

	result := make([]T, b.count)
	for i := 0; i < b.count; i++ {
		idx := (b.start + i) % b.size
		result[i] = b.data[idx]
	}
	return result
}

func (b *CircularBuffer[T]) Contains(element T) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if b.eqFunc == nil {
		return false
	}

	for i := 0; i < b.count; i++ {
		idx := (b.start + i) % b.size
		if b.eqFunc(b.data[idx], element) {
			return true
		}
	}
	return false
}

func (b *CircularBuffer[T]) Clear() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.start = 0
	b.count = 0
	for i := range b.data {
		var zero T
		b.data[i] = zero
	}
}

func (b *CircularBuffer[T]) Len() int {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.count
}
