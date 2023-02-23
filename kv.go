package kv

import (
	"fmt"
	"sync"
)

type Map[K comparable, V any] struct {
	mut  sync.RWMutex
	data map[K]V
}

// New initialaze new map
func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{
		data: make(map[K]V, 0),
	}
}

// Set inserts new or update old value
func (m *Map[K, V]) Set(key K, value V) {
	m.mut.Lock()
	defer m.mut.Unlock()
	m.data[key] = value
}

// Get selecte data
func (m *Map[K, V]) Get(key K) (V, error) {

	m.mut.Lock()
	defer m.mut.Unlock()

	v, ok := m.data[key]
	if !ok {
		return v, fmt.Errorf("%v not found\n", key)
	}
	return v, nil
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
