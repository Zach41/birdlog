package reqs

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/Zach41/birdlog/application/model"
)

type createIdxReq struct {
	*base
	meta *model.Index
}

type dropIdxReq struct {
	*base
	collName  string
	fieldName string
	indexName string
}

func (r *createIdxReq) Decode(kvpairs map[string]interface{}) error {
	if err := r.base.Decode(kvpairs); err != nil {
		return err
	}
	if v, ok := kvpairs["IndexMeta"]; ok {
		meta := &model.Index{}
		encodedStirng := strings.Trim(v.(string), "\"")
		payload, err := base64.StdEncoding.DecodeString(encodedStirng)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(payload, meta); err != nil {
			return err
		}
		r.meta = meta
	}
	return nil
}

func (r *createIdxReq) Do(model.Interface) error {
	// TODO
	return nil
}

func (r *dropIdxReq) Decode(kvpairs map[string]interface{}) error {
	if err := r.base.Decode(kvpairs); err != nil {
		return err
	}
	if collName, ok := kvpairs["CollectionName"]; ok {
		r.collName = collName.(string)
	}
	if fName, ok := kvpairs["FieldName"]; ok {
		r.fieldName = fName.(string)
	}
	if idxName, ok := kvpairs["IndexName"]; ok {
		r.indexName = idxName.(string)
	}
	return nil
}

func (r *dropIdxReq) Do(model.Interface) error {
	return nil
}
