package main

import "fmt"

type (
	Cache[K comparable, V any] struct {
		dados map[K]V
	}
)

func NovoCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		dados: make(map[K]V),
	}
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.dados[key] = value
}

func (c *Cache[K, V]) Get(key K) (value V, ok bool) {
	value, ok = c.dados[key]
	return
}

func main() {
	cache := NovoCache[struct {
		nome  string
		idade int
	}, int]()

	cache.Set(struct {
		nome  string
		idade int
	}{nome: "a", idade: 1}, 1)

	cache.Set(struct {
		nome  string
		idade int
	}{nome: "b", idade: 2}, 2)

	value, ok := cache.Get(struct {
		nome  string
		idade int
	}{nome: "a", idade: 1})
	if ok {
		fmt.Println(value)
	}

}
