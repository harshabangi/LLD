package main

import (
	"container/heap"
	"sync"
	"time"
)

type TTLCache struct {
	ttl      time.Duration
	items    Items
	itemFreq map[string]int
	mu       sync.Mutex
}

type Item struct {
	Key            string
	expirationTime time.Time
}

type Items []Item

func (m *Items) Len() int {
	return len(*m)
}

func (m *Items) Less(i, j int) bool {
	return (*m)[i].expirationTime.Before((*m)[j].expirationTime)
}

func (m *Items) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *Items) Push(x any) {
	*m = append(*m, x.(Item))
}

func (m *Items) Pop() any {
	old := *m
	n := len(old)
	*m = old[:n-1]
	return old[n-1]
}

func NewTTLCache(ttl time.Duration) TTLCache {
	return TTLCache{
		ttl:      ttl,
		itemFreq: make(map[string]int),
	}
}

func (t *TTLCache) Put(key string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.itemFreq[key]++
	t.cleanUpExpiredItems()
	item := Item{
		Key:            key,
		expirationTime: time.Now().Add(t.ttl),
	}
	heap.Push(&t.items, item)
}

func (t *TTLCache) Get() map[string]int {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.cleanUpExpiredItems()
	return t.itemFreq
}

func (t *TTLCache) cleanUpExpiredItems() {
	now := time.Now()

	for len(t.items) > 0 {
		top := t.items[0]

		if top.expirationTime.After(now) {
			break
		}

		heap.Pop(&t.items)

		t.itemFreq[top.Key]--
		if t.itemFreq[top.Key] == 0 {
			delete(t.itemFreq, top.Key)
		}
	}
}
