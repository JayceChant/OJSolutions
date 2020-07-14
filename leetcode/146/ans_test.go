package main

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := Constructor(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)
	cache.Put(4, 4) // 1 out
	if got := cache.Get(2); got != 2 {
		t.Errorf("expect 2, got=%d", got)
	}
	cache.Put(5, 5) // 3 out
	if got := cache.Get(1); got != -1 {
		t.Errorf("expect out, got=%d", got)
	}
	if got := cache.Get(3); got != -1 {
		t.Errorf("expect out, got=%d", got)
	}
}
