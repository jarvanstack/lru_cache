package lru_cache

import (
	"fmt"
	"testing"
)



//* string can not be key,because *key will changed. use string for key.
func TestLRUCache_Set_Key_str_p_err(t *testing.T) {
	cache := NewLRUCache(2)
	k1 := "1"
	v1 := "10"
	k2 := "2"
	v2 := "20"
	cache.Set(&k1,&v1)
	cache.Set(&k2,&v2)
	key := "1"
	fmt.Printf("&key=%#v\n", &key)
	fmt.Printf("&k1=%#v\n", &k1)
	value := cache.Get(&k1)
	no_same_pointer_value := cache.Get(&key)
	fmt.Printf("value=%#v\n", *value)
	fmt.Printf("no_same_pointer_value=%#v\n", *no_same_pointer_value)
}

func Test_Map_ok(t *testing.T) {
	m2 := make(map[int]int)
	m2[1] = 10
	value,ok := m2[1]
	if !ok {
		fmt.Printf("%s\n", "!ok")
	}
	fmt.Printf("value=%#v\n", value)
}