## 测试

1. insert() 插入11881376(26的5次方)条数据
2. queryAll() 查询11881376(26的5次方)条存在的数据
3. deleteAll() 删除11881376(26的5次方)条存在的数据
4. deleteAbsent() 删除1条不存在的数据
5. queryAbsent() 查询11881376(26的5次方)条不存在的数据
```
C:\Users\Administrator\AppData\Local\JetBrains\GoLand2023.3\tmp\GoLand\___1go_build_Ylsm_demo.exe
2024/03/02 20:04:12 E:/Codes/Golang/Self_Write/LSM/Ylsm/Ylsm.go:29: load the configuration...
2024/03/02 20:04:12 E:/Codes/Golang/Self_Write/LSM/Ylsm/Ylsm.go:32: initialize DB...
2024/03/02 20:04:12 E:/Codes/Golang/Self_Write/LSM/Ylsm/Ylsm.go:53: load Wal, recover MemTable...
2024/03/02 20:04:12 E:/Codes/Golang/Self_Write/LSM/Ylsm/wal/Wal.go:50: load the wal.log, consumption of time: 0s
2024/03/02 20:04:12 E:/Codes/Golang/Self_Write/LSM/Ylsm/Ylsm.go:56: load DB...
2024/03/02 20:04:12 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:79: load the TablesTree , consumption of time: 0s
2024/03/02 20:04:12 E:/Codes/Golang/Self_Write/LSM/Ylsm/Ylsm.go:35: start checking in the background...
2024/03/02 20:04:16 E:/Codes/Golang/Self_Write/LSM/Ylsm/check.go:31: MemTable has 896302 Nodes, Wal 107 MB, compressing memory
2024/03/02 20:04:18 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 0, index: 0
2024/03/02 20:04:18 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 0 Sstables, the tableSize is 131 MB
2024/03/02 20:04:23 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 1, index: 0
2024/03/02 20:04:23 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 4.9509762s
2024/03/02 20:04:27 E:/Codes/Golang/Self_Write/LSM/Ylsm/check.go:31: MemTable has 850153 Nodes, Wal 102 MB, compressing memory
2024/03/02 20:04:28 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 0, index: 0
2024/03/02 20:04:28 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 0 Sstables, the tableSize is 124 MB
2024/03/02 20:04:33 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 1, index: 1
2024/03/02 20:04:33 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 4.9797799s
2024/03/02 20:04:37 E:/Codes/Golang/Self_Write/LSM/Ylsm/check.go:31: MemTable has 938481 Nodes, Wal 112 MB, compressing memory
2024/03/02 20:04:38 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 0, index: 0
2024/03/02 20:04:39 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 0 Sstables, the tableSize is 137 MB
2024/03/02 20:04:44 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 1, index: 2
2024/03/02 20:04:44 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 5.4060617s
2024/03/02 20:04:48 E:/Codes/Golang/Self_Write/LSM/Ylsm/check.go:31: MemTable has 966006 Nodes, Wal 116 MB, compressing memory
2024/03/02 20:04:49 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 0, index: 0
2024/03/02 20:04:50 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 0 Sstables, the tableSize is 141 MB
2024/03/02 20:04:56 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 1, index: 3
2024/03/02 20:04:57 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 6.9700893s
2024/03/02 20:04:57 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 1 Sstables, the tableSize is 536 MB
2024/03/02 20:05:20 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 2, index: 0
2024/03/02 20:05:20 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 23.6998844s
2024/03/02 20:05:24 E:/Codes/Golang/Self_Write/LSM/Ylsm/check.go:31: MemTable has 891832 Nodes, Wal 107 MB, compressing memory
2024/03/02 20:05:26 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 0, index: 0
2024/03/02 20:05:26 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 0 Sstables, the tableSize is 130 MB
2024/03/02 20:05:31 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 1, index: 0
2024/03/02 20:05:32 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 5.788168s
2024/03/02 20:05:36 E:/Codes/Golang/Self_Write/LSM/Ylsm/check.go:31: MemTable has 896923 Nodes, Wal 107 MB, compressing memory
2024/03/02 20:05:37 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 0, index: 0
2024/03/02 20:05:37 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 0 Sstables, the tableSize is 131 MB
2024/03/02 20:05:44 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 1, index: 1
2024/03/02 20:05:44 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 6.3357496s
2024/03/02 20:05:48 E:/Codes/Golang/Self_Write/LSM/Ylsm/check.go:31: MemTable has 880366 Nodes, Wal 105 MB, compressing memory
2024/03/02 20:05:49 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 0, index: 0
2024/03/02 20:05:49 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 0 Sstables, the tableSize is 129 MB
2024/03/02 20:05:54 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 1, index: 2
2024/03/02 20:05:54 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 5.1218317s
2024/03/02 20:05:59 E:/Codes/Golang/Self_Write/LSM/Ylsm/check.go:31: MemTable has 970038 Nodes, Wal 116 MB, compressing memory
2024/03/02 20:06:00 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 0, index: 0
2024/03/02 20:06:00 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 0 Sstables, the tableSize is 142 MB
2024/03/02 20:06:06 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 1, index: 3
2024/03/02 20:06:06 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 6.120727s
2024/03/02 20:06:06 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 1 Sstables, the tableSize is 534 MB
2024/03/02 20:06:31 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 2, index: 1
2024/03/02 20:06:31 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 25.3492283s
2024/03/02 20:06:36 E:/Codes/Golang/Self_Write/LSM/Ylsm/check.go:31: MemTable has 931426 Nodes, Wal 111 MB, compressing memory
2024/03/02 20:06:37 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 0, index: 0
2024/03/02 20:06:37 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 0 Sstables, the tableSize is 136 MB
2024/03/02 20:06:43 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 1, index: 0
2024/03/02 20:06:43 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 5.8452117s
2024/03/02 20:06:48 E:/Codes/Golang/Self_Write/LSM/Ylsm/check.go:31: MemTable has 925566 Nodes, Wal 111 MB, compressing memory
2024/03/02 20:06:49 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 0, index: 0
2024/03/02 20:06:49 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 0 Sstables, the tableSize is 135 MB
2024/03/02 20:06:55 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 1, index: 1
2024/03/02 20:06:55 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 5.8888228s
2024/03/02 20:06:59 E:/Codes/Golang/Self_Write/LSM/Ylsm/check.go:31: MemTable has 896213 Nodes, Wal 107 MB, compressing memory
2024/03/02 20:07:00 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 0, index: 0
2024/03/02 20:07:00 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 0 Sstables, the tableSize is 131 MB
2024/03/02 20:07:06 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 1, index: 2
2024/03/02 20:07:06 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 5.26292s
2024/03/02 20:07:06 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 1 Sstables, the tableSize is 404 MB
2024/03/02 20:07:25 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 2, index: 2
2024/03/02 20:07:26 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 19.9101122s
2024/03/02 20:07:30 E:/Codes/Golang/Self_Write/LSM/Ylsm/check.go:31: MemTable has 985239 Nodes, Wal 118 MB, compressing memory
2024/03/02 20:07:31 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 0, index: 0
2024/03/02 20:07:32 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 0 Sstables, the tableSize is 144 MB
2024/03/02 20:07:37 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 1, index: 0
2024/03/02 20:07:37 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 5.8217762s
2024/03/02 20:07:41 E:/Codes/Golang/Self_Write/LSM/Ylsm/demo/main.go:62: insert count: 11881376
2024/03/02 20:07:42 E:/Codes/Golang/Self_Write/LSM/Ylsm/check.go:31: MemTable has 852831 Nodes, Wal 102 MB, compressing memory
2024/03/02 20:07:43 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 0, index: 0
2024/03/02 20:07:43 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 0 Sstables, the tableSize is 125 MB
2024/03/02 20:07:49 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 1, index: 1
2024/03/02 20:07:49 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 6.0174133s
2024/03/02 20:08:48 E:/Codes/Golang/Self_Write/LSM/Ylsm/demo/main.go:90: error count for queryAll: 0
2024/03/02 20:08:58 E:/Codes/Golang/Self_Write/LSM/Ylsm/check.go:31: MemTable has 2265844 Nodes, Wal 110 MB, compressing memory
2024/03/02 20:09:01 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 0, index: 0
2024/03/02 20:09:01 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 0 Sstables, the tableSize is 166 MB
2024/03/02 20:09:12 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 1, index: 2
2024/03/02 20:09:12 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 10.9455794s
2024/03/02 20:09:12 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 1 Sstables, the tableSize is 436 MB
2024/03/02 20:09:36 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 2, index: 3
2024/03/02 20:09:36 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 23.7148932s
2024/03/02 20:09:36 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 2 Sstables, the tableSize is 1911 MB
2024/03/02 20:11:36 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 3, index: 0
2024/03/02 20:11:37 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 2m1.3999628s
2024/03/02 20:11:48 E:/Codes/Golang/Self_Write/LSM/Ylsm/check.go:31: MemTable has 2210278 Nodes, Wal 107 MB, compressing memory
2024/03/02 20:11:51 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 0, index: 0
2024/03/02 20:11:51 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 0 Sstables, the tableSize is 162 MB
2024/03/02 20:12:03 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 1, index: 0
2024/03/02 20:12:04 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 12.5323592s
2024/03/02 20:12:13 E:/Codes/Golang/Self_Write/LSM/Ylsm/check.go:31: MemTable has 2129270 Nodes, Wal 103 MB, compressing memory
2024/03/02 20:12:16 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 0, index: 0
2024/03/02 20:12:16 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 0 Sstables, the tableSize is 156 MB
2024/03/02 20:12:28 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 1, index: 1
2024/03/02 20:12:28 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 11.6673236s
2024/03/02 20:12:39 E:/Codes/Golang/Self_Write/LSM/Ylsm/check.go:31: MemTable has 2306869 Nodes, Wal 112 MB, compressing memory
2024/03/02 20:12:42 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 0, index: 0
2024/03/02 20:12:42 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 0 Sstables, the tableSize is 169 MB
2024/03/02 20:12:55 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 1, index: 2
2024/03/02 20:12:55 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 12.9773208s
2024/03/02 20:12:55 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 1 Sstables, the tableSize is 488 MB
2024/03/02 20:13:32 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 2, index: 0
2024/03/02 20:13:32 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 37.2919593s
2024/03/02 20:13:42 E:/Codes/Golang/Self_Write/LSM/Ylsm/check.go:31: MemTable has 2159496 Nodes, Wal 105 MB, compressing memory
2024/03/02 20:13:46 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 0, index: 0
2024/03/02 20:13:46 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:21: compress level 0 Sstables, the tableSize is 158 MB
2024/03/02 20:13:57 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/TableTree.go:171: create a new SsTable, level: 1, index: 0
2024/03/02 20:13:57 E:/Codes/Golang/Self_Write/LSM/Ylsm/ssTable/compaction.go:31: completed compressing, consumption of time 11.1563201s
2024/03/02 20:14:09 E:/Codes/Golang/Self_Write/LSM/Ylsm/demo/main.go:117: error count for queryAbsent: 0
2024/03/02 20:14:09 E:/Codes/Golang/Self_Write/LSM/Ylsm/demo/main.go:31: func insert, time consumption 3m29.2846348s
2024/03/02 20:14:09 E:/Codes/Golang/Self_Write/LSM/Ylsm/demo/main.go:32: func queryAll, time consumption 1m7.0145864s
2024/03/02 20:14:09 E:/Codes/Golang/Self_Write/LSM/Ylsm/demo/main.go:33: func deleteAll, time consumption 5m12.0962699s
2024/03/02 20:14:09 E:/Codes/Golang/Self_Write/LSM/Ylsm/demo/main.go:34: func deleteAbsent, time consumption 0s
2024/03/02 20:14:09 E:/Codes/Golang/Self_Write/LSM/Ylsm/demo/main.go:35: func queryAbsent, time consumption 8.2424732s

```
