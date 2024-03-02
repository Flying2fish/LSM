package Ylsm

import (
	"Ylsm/config"
	"Ylsm/memTable/skiplist"
	"log"
	"runtime"
	"time"
)

func Check() {
	cfg := config.GetConfig()
	timer := time.NewTimer(time.Duration(cfg.CheckInterval) * time.Millisecond)
	for range timer.C {
		db.Lock()
		checkMemory()
		db.TablesTree.Compaction()
		db.Unlock()
		runtime.GC() //垃圾回收
		timer.Reset(time.Duration(cfg.CheckInterval) * time.Millisecond)
	}
}

func checkMemory() {
	cfg := config.GetConfig()
	count := db.MemTable.GetCount()
	size := int(db.Wal.GetSize() >> 20)
	if count < cfg.Threshold && size < cfg.Level0Size {
		return
	}
	log.Printf("MemTable has %d Nodes, Wal %d MB, compressing memory\n", count, size)
	// 将内存表存储到 SsTable 中
	db.TablesTree.CreateTable(db.MemTable.GetValues(), 0)
	db.MemTable.(*skiplist.SL).Reset()
	db.Wal.Reset()
}
