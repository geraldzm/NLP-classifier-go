package main

import "sync"

// sincronized hash table
type HashTable struct {
	table map[string]int
	mutx  sync.Mutex
}

func NewHashTable() *HashTable {
	return &HashTable{make(map[string]int, 256), sync.Mutex{}}
}

func (h *HashTable) Add(key string) {
	h.mutx.Lock()
	defer h.mutx.Unlock()

	h.table[key]++
}
