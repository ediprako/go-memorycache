package main

import "fmt"

func main() {
	problem1()
}

func problem1() {
	none := &NoneEvictionManager{}
	cache := New(3, none)
	none.InMemoryCache = &cache

	res, _ := cache.Add("key1", "value1") // return 0
	fmt.Println("response key 1 : ", res)
	res, _ = cache.Add("key2", "value2") // return 0
	fmt.Println("response key 2 : ", res)
	res, _ = cache.Add("key3", "value3") // return 0
	fmt.Println("response key 3 : ", res)

	res, _ = cache.Add("key2", "value2.1") // return 1
	fmt.Println("response key 2 : ", res)
	fmt.Println("values", cache.Values())
	str := cache.Get("key3") // return value3
	fmt.Println("value key3 : ", str)
	str = cache.Get("key1") // return value1
	fmt.Println("value key1 : ", str)
	str = cache.Get("key3") // return value3
	fmt.Println("value key3 : ", str)
	keys := cache.Keys() // return ['key1', 'key2', 'key3']
	fmt.Println("keys : ", keys)
	res, err := cache.Add("key4", "value4")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	keys = cache.Keys() // return ['key1', 'key2', 'key3']
	fmt.Println(keys)

	totalClear := cache.Clear()
	fmt.Println("clear : ", totalClear)

	keys = cache.Keys() // return []
	fmt.Println(keys)

}
