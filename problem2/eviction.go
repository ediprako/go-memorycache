package main

type EvictionManager interface {
	Push(key string) int
	Pop() string
	Clear() int
}
