package main

import (
	"reflect"
	"testing"
	"time"
)

func TestTTLCache(t *testing.T) {
	// Create a TTLCache with TTL of 2 seconds
	cache := NewTTLCache(2 * time.Second)

	// Put elements in the cache
	cache.Put("A")
	cache.Put("B")

	// Sleep for 1 second
	time.Sleep(1 * time.Second)

	// Put the element "A" again
	cache.Put("A")

	// Sleep for 1 more second
	time.Sleep(1 * time.Second)

	// Put "A" again at 3 seconds
	cache.Put("A")
	cache.Put("C")

	// Get the elements after 2 seconds, expect {A: 2, B: 1}
	result := cache.Get()
	expected := map[string]int{"A": 2, "C": 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("After 2 seconds, expected: %v, got: %v", expected, result)
	}

	// Sleep for 1 more second
	time.Sleep(1 * time.Second)

	// Get the elements after 3 seconds, expect {A: 2}
	result = cache.Get()
	expected = map[string]int{"A": 1, "C": 1}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("After 3 seconds, expected: %v, got: %v", expected, result)
	}
}

func TestTTLCache_Empty(t *testing.T) {
	// Create a TTLCache with TTL of 2 seconds
	cache := NewTTLCache(2 * time.Second)

	// Get the elements from an empty cache, expect an empty map
	result := cache.Get()
	expected := map[string]int{}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected an empty map, got: %v", result)
	}
}
