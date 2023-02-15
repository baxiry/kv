package kv

import (
	"fmt"
	"sync"
)

type Map[K comparable, V any] struct {
	data sync.Map
}

// New initialaze new map
func New[K comparable, V any]() *Map[K, V] {

	return &Map[K, V]{
		data: sync.Map{},
	}
}

// Set inserts new or update old value
func (m *Map[K, V]) Set(key K, value V) {
	m.data.Store(key, value)
}

// Get selecte data
func (m *Map[K, V]) Get(key K) (V, error) {

	v, ok := m.data.Load(key)
	if !ok {
		return v.(V), fmt.Errorf("%v not found\n", key)
	}
	return v.(V), nil
}

// Delete remove data by key
func (m *Map[K, V]) Delete(key K) {
	m.data.Delete(key)
}

// HasKey inspect key is exist
func (m *Map[K, V]) HasKey(key K) bool {
	_, ok := m.data.Load(key)
	return ok
}
