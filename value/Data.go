package value

type SearchResult int

const (
	None SearchResult = iota
	Deleted
	Success
)

type Data struct {
	Key     string
	Value   []byte
	Deleted bool
}

func (d *Data) Copy() *Data {
	return &Data{
		Key:     d.Key,
		Value:   d.Value,
		Deleted: d.Deleted,
	}
}
