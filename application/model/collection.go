package model

type Collection struct {
	TenantID             string
	CollectionID         int64
	Partitions           []*Partition
	Name                 string
	Description          string
	AutoID               bool
	Fields               []*Field
	FieldIDToIndexID     []Int64Tuple
	VirtualChannelNames  []string
	PhysicalChannelNames []string
	ShardsNum            int32
	StartPositions       []*KeyDataPair
	CreateTime           uint64
	ConsistencyLevel     ConsistencyLevel
	Aliases              []string
	Extra                map[string]string // extra kvs
}

func (c *Collection) Clone() *Collection {
	// TODO
	return &Collection{}
}
