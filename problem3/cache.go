package main

import (
	"container/heap"
)

type Node struct {
	prev  *Node
	next  *Node
	key   string
	value string
	freq  int
}

type keyHeap struct {
	freq int
	key  string
}

type keyHeaps []keyHeap

func (h keyHeaps) Len() int           { return len(h) }
func (h keyHeaps) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h keyHeaps) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *keyHeaps) Push(x interface{}) {
	old := *h
	v := x.(keyHeap)
	var res keyHeaps
	for k := range old {
		if old[k].key == v.key {
			res = append(res, old[:k]...)
			res = append(res, old[k+1:]...)
			*h = res
			break
		}
	}
	*h = append(*h, v)
}

func (h *keyHeaps) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type InMemoryCache struct {
	Limit           int
	evictionManager EvictionManager
	head            *Node
	tail            *Node
	hashMap         map[string]*Node
	size            int
	heaps           keyHeaps
}

func New(limit int, e EvictionManager) InMemoryCache {
	return InMemoryCache{
		Limit:           limit,
		evictionManager: e,
		hashMap:         make(map[string]*Node),
	}
}

func (m *InMemoryCache) Get(key string) string {
	if m.hashMap[key] == nil {
		return ""
	}

	m.hashMap[key].freq++
	heap.Push(&m.heaps, keyHeap{
		freq: m.hashMap[key].freq,
		key:  key,
	})
	return m.hashMap[key].value
}

func (m *InMemoryCache) Keys() (result []string) {
	list := m.head
	for list != nil {
		result = append(result, list.key)
		list = list.next
	}
	return result
}

func (m *InMemoryCache) Add(key string, value string) int {
	if m.hashMap[key] == nil {
		if m.size >= m.Limit {
			x := heap.Pop(&m.heaps)
			m.remove(x.(keyHeap).key)
			m.size--
		}

		n := &Node{
			key:   key,
			value: value,
			freq:  1,
		}
		if m.head == nil {
			m.head = n
		} else {
			m.tail.next = n
			n.prev = m.tail
		}
		m.tail = n

		m.size++
		m.hashMap[key] = n
		heap.Push(&m.heaps, keyHeap{
			freq: n.freq,
			key:  key,
		})
	}

	return 1
}

func (m *InMemoryCache) remove(key string) {
	n := m.hashMap[key]
	if n != nil {
		if n.prev != nil {
			prevNode := n.prev
			prevNode.next = n.next
		}

		if n.next != nil {
			nextNode := n.next
			nextNode.prev = n.prev
		}

		delete(m.hashMap, key)
	}
}

func (m *InMemoryCache) Clear() int {
	count := m.size
	temp := m.head
	for temp.next != nil {
		temp = temp.next
		temp.prev = nil
		m.head = temp
	}

	m.head = nil

	temp = m.tail
	for temp.prev != nil {
		temp = temp.prev
		temp.next = nil
		m.tail = temp
	}
	m.size = 0
	return count
}
