package ssTable

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

// 获取 db 对应的 level 和 index 信息
func getSsTableInfo(filename string) (level int, index int, err error) {
	n, err := fmt.Sscanf(filename, "%d.%d.db", &level, &index)
	if n != 2 || err != nil {
		return 0, 0, fmt.Errorf("incorrect data file filename for SsTable: %q", filename)
	}
	return level, index, nil
}

// 获取 db 数据文件大小
func (t *SsTable) getDBSize() int64 {
	info, err := os.Stat(t.filepath)
	if err != nil {
		log.Fatal(err)
	}
	return info.Size()
}

// 获取指定层的 SsTable 总文件大小
func (tt *TablesTree) getLevelSize(level int) (size int64) {
	curr := tt.levels[level]
	for curr != nil {
		size += curr.table.getDBSize()
		curr = curr.next
	}
	return size
}

// 将数据按顺序 <data, sparseIndex, metaInfo> 写入 db 文件
func writeDataToFile(filepath string, dataArea []byte, indexArea []byte, meta MetaData) {
	//打开文件
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("fail to create file:", err)
	}

	//将data和index写入
	if _, err = f.Write(dataArea); err != nil {
		log.Fatal("fail to write dataArea:", err)
	}
	if _, err = f.Write(indexArea); err != nil {
		log.Fatal("fail to write indexArea:", err)
	}

	if err = binary.Write(f, binary.LittleEndian, &meta.version); err != nil {
		log.Fatal("fail to write metaInfo.version:", err)
	}
	if err = binary.Write(f, binary.LittleEndian, &meta.dataStart); err != nil {
		log.Fatal("fail to write metaInfo.dataStart:", err)
	}
	if err = binary.Write(f, binary.LittleEndian, &meta.dataLen); err != nil {
		log.Fatal("fail to write metaInfo.dataLen:", err)
	}
	if err = binary.Write(f, binary.LittleEndian, &meta.indexStart); err != nil {
		log.Fatal("fail to write metaInfo.indexStart:", err)
	}
	if err = binary.Write(f, binary.LittleEndian, &meta.indexLen); err != nil {
		log.Fatal("fail to write metaInfo.indexLen:", err)
	}
	//把文件刷新到磁盘上，持久化保存
	if err = f.Sync(); err != nil {
		log.Fatal("fail to write metaInfo:", err)
	}
	if err = f.Close(); err != nil {
		log.Fatal("fail to close .db:", err)
	}
}
