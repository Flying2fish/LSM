package ssTable

import (
	"Ylsm/config"
	"Ylsm/value"
	"encoding/json"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

var levelMaxSize []int

type tableNode struct {
	index int
	table *SsTable
	next  *tableNode
}

// TablesTree 用于管理各层 SsTable
type TablesTree struct {
	levels []*tableNode
	sync.RWMutex
}

func (tt *TablesTree) Search(key string) (value.Data, value.SearchResult) {
	tt.RLock()
	defer tt.RUnlock()

	// 依次遍历每层 SsTable
	for _, t := range tt.levels {
		// 整理 SsTable 列表
		var tables []*SsTable
		for t != nil {
			tables = append(tables, t.table)
			t = t.next
		}

		// 从最新的 SsTable 开始查找,链表最后的是最新的
		for i := len(tables) - 1; i >= 0; i-- {
			values, searchResult := tables[i].Search(key)
			// 未找到, 则查找下一个 SsTable
			if searchResult == value.None {
				continue
			}
			// 如果找到或已被删除, 则直接返回结果
			return values, searchResult
		}
	}
	// 没有找到
	return value.Data{}, value.None
}

func (tt *TablesTree) Insert(t *SsTable, level int) (index int) {
	tt.Lock()
	defer tt.Unlock()

	curr := tt.levels[level]
	newNode := &tableNode{table: t}
	// 简单的按序插入逻辑
	if curr == nil {
		tt.levels[level] = newNode
	} else {
		for curr.next != nil {
			curr = curr.next
		}
		newNode.index = curr.index + 1
		curr.next = newNode
	}
	return newNode.index
}

func (tt *TablesTree) Init(dir string) {
	start := time.Now()
	defer func() { log.Println("load the TablesTree , consumption of time:", time.Since(start)) }()
	// 获取各层文件大小
	cfg := config.GetConfig()
	levelMaxSize = make([]int, 10)
	levelMaxSize[0] = cfg.Level0Size

	for i := 1; i < 10; i++ {
		levelMaxSize[i] = levelMaxSize[i-1] << 2
	}
	// 加载各层 db 文件
	tt.levels = make([]*tableNode, cfg.PartSize)
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Panicln("failed to read the database files:", err.Error())
	}
	for _, f := range files {
		if path.Ext(f.Name()) == ".db" {
			tt.loadDBFile(path.Join(dir, f.Name()))
		}
	}
}

func (tt *TablesTree) loadDBFile(path string) {
	start := time.Now()
	defer func() {
		log.Printf("load the %s, consumption of time: %v", path, time.Since(start))
	}()
	// 获取 db 对应的 level 和 index 信息
	level, index, err := getSsTableInfo(filepath.Base(path))
	if err != nil {
		log.Println("can not load the", path)
		return
	}
	t := &SsTable{}
	t.Load(path)
	newNode := &tableNode{index: index, table: t}

	// 根据 index 将 SsTable 插入到合适的位置
	curr := tt.levels[level]
	if curr == nil || curr.index > newNode.index {
		tt.levels[level] = newNode
		newNode.next = curr
		return
	}
	for curr.next != nil && curr.next.index <= newNode.index {
		curr = curr.next
	}
	newNode.next = curr.next
	curr.next = newNode
}

func (tt *TablesTree) CreateTable(values []value.Data, level int) *SsTable {
	//生成稀疏索引
	indexes := map[string]Index{}
	//生成数据区
	var dataArea []byte

	for _, v := range values {
		data, err := json.Marshal(v)
		if err != nil {
			log.Println("failed to Insert Key:", v.Key, err)
			continue
		}

		indexes[v.Key] = Index{
			start:   int64(len(dataArea)),
			Len:     int64(len(data)),
			Deleted: v.Deleted,
		}

		dataArea = append(dataArea, data...)
	}

	indexArea, err := json.Marshal(indexes)
	if err != nil {
		log.Fatal("an SsTable file cannot be created,", err)
	}

	meta := MetaData{
		version:    0,
		dataStart:  0,
		dataLen:    int64(len(dataArea)),
		indexStart: int64(len(dataArea)),
		indexLen:   int64(len(indexArea)),
	}

	table := &SsTable{
		metaData:    meta,
		sparseIndex: indexes,
	}

	index := tt.Insert(table, level)
	log.Printf("create a new SsTable, level: %d, index: %d\n", level, index)

	cfg := config.GetConfig()
	filePath := cfg.DataDir + "/" + strconv.Itoa(level) + "." + strconv.Itoa(index) + ".db"

	table.filepath = filePath

	writeDataToFile(filePath, dataArea, indexArea, meta)
	// 以只读的形式打开文件
	f, err := os.OpenFile(table.filepath, os.O_RDONLY, 0666)
	if err != nil {
		log.Println("fail to open file", table.filepath)
		panic(err)
	}
	table.f = f

	return table
}
