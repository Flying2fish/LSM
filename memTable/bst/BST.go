package bst

import (
	"Ylsm/memTable"
	"Ylsm/value"
	"log"
	"sync"
)

type treeNode struct {
	Data  value.Data
	Left  *treeNode
	Right *treeNode
}

type Tree struct {
	root  *treeNode
	count int //节点个数
	*sync.RWMutex
}

func (tree *Tree) GetCount() int {
	return tree.count
}

func (tree *Tree) Search(key string) (value.Data, value.SearchResult) {
	tree.RLock()
	defer tree.RUnlock()

	if tree == nil {
		log.Fatal("The tree is nil!")
	}

	curr := tree.root
	for curr != nil {
		if key == curr.Data.Key {
			if curr.Data.Deleted == false {
				return curr.Data, value.Success
			} else {
				return value.Data{}, value.Deleted
			}
		}
		if key < curr.Data.Key {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return value.Data{}, value.None
}

// 插入节点或者更改节点值
func (tree *Tree) Set(key string, values []byte) (oldValue value.Data, hasOld bool) {
	tree.RLock()
	defer tree.RUnlock()

	if tree == nil {
		log.Fatal("The tree is nil!")
	}

	curr := tree.root

	newtn := &treeNode{
		Data: value.Data{
			Key:   key,
			Value: values,
		},
	}

	for curr != nil {
		//有节点，替换即可
		if key == curr.Data.Key {
			oldValue = *curr.Data.Copy()
			curr.Data.Value = values
			curr.Data.Deleted = false

			if oldValue.Deleted {
				return value.Data{}, false
			} else {
				return oldValue, true
			}
		}
		//往左子树找
		if key < curr.Data.Key {
			if curr.Left == nil {
				curr.Left = newtn
				tree.count++
				return value.Data{}, false
			} else {
				curr = curr.Left
			}
		} else { //往右子树找
			if curr.Right == nil {
				curr.Right = newtn
				tree.count++
				return value.Data{}, false
			} else {
				curr = curr.Right
			}
		}
	}
	log.Fatalf("The tree fail to Set value, key: %s, value: %v", key, values)
	return value.Data{}, false

}

// 删除节点
func (tree *Tree) Delete(key string) (oldValue value.Data, hasOld bool) {
	tree.RLock()
	defer tree.RUnlock()

	if tree == nil {
		log.Fatal("The tree is nil")
	}

	curr := tree.root

	newtn := &treeNode{
		Data: value.Data{
			Key:     key,
			Value:   nil,
			Deleted: true,
		},
	}

	for curr != nil {
		//查找到了key
		if curr.Data.Key == key {
			if curr.Data.Deleted {
				return value.Data{}, false
			} else {
				oldValue = *curr.Data.Copy()
				curr.Data.Value = nil
				curr.Data.Deleted = true
				tree.count--
				return oldValue, true
			}
		}

		//向下继续查找
		if curr.Data.Key < key {
			if curr.Left == nil {
				curr.Left = newtn
				tree.count++
				return value.Data{}, false
			} else {
				curr = curr.Left
			}
		} else {
			if curr.Right == nil {
				curr.Right = newtn
				tree.count++
				return value.Data{}, false
			} else {
				curr = curr.Right
			}
		}
	}

	log.Fatalf("The tree fail to delete key, key: %s", key)
	return value.Data{}, false
}

// 获取所有kvdata
func (tree *Tree) GetValues() []value.Data {
	tree.RLock()
	defer tree.RUnlock()

	s := &Stack{}
	values := make([]value.Data, 0)

	curr := tree.root
	for {
		if curr != nil {
			s.Push(curr)
			curr = curr.Left
		} else {
			p, result := s.Pop()
			if !result {
				break
			}

			values = append(values, p.Data)
			curr = p.Right
		}
	}
	return values
}

// 数据迁移
func (tree *Tree) Swap() memTable.MemTable {
	tree.RLock()
	defer tree.RUnlock()

	newTree := &Tree{}

	newTree.count = tree.count
	newTree.root = tree.root
	tree.root = nil
	tree.count = 0

	return newTree
}
