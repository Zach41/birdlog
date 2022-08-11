package model

type SegmentState int32

const (
	SegmentState_SegmentStateNone SegmentState = 0
	SegmentState_NotExist         SegmentState = 1
	SegmentState_Growing          SegmentState = 2
	SegmentState_Sealed           SegmentState = 3
	SegmentState_Flushed          SegmentState = 4
	SegmentState_Flushing         SegmentState = 5
	SegmentState_Dropped          SegmentState = 6
	SegmentState_Importing        SegmentState = 7
)

type Segment struct {
	SegmentID           int64
	PartitionID         int64
	NumRows             int64
	MemSize             int64
	DmChannel           string
	CompactionFrom      []int64
	CreatedByCompaction bool
	SegmentState        SegmentState
	IndexInfos          []*SegmentIndex
	ReplicaIds          []int64
	NodeIds             []int64
}

func (s *Segment) Clone() *Segment {
	// TODO
	return &Segment{}
}
