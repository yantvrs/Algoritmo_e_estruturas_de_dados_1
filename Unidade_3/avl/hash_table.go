package main

import("fmt" "errors")

type Map interface {
	Put(key int, value string)
	Get(key int) (string, error)
	Remove(key int) error
	Size() int
	LoadFactor() float32
	Init()
}

/* Usando generics
type Pair[T, U any] struct {
    key  T
    value U
}*/

type Tuple struct {
	key   int
	value string
}

var emptyTuple Tuple = Tuple{}

type HashTable struct {
	buckets          []Tuple
	elementsInserted int
}

