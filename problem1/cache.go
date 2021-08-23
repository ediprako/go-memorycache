package main

import (
	"errors"
)

type Node struct {
	prev  *Node
	next  *Node
	key   string
	value string
}

type InMemoryCache struct {
	Limit           int
	evictionManager EvictionManager
	head            *Node
	tail            *Node
}

func New(limit int, e EvictionManager) InMemoryCache {
	return InMemoryCache{
		Limit:           limit,
		evictionManager: e,
	}
}

func (m *InMemoryCache) Get(key string) string {
	return get(m.head, key)
}

func get(head *Node, key string) string {
	if head == nil {
		return ""
	}

	if head.key == key {
		return head.value
	}

	return get(head.next, key)
}

func (m *InMemoryCache) Clear() int {
	return clearNode(m.tail)
}

func (m *InMemoryCache) Keys() []string {
	var keys []string
	list := m.tail
	for list != nil {
		keys = append(keys, list.key)
		list = list.prev
	}
	return keys
}

func (m *InMemoryCache) Values() []string {
	var keys []string
	list := m.tail
	for list != nil {
		keys = append(keys, list.value)
		list = list.prev
	}
	return keys
}

func (m *InMemoryCache) Add(key string, value string) (int, error) {
	t := m.evictionManager.Push(key)

	c := countNode(m.head)
	if c+1 > m.Limit && t == 0 {
		return 0, errors.New("key_limit_exceeded")
	}

	if t == 0 {
		m.add(key, value)
	} else {
		m.updateNode(key, value)
	}

	return t, nil
}

func (m *InMemoryCache) add(key string, value string) {
	list := &Node{
		next:  m.head,
		key:   key,
		value: value,
	}
	if m.head != nil {
		m.head.prev = list
	}
	m.head = list

	l := m.head
	for l.next != nil {
		l = l.next
	}
	m.tail = l
}

func (m *InMemoryCache) updateNode(key string, value string) {
	updateNode(m.head, key, value)
}

func updateNode(node *Node, key string, value string) {
	if node.key == key {
		node.value = value
		return
	}

	updateNode(node.next, key, value)
}

func isKeyExist(head *Node, key string) bool {
	if head == nil {
		return false
	}

	if head.key == key && head.prev != nil {
		return true
	}

	return isKeyExist(head.next, key)
}

func countNode(head *Node) int {
	var count int
	for head != nil {
		head = head.next
		count++
	}

	return count
}

func clearNode(tail *Node) int {
	var count int
	for tail != nil {
		tail = tail.prev

		if tail != nil {
			tail.next = nil
		}

		count++
	}

	return count
}
