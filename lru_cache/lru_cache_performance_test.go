package lru_cache

import (
	"github.com/dengjiawen8955/go_utils/string_util"
	"github.com/dengjiawen8955/go_utils/test_util"
	"testing"
)

// ----MAX_SIZE:500_000,LOOP_TIMES:100 - kv<64,1024>----
// type     |  set     |   get
// lru2        209/115     955/735
// ----MAX_SIZE:50_000,LOOP_TIMES:1000 - kv<64,1024>----
// type     |  set     |   get
// lru2        271         1461
// ----MAX_SIZE:5000,LOOP_TIMES:10000 - kv<64,1024>----
// type     |  set     |   get
// lru2        217         1822
// ----MAX_SIZE:500,LOOP_TIMES:100_000 - kv<64,1024>----
// type     |  set     |   get
// lru2        274         2237
const LOOP_TIMES = 100000
const MAX_SIZE = 500
const KEY_SIZE = 64
const VALUE_SIZE = 1024 //1000 kb
func Test_performance(t *testing.T) {
	var maxSize = MAX_SIZE
	//create data
	keys := make([]string, maxSize)
	values := make([]string, maxSize)
	cache := NewLRUCache(uint32(maxSize))
	for i := 0; i < maxSize; i++ {
		key := string_util.GetLengthString(KEY_SIZE)
		//1kb
		value := string_util.GetLengthString(VALUE_SIZE)
		keys[i] = key
		values[i] = value
	}
	//-------
	tu := test_util.NewTestUtil(uint32(maxSize))
	tu.StartWithComment("my")
	for i := 0; i < maxSize; i++ {
		cache.Set(keys[i], values[i])
	}
	tu.End()
	//-------
	tu2 := test_util.NewTestUtil(uint32(maxSize)*LOOP_TIMES)
	tu2.StartWithComment("get ")
	for i := 0; i < LOOP_TIMES; i++ {
		for i := 0; i < maxSize; i++ {
			_, err := cache.Get(keys[i])
			if err != nil {
				panic(err)
			}
		}
	}
	tu2.End()
}