package lrucache

import (
	"fmt"
	"testing"
)

func TestLRCCache_Self1(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(1, 3)
}

func TestLRUCache_Case1(t *testing.T) {
	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // 返回  1
	cache.Put(3, 3)           // 该操作会使得密钥 2 作废
	fmt.Println(cache.Get(2)) // 返回 -1 (未找到)
	cache.Put(4, 4)           // 该操作会使得密钥 1 作废
	fmt.Println(cache.Get(1)) // 返回 -1 (未找到)
	fmt.Println(cache.Get(3)) // 返回  3
	fmt.Println(cache.Get(4)) // 返回  4
}

func TestLRUCache_Case2(t *testing.T) {
	cache := Constructor(10)
	cache.Put(10, 13)
	cache.Put(3, 17)
	cache.Put(6, 11)
	cache.Put(10, 5)
	cache.Put(9, 10)
	cache.Get(13)
	cache.Put(2, 19)
	cache.Get(2)
	cache.Get(3)
	cache.Put(5, 25)
	cache.Get(8)
	cache.Put(9, 22)
	cache.Put(5, 5)
	cache.Put(1, 30)
	cache.Get(11)
	cache.Put(9, 12)
	cache.Get(7)
	cache.Get(5)
	cache.Get(8)
	cache.Get(9)
	cache.Put(4, 30)
	cache.Put(9, 3)
	cache.Get(9)
	cache.Get(10)
	cache.Get(10)
	cache.Put(6, 14)
	cache.Put(3, 1)
}

func TestConstructor_1(t *testing.T) {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	cache.Get(2)
	cache.Put(1, 1)
	cache.Put(4, 1)
	cache.Get(2)
}