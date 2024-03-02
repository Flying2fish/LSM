package ssTable

import (
	"Ylsm/value"
	"encoding/binary"
	"encoding/json"
	"log"
	"os"
	"sync"
)

/*
索引从数据区开始
0 ─────────────────────────────────────────────────────────►
◄──────────────────────────

	dataLen           ◄────────────────
	                      indexLen      ◄─────────────┐

┌──────────────────────────┬─────────────────┬──────────────┤
│                          │                 │              │
│           Data           │   SparseIndex   │   MetaData   │
│                          │                 │              │
└──────────────────────────┴─────────────────┴──────────────┘
*/

//SSTable 文件由 {level}.{index}.db 组成，第一个数字代表文件所在的 SSTable 层，第二个数字，表示在该层中的索引。

type SsTable struct {
	f           *os.File         //文件句柄,只读
	filepath    string           // SsTable 文件路径
	metaData    MetaData         // SsTable 元数据
	sparseIndex map[string]Index // 文件的稀疏索引列表，存在内存中
	sync.Mutex
}

type MetaData struct {
	// 版本号
	version int64
	//起始索引
	dataStart int64
	//数据区长度
	dataLen int64
	//稀疏索引区起始索引
	indexStart int64
	//稀疏索引区长度
	indexLen int64
}

type Index struct {
	start   int64 //数据起始
	Len     int64 //数据长度
	Deleted bool
}

// Load 将 db 文件 加载成 SsTable, sparseIndex 常驻内存
func (t *SsTable) Load(filepath string) {
	//初始化索引及文件地址赋值
	t.filepath = filepath
	t.sparseIndex = map[string]Index{}

	// 加载文件句柄
	f, err := os.OpenFile(t.filepath, os.O_RDONLY, 0666)
	if err != nil {
		log.Panic("fail to open file", t.filepath, ":", err.Error())
	}
	t.f = f

	// 加载元数据
	//0（io.SeekStart）：从文件的起始位置开始偏移。
	//1（io.SeekCurrent）：从当前文件指针位置开始偏移。
	//2（io.SeekEnd）：从文件的末尾位置开始偏移。
	if _, err = f.Seek(-8*5, 2); err != nil {
		log.Panic("fail to seek metadata for version:", err.Error())
	}
	_ = binary.Read(f, binary.LittleEndian, &t.metaData.version)

	if _, err = f.Seek(-8*4, 2); err != nil {
		log.Panic("fail to seek metadata for dataStart:", err.Error())
	}
	_ = binary.Read(f, binary.LittleEndian, &t.metaData.dataStart)

	if _, err = f.Seek(-8*3, 2); err != nil {
		log.Panic("fail to seek metadata for dataLen:", err.Error())
	}
	_ = binary.Read(f, binary.LittleEndian, &t.metaData.dataLen)

	if _, err = f.Seek(-8*2, 2); err != nil {
		log.Panic("fail to seek metadata for indexStart:", err.Error())
	}
	_ = binary.Read(f, binary.LittleEndian, &t.metaData.indexStart)

	if _, err = f.Seek(-8*1, 2); err != nil {
		log.Panic("fail to seek metadata for indexLen:", err.Error())
	}
	_ = binary.Read(f, binary.LittleEndian, &t.metaData.indexLen)

	// 加载稀疏索引区
	bs := make([]byte, t.metaData.indexLen)
	if _, err = f.Seek(t.metaData.indexStart, 0); err != nil {
		log.Panic("fail to seek sparseIndex:", err.Error())
	}
	if _, err = f.Read(bs); err != nil {
		log.Panic("fail to read sparseIndex:", err.Error())
	}
	if err = json.Unmarshal(bs, &t.sparseIndex); err != nil {
		log.Panic("fail to unmarshal for sparseIndex:", err.Error())
	}
	_, _ = f.Seek(0, 0)
}

func (t *SsTable) Search(key string) (values value.Data, result value.SearchResult) {
	t.Lock()
	defer t.Unlock()

	index, exist := t.sparseIndex[key]
	if !exist {
		return value.Data{}, value.None
	}
	if index.Deleted {
		return value.Data{}, value.Deleted
	}

	// 从磁盘文件中查找
	bs := make([]byte, index.Len)
	if _, err := t.f.Seek(index.start, 0); err != nil {
		log.Println("fail to seek for data:", key, err)
		return value.Data{}, value.None
	}
	if _, err := t.f.Read(bs); err != nil {
		log.Println("fail to read for data:", err)
		return value.Data{}, value.None
	}
	if err := json.Unmarshal(bs, &values); err != nil {
		log.Println("fail to unmarshal for data:", err)
		return value.Data{}, value.None
	}
	return values, value.Success
}
