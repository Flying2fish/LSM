package wal

import (
	"Ylsm/memTable"
	"Ylsm/memTable/skiplist"
	"Ylsm/value"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
	"os"
	"path"
	"sync"
	"time"
)

type Wal struct {
	f    *os.File
	path string
	sync.Mutex
}

// GetSize 获取 Wal大小, 单位字节
func (w *Wal) GetSize() int64 {
	info, err := w.f.Stat()
	if err != nil {
		log.Fatal(err)
		return -1
	}
	return info.Size()
}

// Write 将数组增加、修改和删除操作写入Wal
func (w *Wal) Write(value value.Data) {
	w.Lock()
	defer w.Unlock()
	data, _ := json.Marshal(value)
	if err := binary.Write(w.f, binary.LittleEndian, int64(len(data))); err != nil {
		log.Panicln("fail to write the wal.log: " + err.Error())
	}

	if err := binary.Write(w.f, binary.LittleEndian, data); err != nil {
		log.Panicln("fail to write the wal.log: " + err.Error())
	}
}

func (w *Wal) Load(dir string) memTable.MemTable {
	start := time.Now()
	defer func() {
		log.Println("load the wal.log, consumption of time:", time.Since(start))
	}()

	walPath := path.Join(dir, "wal.log")
	f, err := os.OpenFile(walPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panicln("fail to open the wal.log:", err)
	}

	w.Lock()
	defer w.Unlock()

	w.f = f
	w.path = walPath
	size := w.GetSize()
	t := skiplist.New()

	if size == 0 {
		return t
	}
	if _, err = w.f.Seek(0, 0); err != nil {
		log.Panicln("fail to open the wal.log:", err)
	}

	// 文件指针移动到最后，以便追加
	defer func(f *os.File) {
		_, err = f.Seek(0, 2)
		if err != nil {
			log.Println("fail to open the wal.log")
			panic(err)
		}
	}(w.f)

	// 将文件内容全部读取到内存
	data := make([]byte, size)
	if _, err = w.f.Read(data); err != nil {
		log.Panicln("fail to read the wal.log:", err)
	}

	dataLen := int64(0) // 元素的字节数量初始化
	index := int64(0)   // 当前索引

	for index < size {
		// 获取元素的字节长度
		dataLenArea := data[index : index+8]
		buf := bytes.NewReader(dataLenArea) //声明reader
		if err = binary.Read(buf, binary.LittleEndian, &dataLen); err != nil {
			log.Panicln("fail to read for dataLen:", err)
		}
		// 将元素的所有字节读取出来，并还原为 kv.Data
		index += 8
		var values value.Data
		dataArea := data[index : index+dataLen]
		if err = json.Unmarshal(dataArea, &values); err != nil {
			log.Panicln("fail to unmarshal the data:", err)
		}
		if values.Deleted {
			t.Delete(values.Key)
		} else {
			t.Set(values.Key, values.Value)
		}
		index += dataLen
	}
	return t
}

func (w *Wal) Reset() {
	w.Lock()
	defer w.Unlock()
	if err := w.f.Close(); err != nil {
		log.Panicln(err)
	}
	if err := os.Remove(w.f.Name()); err != nil {
		log.Panicln(err)
	}
	time.Sleep(time.Millisecond)
	f, err := os.OpenFile(w.path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panicln(err)
	}
	w.f = f
}
