package Ylsm

import (
	"Ylsm/config"
	"Ylsm/memTable"
	"Ylsm/memTable/skiplist"
	"Ylsm/ssTable"
	"Ylsm/wal"
	"log"
	"os"
	"sync"
)

type DB struct {
	MemTable   memTable.MemTable
	TablesTree *ssTable.TablesTree
	Wal        *wal.Wal
	sync.RWMutex
}

var db *DB

// Start 启动数据库
func Start(cfg config.Config) {
	if db != nil {
		return
	}

	log.Println("load the configuration...")
	config.Init(cfg)

	log.Println("initialize DB...")
	initDatabase(cfg.DataDir)

	log.Println("start checking in the background...")
	go Check()
}

// 初始化 DB, 从磁盘文件中还原 SsTable, Wal, MemTable
func initDatabase(dir string) {
	if _, err := os.Stat(dir); err != nil {
		log.Printf("the %s directory does not exist, the %s directory is being created.\n", dir, dir)
		if err = os.Mkdir(dir, 0666); err != nil {
			log.Panicln("fail to create the db directory:", err)
		}
	}
	db = &DB{
		MemTable:   skiplist.New(),
		Wal:        &wal.Wal{},
		TablesTree: &ssTable.TablesTree{},
	}

	log.Println("load Wal, recover MemTable...")
	db.MemTable = db.Wal.Load(dir)

	log.Println("load DB...")
	db.TablesTree.Init(dir)
}
