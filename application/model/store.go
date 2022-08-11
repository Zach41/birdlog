package model

type Interface interface {
	CreateCollection(coll *Collection) error
	DropCollection(collID int64) error
	CreateCollectionAlias(collName, alias string) error
	DropCollectionAlias(alias string) error
	AlterCollectionAlias(collName, alias string) error
	CreateIndex(idx *Index) error
	DropIndex(idxID int64) error
	CreatePartition(part *Partition) error
	DropPartition(partID int64) error
	// AddSegment(seg *Segment) error
}
