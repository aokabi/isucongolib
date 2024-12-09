package isucongolib

import (
	"sync"
)

// thread-safe
type randMap[K comparable, V any] struct {
	mu       sync.Mutex
	m        map[K]V
	keys     []K // slice of keys
	randFunc func(int) int
}

func zero[T any]() T {
	var zero T
	return zero
}

func NewRandMap[K comparable, V any](f func(int) int) *randMap[K, V] {
	return &randMap[K, V]{
		m:        make(map[K]V),
		keys:     make([]K, 0),
		randFunc: f,
	}
}

func (m *randMap[K, V]) Random() (key K, value V, ok bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if len(m.keys) == 0 {
		return zero[K](), zero[V](), false
	}
	key = m.keys[m.randFunc(len(m.keys))]
	return key, m.m[key], true
}

func (m *randMap[K, V]) PopRandom() (key K, value V, ok bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if len(m.keys) == 0 {
		return zero[K](), zero[V](), false
	}
	i := m.randFunc(len(m.keys))
	key = m.keys[i]
	value = m.m[key]
	m.keys = append(m.keys[:i], m.keys[i+1:]...)
	delete(m.m, key)
	return key, value, true
}

// 既に存在するキーに対しては上書き
func (m *randMap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.m[key]; !ok {
		m.keys = append(m.keys, key)
	}
	m.m[key] = value
}

func (m *randMap[K, V]) Get(key K) (value V, ok bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	value, ok = m.m[key]
	return
}

func (m *randMap[K, V]) Pop(key K) (value V, ok bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.m[key]; !ok {
		return zero[V](), false
	}
	value = m.m[key]
	delete(m.m, key)
	for i, k := range m.keys {
		if k == key {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
	return value, true
}

func (m *randMap[K, V]) Len() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.keys)
}
