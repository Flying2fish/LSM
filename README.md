
测试
本测试主要测试qlsm的增删改查功能性与效率，以及数据库启动时 WAL 和 SsTable 的载入功能性与效率。

实验环境
CPU: Intel(R) Xeon(R) Gold 5218R CPU @ 2.10GHz
操作系统: Ubuntu20.04
内存大小: 128GB
测试1
流程
insert() 插入11881376(26的5次方)条数据
queryAll() 查询11881376(26的5次方)条存在的数据
deleteAll() 删除11881376(26的5次方)条存在的数据
deleteAbsent() 删除1条不存在的数据
queryAbsent() 查询11881376(26的5次方)条不存在的数据
程序退出
insert() 插入11881376(26的5次方)条数据
queryAll() 查询11881376(26的5次方)条存在的数据
deleteAll() 删除11881376(26的5次方)条存在的数据
deleteAbsent() 删除1条不存在的数据
queryAbsent() 查询11881376(26的5次方)条不存在的数据
结果1
...
2023/06/28 23:03:10 /home/qian/todo/qlsm/demo/main.go:62: insert count: 11881376
2023/06/28 23:04:03 /home/qian/todo/qlsm/demo/main.go:90: error count for queryAll: 0
...
2023/06/28 23:08:08 /home/qian/todo/qlsm/demo/main.go:117: error count for queryAbsent: 0
2023/06/28 23:08:08 /home/qian/todo/qlsm/demo/main.go:31: func insert, time consumption 6m48.954061763s
2023/06/28 23:08:08 /home/qian/todo/qlsm/demo/main.go:32: func queryAll, time consumption 52.912864036s
2023/06/28 23:08:08 /home/qian/todo/qlsm/demo/main.go:33: func deleteAll, time consumption 3m33.146243216s
2023/06/28 23:08:08 /home/qian/todo/qlsm/demo/main.go:34: func deleteAbsent, time consumption 7.404µs
2023/06/28 23:08:08 /home/qian/todo/qlsm/demo/main.go:35: func queryAbsent, time consumption 31.21362344s
[CRASH]
2023/06/28 23:12:43 /home/qian/todo/qlsm/qlsm.go:29: load the configuration...
2023/06/28 23:12:43 /home/qian/todo/qlsm/qlsm.go:32: initialize DB...
2023/06/28 23:12:43 /home/qian/todo/qlsm/qlsm.go:53: load Wal, recover MemTable...
2023/06/28 23:12:43 /home/qian/todo/qlsm/wal/Wal.go:36: load the wal.log, consumption of time: 48.38µs
2023/06/28 23:12:43 /home/qian/todo/qlsm/qlsm.go:56: load DB...
2023/06/28 23:12:47 /home/qian/todo/qlsm/ssTable/TableTree.go:105: load the ../lsmDB/1.0.db, consumption of time: 3.579540708s
2023/06/28 23:12:56 /home/qian/todo/qlsm/ssTable/TableTree.go:105: load the ../lsmDB/2.0.db, consumption of time: 9.539435985s
2023/06/28 23:13:04 /home/qian/todo/qlsm/ssTable/TableTree.go:105: load the ../lsmDB/2.1.db, consumption of time: 7.807468406s
2023/06/28 23:13:25 /home/qian/todo/qlsm/ssTable/TableTree.go:105: load the ../lsmDB/3.0.db, consumption of time: 20.764075196s
2023/06/28 23:13:25 /home/qian/todo/qlsm/ssTable/TableTree.go:80: load the TablesTree , consumption of time: 41.690846983s
...
2023/06/28 23:20:45 /home/qian/todo/qlsm/demo/main.go:62: insert count: 11881376
2023/06/28 23:21:57 /home/qian/todo/qlsm/demo/main.go:90: error count for queryAll: 0
...
2023/06/28 23:28:34 /home/qian/todo/qlsm/demo/main.go:117: error count for queryAbsent: 0
2023/06/28 23:28:34 /home/qian/todo/qlsm/demo/main.go:31: func insert, time consumption 7m20.617307079s
2023/06/28 23:28:34 /home/qian/todo/qlsm/demo/main.go:32: func queryAll, time consumption 1m11.424989818s
2023/06/28 23:28:34 /home/qian/todo/qlsm/demo/main.go:33: func deleteAll, time consumption 6m27.354441512s
2023/06/28 23:28:34 /home/qian/todo/qlsm/demo/main.go:34: func deleteAbsent, time consumption 7.034µs
2023/06/28 23:28:34 /home/qian/todo/qlsm/demo/main.go:35: func queryAbsent, time consumption 9.54136329s
测试2
流程
insert() 插入11881376(26的5次方)条数据
程序退出
queryAll() 查询11881376(26的5次方)条存在的数据
deleteAll() 删除11881376(26的5次方)条存在的数据
程序退出
queryAbsent() 查询11881376(26的5次方)条不存在的数据
结果2
...
2023/06/28 23:54:32 /home/qian/todo/qlsm/demo/main.go:62: insert count: 11881376
2023/06/28 23:54:32 /home/qian/todo/qlsm/demo/main.go:31: func insert, time consumption 6m55.616976789s
[CRASH]
2023/06/28 23:55:29 /home/qian/todo/qlsm/qlsm.go:29: load the configuration...
2023/06/28 23:55:29 /home/qian/todo/qlsm/qlsm.go:32: initialize DB...
2023/06/28 23:55:29 /home/qian/todo/qlsm/qlsm.go:53: load Wal, recover MemTable...
2023/06/28 23:55:30 /home/qian/todo/qlsm/wal/Wal.go:36: load the wal.log, consumption of time: 386.488948ms
2023/06/28 23:55:30 /home/qian/todo/qlsm/qlsm.go:56: load DB...
2023/06/28 23:55:31 /home/qian/todo/qlsm/ssTable/TableTree.go:105: load the ../lsmDB/1.0.db, consumption of time: 1.422849222s
2023/06/28 23:55:52 /home/qian/todo/qlsm/ssTable/TableTree.go:105: load the ../lsmDB/3.0.db, consumption of time: 21.087827242s
2023/06/28 23:55:52 /home/qian/todo/qlsm/ssTable/TableTree.go:80: load the TablesTree , consumption of time: 22.510812842s
2023/06/28 23:55:52 /home/qian/todo/qlsm/qlsm.go:35: start checking in the background...
...
2023/06/28 23:56:57 /home/qian/todo/qlsm/demo/main.go:90: error count for queryAll: 0
2023/06/29 00:00:38 /home/qian/todo/qlsm/demo/main.go:32: func queryAll, time consumption 1m4.720664143s
2023/06/29 00:00:38 /home/qian/todo/qlsm/demo/main.go:33: func deleteAll, time consumption 3m41.399412966s
[CRASH]
2023/06/29 00:03:36 /home/qian/todo/qlsm/qlsm.go:29: load the configuration...
2023/06/29 00:03:36 /home/qian/todo/qlsm/qlsm.go:32: initialize DB...
2023/06/29 00:03:36 /home/qian/todo/qlsm/qlsm.go:53: load Wal, recover MemTable...
2023/06/29 00:03:38 /home/qian/todo/qlsm/wal/Wal.go:36: load the wal.log, consumption of time: 2.204609131s
2023/06/29 00:03:38 /home/qian/todo/qlsm/qlsm.go:56: load DB...
2023/06/29 00:03:46 /home/qian/todo/qlsm/ssTable/TableTree.go:105: load the ../lsmDB/2.0.db, consumption of time: 8.601862245s
2023/06/29 00:03:58 /home/qian/todo/qlsm/ssTable/TableTree.go:105: load the ../lsmDB/2.1.db, consumption of time: 11.291855699s
2023/06/29 00:04:17 /home/qian/todo/qlsm/ssTable/TableTree.go:105: load the ../lsmDB/3.0.db, consumption of time: 19.305720113s
2023/06/29 00:04:17 /home/qian/todo/qlsm/ssTable/TableTree.go:80: load the TablesTree , consumption of time: 39.199662911s
2023/06/29 00:04:17 /home/qian/todo/qlsm/qlsm.go:35: start checking in the background...
2023/06/29 00:04:26 /home/qian/todo/qlsm/demo/main.go:117: error count for queryAbsent: 0
2023/06/29 00:04:26 /home/qian/todo/qlsm/demo/main.go:35: func queryAbsent, time consumption 8.988589246s
总体架构
qlsm 组件可以分为内存组件和磁盘组件两大类。其中常驻内存的有 MemTable, TableTree 和 SsTable 中的索引信息；而常驻磁盘的有 WAL, SsTable的数据部分。 qlsm 主要支持的操作是数据的增删改查操作，同时，还有一个监控协程对 MemTable 中节点数目，WAL 文件大小，各层 SsTable 数量与总大小进行监控，当各指标到达一定阈值时，则会触发相应的操作，如数据落盘，文件压实等操作。

当发生插入，删除或修改操作时，数据会写入 MemTable 中，同时，为了防止崩溃导致数据丢失，数据也会写入 WAL 文件中。值得说明的是，在 qlsm 中这两步并不是原子操作。

当 MemTable 中节点数目过多或者 WAL 文件大小过大时，将会触发落盘操作，将 MemTable 落盘为 SsTable，并将 MemTable 对应的 WAL 文件删去。

当 SsTable 过大或者数目过多的时候，就会触发压实操作，将低层次的小容量 SsTable 压实为高层次大容量的 SsTable

当发生查询操作时，qlsm 会先查询 MemTable，如果 MemTable 不命中，则会逐层倒序查找 SsTable。值得说明的是，在 qlsm 中，各层的 SsTable 依然会存在键重叠的情况，因此需要倒序查找每一个 SsTable 来确保查找到最新数据。一般 lsm 实现在 level > 0 的时候不会出现 key 重叠情况，qlsm 则不同。

qlsm 基本配置
qlsm 的配置主要影响监控协程的检查操作。

DataDir 数据库目录

Level0Size 0 层 SsTable 文件总大小 (MB)

i + 1 层 SsTable 文件总大小 = i 层 SsTable 文件总大小 << 2

PartSize 每层 SsTable 数量的最大值

MemTable 中节点最大数量

CheckInterval 监控协程检查的时间间隔 (ms)

基本组件
接下来介绍qlsm的基本组件的一些关键介绍。

KV
KV 是数据的抽象，包含数据的键、值以及状态。

type Data struct {
	Key     string
	Value   []byte
	Deleted bool
}

// Copy 返回 Data 的一个复制
func (d *Data) Copy() *Data
MemTable
MemTable 是内存表的抽象，qlsm 实现了两种数据结构支持 MemTable，分别是 BST 和跳表，BST 为早期版本。

type MemTable interface {
	GetCount() int
	Search(key string) (kv.Data, kv.SearchResult)
	Set(key string, value []byte) (oldValue kv.Data, hasOld bool)
	Delete(key string) (oldValue kv.Data, hasOld bool)
	GetValues() (values []kv.Data)
	Swap() MemTable
}
Write Ahead Log
Write Ahead Log 是内存表在磁盘的映射，用于在崩溃后对数据库进行恢复。每次数据库启动时，会读取 Write Ahead Log 并生成对应的 MemTable

type Wal struct {
	f    *os.File
	path string
	sync.Mutex
}

// Load 通过 wal.log 文件初始化 WAL, 并生成对应的 MemTable
func (w *Wal) Load(dir string) memTable.MemTable
// Write 将数组增加、修改和删除操作写入 WAL
func (w *Wal) Write(value kv.Data)
SsTable
MemTable 的节点数目或 WAL 大小达到阈值时会将 MemTable 落盘为 SsTable，值得一提的是 SsTable 的 sparseIndex 常驻内存

type SsTable struct {
	f           *os.File            //文件句柄
	filepath    string              // SsTable 文件路径
	metaInfo    MetaInfo            // SsTable 元数据
	sparseIndex map[string]Position // 文件的稀疏索引列表
	sync.Mutex
}

// MetaInfo 是 SsTable 的元数据, 存储在文件的末尾
type MetaInfo struct {
	version    int64 // 版本号
	dataStart  int64 // 数据区起始索引
	dataLen    int64 // 数据区长度
	indexStart int64 // 稀疏索引区起始索引
	indexLen   int64 // 稀疏索引区长度
}

// Position 存储在 SparseIndex 中, 表示 KV 的起始位置和长度
type Position struct {
	Start   int64 // 起始位置
	Len     int64 // 长度
	Deleted bool  //删除标志
}

// Search 先通过常驻内存 sparseIndex 找到 Position, 再从磁盘数据区加载数据
func (t *SsTable) Search(key string) (value kv.Data, result kv.SearchResult)
TablesTree
TablesTree 用于管理 各层 SsTable 文件

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
// Search 从所有 SsTable 表中查找数据
func (tt *TablesTree) Search(key string) (kv.Data, kv.SearchResult)
// Insert 在 TablesTree 的 level 层的末尾插入 SsTable, 并返回新插入的 SsTable 的 index
func (tt *TablesTree) Insert(t *SsTable, level int) (index int)
// CreateTable 为对应层生成 SsTable
func (tt *TablesTree) CreateTable(values []kv.Data, level int) *SsTable
监控协程
// checkMemory 将 MemTable 落库成 SsTable [ MemTable节点数 >= Threshold 或 WAL >= Level0Size ]
func checkMemory()
// Compaction 对 SsTable 进行压实 [db 文件数量 > PartSize 或者 db 文件总大小 > levelMaxSize]
func (tt *TablesTree) Compaction()
存在问题
Compaction 合并策略问题
目前 qlsm 在进行 compaction 操作时，需要将对应层的 SsTable 全部读入内存，形成一个新的且内存耗费较大的 MemTable，在 16G 内存的机器上很容易造成内存溢出。后续可以探索如何利用硬盘来进行 compaction 操作

SsTable 键重叠问题
目前 qlsm 中高层 SsTable 依然会出现键重叠情况，因此在 TablesTree 的 Search 方法并不能使用二分来直接找到键，而是需要倒序查找该层的每一个 SsTable，当出现 Absence 情况时，则需要查找完所有的 SsTable，造成极大的时间开销。除此之外，布隆过滤器也是应对queryAbsent 的一个有效途径，但是 qlsm 也没有实现

锁粒度大问题
目前 qlsm 通过比较粗暴的全局锁来保证业务操作和落盘、压实操作不会并发进行来保证不会查找到旧数据。
