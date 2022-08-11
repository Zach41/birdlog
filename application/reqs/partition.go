package reqs

import "github.com/Zach41/birdlog/application/model"

type createPartReq struct {
	*base
	collID   int64
	partID   int64
	partName string
}

type dropPartReq struct {
	*base
	collID   int64
	partName string
}

func (r *createPartReq) Decode(kvpairs map[string]interface{}) error {
	if err := r.base.Decode(kvpairs); err != nil {
		return err
	}
	if collID, ok := kvpairs["CollectionID"]; ok {
		r.collID = collID.(int64)
	}
	if partID, ok := kvpairs["PartitionID"]; ok {
		r.partID = partID.(int64)
	}
	if partName, ok := kvpairs["PartitionName"]; ok {
		r.partName = partName.(string)
	}
	return nil
}

func (r *createPartReq) Do(model.Interface) error {
	// TODO
	return nil
}

func (r *dropPartReq) Decode(kvpairs map[string]interface{}) error {
	if err := r.base.Decode(kvpairs); err != nil {
		return err
	}
	if collID, ok := kvpairs["CollectionID"]; ok {
		r.collID = collID.(int64)
	}
	if partName, ok := kvpairs["PartitionName"]; ok {
		r.partName = partName.(string)
	}
	return nil
}

func (r *dropPartReq) Do(model.Interface) error {
	// TODO
	return nil
}
