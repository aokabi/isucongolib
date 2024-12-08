package isucongolib

import (
	"sync"
)

type randMap struct {
	mu       sync.Mutex
	m        map[string]string
	keys     []string // slice of keys
	randFunc func(int) int
}

func NewRandMap(f func(int) int) *randMap {
	return &randMap{
		m:        make(map[string]string),
		keys:     make([]string, 0),
		randFunc: f,
	}
}

func (m *randMap) Random() (key, value string, ok bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if len(m.keys) == 0 {
		return "", "", false
	}
	key = m.keys[m.randFunc(len(m.keys))]
	return key, m.m[key], true
}

func (m *randMap) PopRandom() (key, value string, ok bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if len(m.keys) == 0 {
		return "", "", false
	}
	i := m.randFunc(len(m.keys))
	key = m.keys[i]
	value = m.m[key]
	m.keys = append(m.keys[:i], m.keys[i+1:]...)
	delete(m.m, key)
	return key, value, true
}

// 既に存在するキーに対しては上書き
func (m *randMap) Set(key, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.m[key]; !ok {
		m.keys = append(m.keys, key)
	}
	m.m[key] = value
}

func (m *randMap) Len() {
	// TODO: implement
}
