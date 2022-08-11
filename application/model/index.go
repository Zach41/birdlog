package model

type SegmentIndex struct {
	Segment
	EnableIndex    bool
	CreateTime     uint64
	BuildID        int64
	IndexSize      uint64
	IndexFilePaths []string
}

type Index struct {
	CollectionID   int64
	FieldID        int64
	IndexID        int64
	IndexName      string
	IsDeleted      bool
	CreateTime     uint64
	IndexParams    []*KeyValuePair
	SegmentIndexes map[int64]SegmentIndex //segmentID -> segmentIndex
	Extra          map[string]string
}

func (i *Index) Clone() *Index {
	// TODO
	return &Index{}
}
