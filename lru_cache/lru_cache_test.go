package lru_cache

import (
	"fmt"
	"testing"
)

func Test_function(t *testing.T) {
	k1 := "1"
	v1 := "10"
	k2 := "2"
	v2 := "20"
	cache := NewLRUCache(1)
	cache.Set(k1, v1)
	cache.Set(k2, v2)
	value1, _ := cache.Get(k1)
	fmt.Printf("value1=%#v\n", value1)
	value2, _ := cache.Get(k2)
	fmt.Printf("value2=%#v\n", value2)
}
