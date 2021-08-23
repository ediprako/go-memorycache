package main

type LRUEvictionManager struct {
	*InMemoryCache
}

func (lru *LRUEvictionManager) Push(key string) int {
	if isKeyExist(lru.head, key) {
		return 1
	}
	return 0
}

func isKeyExist(n *Node, key string) bool {
	if n == nil {
		return false
	}
	if n.key == key {
		return true
	}

	return isKeyExist(n.Next(), key)
}

func (lru *LRUEvictionManager) Pop() string {
	if lru.head == nil {
		return ""
	}

	currentHead := lru.head
	currentValueHead := currentHead.value
	lru.head = lru.head.next
	if lru.head == nil {
		lru.tail = nil
	}
	return currentValueHead
}

func (lru *LRUEvictionManager) Clear() int {
	return 0
}
