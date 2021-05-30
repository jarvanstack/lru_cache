package lru_cache

type LRUCache struct {
	Map     map[string]*DLinkedNode
	Head    *DLinkedNode
	Tail    *DLinkedNode
	MaxSize uint32
	Size    uint32
}
type Key struct {
	DataP *string
}


// DLinkedNode double linked node for lru cache.
type DLinkedNode struct {
	NodeKey *string
	NodeValue *string
	left    *DLinkedNode
	right   *DLinkedNode
}

func NewLRUCache(maxSize uint32) *LRUCache {
	return &LRUCache{
		Map:     make(map[string]*DLinkedNode, 0),
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
		node := lru.Map[*key]
		node.NodeValue = value
	} else { //not contains
		//(1) add node to head
		node := &DLinkedNode{NodeKey: key}
		//double linked list.
		node.right = lru.Head
		//size is 0,1,lru.Head.left is null
		if lru.Size >= 2 {
			lru.Head.left = node
		}
		lru.Head = node
		//(2)set map
		newNode := &DLinkedNode{NodeValue: value,NodeKey: key}
		lru.Map[*key] = newNode
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
		//nodeLeft maybe is nil.
		nodeLeft := value.left
		// if is Head not node.
		if nodeLeft != nil {
			nodeRight := value.right
			value.right = lru.Head
			nodeLeft.right = nodeRight
			lru.Head = value
		} else {
			//is Head node,init Tail
			lru.Tail = value
		}
	}
	return value.NodeValue

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
