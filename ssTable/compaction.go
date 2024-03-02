package ssTable

import (
	"Ylsm/config"
	"Ylsm/memTable/skiplist"
	"Ylsm/value"
	"encoding/json"
	"log"
	"os"
	"time"
)

func (tt *TablesTree) Compaction() {
	cfg := config.GetConfig()

	for levelIndex := range tt.levels {
		// 转为 MB
		tableSize := int(tt.getLevelSize(levelIndex) >> 20)
		// 如果 db 文件数量 > PartSize 或者 db 文件总大小 > levelMaxSize, 触发对应层的 compaction
		if tt.getCount(levelIndex) >= cfg.PartSize || tableSize >= levelMaxSize[levelIndex] {
			log.Printf("compress level %d Sstables, the tableSize is %d MB", levelIndex, tableSize)
			tt.majorCompactionLevel(levelIndex)
		}
	}
}

// 合并第level层
func (tt *TablesTree) majorCompactionLevel(level int) {
	start := time.Now()
	defer func() {
		log.Println("completed compressing, consumption of time", time.Since(start))
	}()

	curr := tt.levels[level]
	// 将当前层的 SsTable 合并到一个 MemTable 中
	mt := skiplist.New()
	tt.Lock()
	for curr != nil {
		t := curr.table
		data := make([]byte, t.metaData.dataLen)
		// 读取 SsTable 的数据区
		if _, err := t.f.Seek(0, 0); err != nil {
			log.Println("fail to open file", t.filepath)
			panic(err)
		}
		if _, err := t.f.Read(data); err != nil {
			log.Println("fail to read file", t.filepath)
			panic(err)
		}
		// 读取每一个元素
		for k, p := range t.sparseIndex {
			if !p.Deleted { //如果没被删除
				var values value.Data
				if err := json.Unmarshal(data[p.start:(p.start+p.Len)], &values); err != nil {
					log.Fatal(err)
				}
				mt.Set(k, values.Value)
			} else {
				mt.Delete(k)
			}
		}
		curr = curr.next
	}
	tt.Unlock()
	// 将 MemTable 压缩合并成一个 SsTable
	values := mt.GetValues()
	// 最多支持 10 层, 不过也不可能到达
	newLevel := level + 1
	if newLevel > 10 {
		newLevel = 10
	}
	// 创建新的 SsTable
	tt.CreateTable(values, newLevel)
	// 重置该层
	if level < 10 {
		tt.clearLevel(level)
	}
}

// 获取第level层的文件个数
func (tt *TablesTree) getCount(level int) int {
	curr := tt.levels[level]
	count := 0
	for curr != nil {
		count++
		curr = curr.next
	}
	return count
}

// 清理第level层的文件
func (tt *TablesTree) clearLevel(level int) {
	tt.Lock()
	defer tt.Unlock()

	oldNode := tt.levels[level]
	// 清理当前层所有 SsTable
	for oldNode != nil {
		oldNode.table.Lock()
		if err := oldNode.table.f.Close(); err != nil {
			log.Println("fail to close file", oldNode.table.filepath)
			panic(err)
		}
		if err := os.Remove(oldNode.table.filepath); err != nil {
			log.Println("fail to delete file", oldNode.table.filepath)
			panic(err)
		}
		oldNode.table.Unlock()
		oldNode.table.f = nil
		oldNode.table = nil
		oldNode = oldNode.next
	}
	tt.levels[level] = nil
}
