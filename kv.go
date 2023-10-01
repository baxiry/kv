package kv

import (
	"sync"
)

type Map[K comparable, V any] struct {
	mut  sync.RWMutex
	data map[K]V
}

// New initialaze new map
func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		data: make(map[K]V, 1),
	}
}

// Set inserts new or update old value
func (m *Map[K, V]) Set(key K, value V) {
	m.mut.Lock()
	m.data[key] = value
	m.mut.Unlock()
}

// Get selecte data
func (m *Map[K, V]) Get(key K) (V, bool) {

	m.mut.Lock()
	defer m.mut.Unlock()
	v, ok := m.data[key]
	return v, ok
}

// Delete remove data by key
func (m *Map[K, V]) Delete(key K) {
	m.mut.Lock()
	defer m.mut.Unlock()
	delete(m.data, key)
}

// HasKey inspect key is exist
func (m *Map[K, V]) HasKey(key K) bool {
	m.mut.Lock()
	defer m.mut.Unlock()
	_, ok := m.data[key]
	return ok
}
