package lru_cache

import (
	"go_utils/string_util"
	"go_utils/test_util"
	"testing"
)

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
