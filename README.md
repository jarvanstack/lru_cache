## LRU_CACHE

this is  lru cache.



### WHAT IS LRU_CACHE?

there are some instruction for goland/goupcache/lru same as this package `lru_cache`

[Golang groupcache LRU ](https://cloud.tencent.com/developer/article/1478020)


### QUICK START

1. install
```bash
go get -u  github.com/dengjiawen8955/lru_cache
```
2. example for use

```go

```
out put
```go
cache.Size=2
cache.Size=2
cache.Get(&k1)="10"
cache.Get(&k2)=(*string)(nil)
cache.Get(&k3)="30"
```

### PERFORMANCE 

* number of keys is 500000,and size of value is 1KB,
  lru_cache QPS is about 1000(W), goland/lru is  about 700(W)  

* number of keys is 500,and size value is 1MB
  lru_cache QPS is about 3827(W), 
  goland/lru is  about 1890(W)
  
there are two advantages for my_cache
1. `map[string]string` is little faster than `map[interface{]]interface{}`
2. string pointer can save some time of copy string(because string is goland is value pass)
```go
 ----50_0000 - kv<64,1024>----
 type     |  set(W)     |   get(W)
 my         170        1012/999/1080/1043
 goland     251        699/599/721/706
 ----5_0000 - kv<64,10240>----
 type     |  set     |   get
 my         310       1971/1971/[10]1805/[100]1668/
 goland     352         1492/1611/1071/[100]1310/
 ----5000 - kv<64,100K> 10000 LOOP----
 type     |  set     |   get
 my         434        2367/2100/2353
 goland     347       1614/1568/1496/1610
 ----500 - kv<64,1MB> 100_000 LOOP----
 type     |  set     |   get
 my                    3827/3523
 goland                1542/1890
```