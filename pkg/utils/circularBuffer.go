package utils

import "sync"

type CircularBuffer[T any] struct {
	mu sync.RWMutex

	data  []T
	size  int
	start int
	count int
}

func NewCircularBuffer[T any](size int) *CircularBuffer[T] {
	return &CircularBuffer[T]{
		data:  make([]T, size),
		size:  size,
		start: 0,
		count: 0,
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

func (b *CircularBuffer[T]) Len() int {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.count
}
