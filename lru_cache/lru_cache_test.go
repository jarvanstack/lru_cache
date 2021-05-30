package lru_cache

import (
	"fmt"
	"github.com/dengjiawen8955/go_utils/string_util"
	"github.com/dengjiawen8955/go_utils/test_util"
	"testing"
)

func Test_lru_cache(t *testing.T) {
	k1 := "1"
	v1 := "10"
	k2 := "2"
	v2 := "20"
	k3 := "3"
	v3 := "30"
	cache := NewLRUCache(2)
	cache.Set(&k1, &v1)
	cache.Set(&k2, &v2)
	fmt.Printf("cache.Size=%d\n", cache.Size)
	cache.Set(&k3, &v3)
	fmt.Printf("cache.Size=%d\n", cache.Size)
	fmt.Printf("cache.Get(&k1)=%#v\n", *cache.Get(&k1))
	fmt.Printf("cache.Get(&k2)=%#v\n", cache.Get(&k2))
	fmt.Printf("cache.Get(&k3)=%#v\n", *cache.Get(&k3))
}

func Test_LRU_cache(t *testing.T) {
	var maxSize = 50_0000
	//create data
	keys := make([]string, maxSize)
	values := make([]string, maxSize)
	cache := NewLRUCache(uint32(maxSize))
	for i := 0; i < maxSize; i++ {
		key := string_util.GetLengthString(64)
		//1kb
		value := string_util.GetLengthString(1024)
		keys[i] = key
		values[i] = value
	}
	//-------
	tu := test_util.NewTestUtil(uint32(maxSize))
	tu.StartWithComment("my")
	for i := 0; i < maxSize; i++ {
		cache.Set(&keys[i], &values[i])
	}
	tu.End()
	//-------
	tu2 := test_util.NewTestUtil(uint32(maxSize))
	tu2.StartWithComment("get ")
	for i := 0; i < maxSize; i++ {
		cache.Get(&keys[i])
	}
	tu2.End()
}
