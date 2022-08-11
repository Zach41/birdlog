package log

import (
	"github.com/Zach41/birdlog/application/model"
	"github.com/Zach41/birdlog/application/reqs"
)

type MetaStore struct {
	collections map[int64]*model.Collection
	indexes     map[int64]*model.Index
	segments    map[int64]*model.Segment
	ops         []reqs.DDLOperation
}

var _ model.Interface = &MetaStore{}

func (m *MetaStore) Recover(tso int64) error {
	end := lowerbound(m.ops, tso)
	if end == -1 {
		end = len(m.ops)
	}
	ops := m.ops[:end]
	for _, op := range ops {
		if err := op.Do(m); err != nil {
			return err
		}
	}
	return nil
}

func (m *MetaStore) CreateCollection(coll *model.Collection) error {
	// TODO
	return nil
}

func (m *MetaStore) DropCollection(collID int64) error {
	// TODO
	return nil
}

func (m *MetaStore) CreateCollectionAlias(collName, alias string) error {
	// TODO
	return nil
}

func (m *MetaStore) DropCollectionAlias(alias string) error {
	// TODO
	return nil
}

func (m *MetaStore) AlterCollectionAlias(collName, alias string) error {
	// TODO
	return nil
}

func (m *MetaStore) CreateIndex(idx *model.Index) error {
	// TODO
	return nil
}

func (m *MetaStore) DropIndex(idxID int64) error {
	// TODO
	return nil
}

func (m *MetaStore) CreatePartition(part *model.Partition) error {
	// TODO
	return nil
}

func (m *MetaStore) DropPartition(partID int64) error {
	// TODO
	return nil
}
