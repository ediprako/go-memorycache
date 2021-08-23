package main

import "fmt"

func main() {
	problem2()
}

func problem2() {
	lru := &LRUEvictionManager{}
	cache := New(3, lru)
	lru.InMemoryCache = &cache

	res := cache.Add("key1", "value1") // return 0
	fmt.Println("response key 1 : ", res)
	res = cache.Add("key2", "value2") // return 0
	fmt.Println("response key 2 : ", res)
	res = cache.Add("key3", "value3") // return 0
	fmt.Println("response key 3 : ", res)
	res = cache.Add("key2", "value2.1") // return 1
	fmt.Println("response key 2 : ", res)
	res = cache.Add("key4", "value2.1") // return 1
	fmt.Println("response key 4 : ", res)

	str := cache.Get("key3") // return value3
	fmt.Println("value key3 : ", str)

	keys1 := cache.Keys() // return ['key1', 'key2', 'key3']
	fmt.Println("keys : ", keys1)

	str = cache.Get("key1") // return value1
	fmt.Println("value key1 : ", str)
	str = cache.Get("key3") // return value3
	fmt.Println("value key3 : ", str)
	keys := cache.Keys() // return ['key1', 'key2', 'key3']
	fmt.Println("keys : ", keys)

	totalClear := cache.Clear()
	fmt.Println("clear : ", totalClear)

	keys = cache.Keys() // return []
	fmt.Println(keys)

}
