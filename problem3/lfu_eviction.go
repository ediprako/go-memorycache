package main

type LFUEvictionManager struct {
	*InMemoryCache
}

func (lfu *LFUEvictionManager) Push(key string) int {
	if lfu.hashMap[key] != nil {
		return 1
	}
	return 0
}

func (lfu *LFUEvictionManager) Pop() string {
	if lfu.head == nil {
		return ""
	}

	currentHead := lfu.head
	currentValueHead := currentHead.value
	lfu.head = lfu.head.next
	if lfu.head == nil {
		lfu.tail = nil
	}
	return currentValueHead
}

func (lfu *LFUEvictionManager) Clear() int {
	return 0
}
