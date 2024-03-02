package memTable

import "Ylsm/value"

type MemTable interface {
	GetCount() int
	Search(key string) (value.Data, value.SearchResult)
	Set(key string, values []byte) (oldValue value.Data, hasOld bool)
	Delete(key string) (oldValue value.Data, hasOld bool)
	GetValues() (values []value.Data)
	Swap() MemTable
}
