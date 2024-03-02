package skiplist

import (
	"Ylsm/memTable"
	"Ylsm/value"
	"log"
	"math/rand"
	"sync"
)

const maxLevel = 32
const pFactor = 0.25

type Node struct {
	value   value.Data
	forward []*Node //forward[i]代表第i层的下一个节点
}

type SL struct {
	head  *Node //头节点，不保存数据，数据从头节点的下一个开始
	level int
	count int //节点总个数
	sync.RWMutex
}

var _ memTable.MemTable = (*SL)(nil)

func New() *SL {
	sl := SL{}
	sl.head = &Node{
		value:   value.Data{Key: "", Value: nil, Deleted: true},
		forward: make([]*Node, maxLevel),
	}
	return &sl
}

func (sl *SL) randomLevel() int {
	lv := 1
	for lv < maxLevel && rand.Float64() < pFactor {
		lv++
	}
	return lv
}

// GetCount 获取树中的元素数量
func (sl *SL) GetCount() int {
	return sl.count
}

// Search 查找 Key 的值
func (sl *SL) Search(key string) (value.Data, value.SearchResult) {
	sl.RLock()
	defer sl.RUnlock()
	curr := sl.head
	for i := sl.level - 1; i >= 0; i-- { //逐层向下寻找
		for curr.forward[i] != nil && curr.forward[i].value.Key < key {
			curr = curr.forward[i]
		}
	}
	curr = curr.forward[0]
	if curr != nil && curr.value.Key == key {
		if curr.value.Deleted {
			return value.Data{}, value.Deleted
		}
		return curr.value, value.Success
	}
	return value.Data{}, value.None
}

// Set 设置 Key 的值并返回旧值
func (sl *SL) Set(key string, values []byte) (oldValue value.Data, hasOld bool) {
	sl.Lock()
	defer sl.Unlock()
	update := make([]*Node, maxLevel)
	for i := range update {
		update[i] = sl.head
	}
	curr := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for curr.forward[i] != nil && curr.forward[i].value.Key < key {
			curr = curr.forward[i]
		}
		update[i] = curr //给每一层的forward赋值，如果没有节点需要用这个赋值
	}
	curr = curr.forward[0]
	// 如果有这个节点
	if curr != nil && curr.value.Key == key {
		if curr.value.Deleted {
			curr.value.Deleted = false
			curr.value.Value = values
			return value.Data{}, false
		} else {
			oldValue = *curr.value.Copy()
			curr.value.Value = values
			return oldValue, true
		}
	}
	sl.count++
	lv := sl.randomLevel()
	sl.level = max(sl.level, lv)
	newNode := &Node{
		value:   value.Data{Key: key, Value: values, Deleted: false},
		forward: make([]*Node, lv),
	}
	for i, node := range update[:lv] { //插入
		newNode.forward[i] = node.forward[i]
		node.forward[i] = newNode
	}
	return value.Data{}, false
}

// Delete 删除 key 并返回旧值
func (sl *SL) Delete(key string) (oldValue value.Data, hasOld bool) {
	sl.Lock()
	defer sl.Unlock()
	update := make([]*Node, maxLevel)
	for i := range update {
		update[i] = sl.head
	}
	curr := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for curr.forward[i] != nil && curr.forward[i].value.Key < key {
			curr = curr.forward[i]
		}
		update[i] = curr
	}
	curr = curr.forward[0]
	// 如果有这个节点
	if curr != nil && curr.value.Key == key {
		if curr.value.Deleted {
			return value.Data{}, false
		} else {
			curr.value.Value = nil
			curr.value.Deleted = true
			return *curr.value.Copy(), true
		}
	}
	// 没有这节点的话就要插入
	sl.count++
	lv := sl.randomLevel()
	sl.level = max(sl.level, lv)
	newNode := &Node{
		value:   value.Data{Key: key, Value: nil, Deleted: true},
		forward: make([]*Node, lv),
	}
	for i, node := range update[:lv] {
		newNode.forward[i] = node.forward[i]
		node.forward[i] = newNode
	}
	return value.Data{}, false
}

// GetValues 获取树中的所有元素
func (sl *SL) GetValues() (values []value.Data) {
	sl.RLock()
	defer sl.RUnlock()
	curr := sl.head.forward[0]
	for curr != nil {
		values = append(values, curr.value)
		curr = curr.forward[0]
	}
	return values
}

func (sl *SL) Reset() {
	sl.Lock()
	defer sl.Unlock()
	sl.count = 0
	sl.level = 0
	sl.head = &Node{
		value:   value.Data{Key: "", Value: nil, Deleted: true},
		forward: make([]*Node, maxLevel),
	}
}

func (sl *SL) Swap() memTable.MemTable {
	sl.Lock()
	defer sl.Unlock()
	// 生成tmpSL
	tmpSL := &SL{}
	tmpSL.count = sl.count
	tmpSL.level = sl.level
	tmpSL.head = sl.head

	// 将 sl 初始化
	sl.count = 0
	sl.level = 0
	sl.head = &Node{
		value:   value.Data{Key: "", Value: nil, Deleted: true},
		forward: make([]*Node, maxLevel),
	}
	return tmpSL
}

func (sl *SL) Show() {
	curr := sl.head.forward[0]
	for curr != nil {
		log.Println(curr.value.Key, curr.value.Deleted)
		curr = curr.forward[0]
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
