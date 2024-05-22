package main

import "sync"

type SafeMap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

func NewMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V),
	}
}

func (sm *SafeMap[K, V]) Length() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	return len(sm.data)
}

func (sm *SafeMap[K, V]) put(key K, value V) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.data[key] = value
}

func (sm *SafeMap[K, V]) Get(key K) (V, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	value, ok := sm.data[key]
	return value, ok
}
