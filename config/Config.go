package config

import "sync"

type Config struct {
	DataDir       string // 数据目录
	Level0Size    int    // 0 层所有 SsTable 文件大小总和的最大值 (MB)
	PartSize      int    // 每层 SsTable 数量的最大值
	Threshold     int    // MemTable 中 kv 最大数量
	CheckInterval int    // 监控协程检查的时间间隔 (ms)
}

var config Config
var once sync.Once

func Init(cfg Config) {
	once.Do(func() {
		config = cfg
	})
}

func GetConfig() Config {
	return config
}
