package main

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
	hashMap         map[string]*Node
	size            int
}

func New(limit int, e EvictionManager) InMemoryCache {
	return InMemoryCache{
		Limit:           limit,
		evictionManager: e,
		hashMap:         make(map[string]*Node),
	}
}

func (m *InMemoryCache) Get(key string) string {
	n := m.get(m.First(), key)
	if n == nil {
		return ""
	}

	return n.value
}

func (m *InMemoryCache) addHead(node *Node) bool {
	if m.Limit == 0 {
		return false
	}
	if m.head == nil {
		m.head = node
		m.tail = node
	} else {
		m.head.prev = m.head
		node.next = m.head
		m.head = node
		m.head.prev = nil
	}
	m.size++
	return true
}

func (m *InMemoryCache) get(n *Node, key string) *Node {
	if n == nil {
		return nil
	}
	if n.key == key {
		return n
	}
	return m.get(n.Next(), key)
}

func (m *InMemoryCache) Clear() int {
	return 0
}

func (m *InMemoryCache) First() *Node {
	return m.head
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (m *InMemoryCache) Keys() []string {
	var keys []string
	list := m.First()
	for list != nil {
		keys = append(keys, list.key)
		list = list.Next()
	}

	return keys
}

func (m *InMemoryCache) Add(key string, value string) int {
	t := m.evictionManager.Push(key)
	if t == 0 {
		if m.size >= m.Limit {
			temp := m.head
			for temp.next.next != nil {
				temp = temp.next
			}
			temp.next = nil
		} else {
			m.size++
		}
		m.add(key, value)
		return t
	}

	m.update(key, value)
	return t
}

func (m *InMemoryCache) update(key string, value string) {
	for n := m.First(); n != nil; n = n.Next() {
		if n.key == key {
			n.value = value
			return
		}
	}
}

func (m *InMemoryCache) removeNode(key string) {
	n := m.get(m.First(), key)
	if n != nil {
		if n.prev != nil {
			prevNode := n.prev
			prevNode.next = n.next
		}

		if n.next != nil {
			nextNode := n.next
			nextNode.prev = n.prev
		}
	}
}

func (m *InMemoryCache) add(key string, value string) {
	n := &Node{
		key:   key,
		value: value,
	}
	if m.head == nil {
		m.head = n
	} else {
		m.tail.next = n
		n.prev = m.tail
	}
	m.tail = n
}
