package main

type NoneEvictionManager struct {
	*InMemoryCache
}

func (n *NoneEvictionManager) Push(key string) int {
	var result int
	if isKeyExist(n.head, key) {
		result++
	}
	return result
}

func (n *NoneEvictionManager) Pop() string {
	if n.head == nil {
		return ""
	}

	currentHead := n.head
	currentValueHead := currentHead.value
	n.head = n.head.next
	if n.head == nil {
		n.tail = nil
	}
	return currentValueHead
}

func (n *NoneEvictionManager) Clear() int {
	return 0
}
