package main

import "fmt"

func main() {
	problem3()
}

func problem3() {
	lru := &LFUEvictionManager{}
	cache := New(3, lru)
	lru.InMemoryCache = &cache

	res := cache.Add("key1", "value1") // return 0
	fmt.Println("response key 1 : ", res)
	res = cache.Add("key2", "value2") // return 0
	fmt.Println("response key 2 : ", res)
	res = cache.Add("key3", "value3") // return 0
	fmt.Println("response key 3 : ", res)
	str1 := cache.Get("key1") // return value1
	fmt.Println("value key1 : ", str1)
	res = cache.Add("key2", "value2.1") // return 1
	fmt.Println("response key 2 : ", res)

	str := cache.Get("key3") // return value3
	fmt.Println("value key3 : ", str)
	str = cache.Get("key1") // return value1
	fmt.Println("value key1 : ", str)
	str = cache.Get("key3") // return value3
	fmt.Println("value key3 : ", str)
	str = cache.Get("key2") // return value2
	fmt.Println("value key2 : ", str)

	fmt.Println(cache.Keys())
	fmt.Println(cache.Add("key4", "value4"))
	fmt.Println(cache.Keys())
	fmt.Println(cache.Clear())
	fmt.Println(cache.Keys())
}
