package reqs

import (
	"time"

	"github.com/Zach41/birdlog/application/model"
)

type DDLOperation interface {
	Operation() int
	TSO() int64
	Time() time.Time
	Decode(map[string]interface{}) error
	Do(model.Interface) error
}

const (
	InvalidMetaOperation int = iota - 1
	CreateCollection
	DropCollection
	CreateCollectionAlias
	AlterCollectionAlias
	DropCollectionAlias
	CreatePartition
	DropPartition
	CreateIndex
	DropIndex
	BuildSegmentIndex
)

const (
	LogTimeKey = "ParserLogTime"
)

type base struct {
	op   int
	tso  int64
	logT time.Time
}

var _ DDLOperation = &base{}

func (r *base) Operation() int {
	return r.op
}

func (r *base) TSO() int64 {
	return r.tso
}

func (r *base) Time() time.Time {
	return r.logT
}

func (r *base) Do(model.Interface) error {
	// do nothing
	return nil
}

func (r *base) Decode(kvpairs map[string]interface{}) error {
	r.op = InvalidMetaOperation
	r.tso = -1

	for k, v := range kvpairs {
		switch k {
		case "TSO":
			r.tso = v.(int64)
		case "Operation":
			r.op = v.(int)
		case LogTimeKey:
			r.logT = v.(time.Time)
		}
	}
	return nil
}

func NewRequest(op int) DDLOperation {
	switch op {
	case CreateCollection:
		return &createCollReq{}
	case DropCollection:
		return &dropCollReq{}
	case CreateCollectionAlias:
		return &createCollAliasReq{}
	case AlterCollectionAlias:
		return &alterCollAliasReq{}
	case DropCollectionAlias:
		return &dropCollAliasReq{}
	case CreateIndex:
		return &createIdxReq{}
	case DropIndex:
		return &dropIdxReq{}
	case CreatePartition:
		return &createPartReq{}
	case DropPartition:
		return &dropPartReq{}
	default:
		return nil
	}
}
