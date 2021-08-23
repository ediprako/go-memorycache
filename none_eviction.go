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
	currentHead := n.head
	for currentHead.next.next != nil {
		currentHead = currentHead.next
	}
	n.tail = currentHead
	currentHead.next = nil
	return currentHead.value
}

func (n *NoneEvictionManager) Clear() int {
	return 0
}
