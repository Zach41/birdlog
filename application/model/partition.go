package model

type Partition struct {
	PartitionID               int64
	PartitionName             string
	PartitionCreatedTimestamp uint64
	Extra                     map[string]string
}

func (p *Partition) Clone() *Partition {
	// TODO
	return &Partition{}
}
