package lru_cache

type LRUCache struct {
	Map     map[string]*Value
	Head    *DLinkedNode
	Tail    *DLinkedNode
	MaxSize uint32
	Size    uint32
}
type Key struct {
	DataP *string
}
type Value struct {
	DataP       *string
	LinkedNodeP *DLinkedNode
}

// DLinkedNode double linked node for lru cache.
type DLinkedNode struct {
	Data  *string
	left  *DLinkedNode
	right *DLinkedNode
}

func NewLRUCache(maxSize uint32) *LRUCache {
	return &LRUCache{
		Map:     make(map[string]*Value, 0),
		Size:    0,
		MaxSize: maxSize,
		Tail:    nil,
		Head:    nil,
	}
}

//Set set a kv
//1.map set
// 1.1 map contains ? update and skip step 2 : 1.2
// 1.2 (1) add node to head  (2)set map   (3)size ++ (4)to step 2
//2.check to size and rm node.
// 2.1 if size less or equal: skip
// 2.2 other: remove Tail node and Size --.
func (lru *LRUCache) Set(key, value *string) {
	_, contains := lru.Map[*key]
	if contains {
		mapValue := lru.Map[*key]
		mapValue.LinkedNodeP.Data = value
	} else { //not contains
		//(1) add node to head
		node := &DLinkedNode{Data: key}
		//double linked list.
		node.right = lru.Head
		//size is 0,1,lru.Head.left is null
		if lru.Size >= 2 {
			lru.Head.left = node
		}
		lru.Head = node
		//(2)set map
		v := &Value{DataP: value, LinkedNodeP: node}
		lru.Map[*key] = v
		//(3)size ++
		lru.Size++
		//(4)if Size greater that MaxSize, remove Tail node and Size --.
		lru.checkSizeRemoveNode()
	}
}

//Get return *string value if kv exist.
//if contains kv.
//1. let value to be Head
func (lru *LRUCache) Get(key *string) *string {
	value, contains := lru.Map[*key]
	if !contains {
		return nil
	} else { //contains.
		node := value.LinkedNodeP
		//nodeLeft maybe is nil.
		nodeLeft := node.left
		// if is Head not node.
		if nodeLeft != nil {
			nodeRight := node.right
			node.right = lru.Head
			nodeLeft.right = nodeRight
			lru.Head = node
		} else {
			//is Head node,init Tail
			lru.Tail = node
		}
	}
	return value.DataP

}

//checkSizeRemoveNode if Size greater that MaxSize, remove Tail node and Size --.
func (lru *LRUCache) checkSizeRemoveNode() {
	if lru.Size > lru.MaxSize {
		//1. remove Tail node
		// if have one node. tailLeft is null.
		tailLeft := lru.Tail.left
		lru.Tail = nil
		tailLeft.right = nil
		lru.Tail = tailLeft
		//2. Size --
		lru.Size--
	}
}
