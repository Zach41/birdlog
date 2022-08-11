package reqs

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/Zach41/birdlog/application/model"
)

type createCollReq struct {
	*base
	meta *model.Collection
}

type dropCollReq struct {
	*base
	collID int64
}

type createCollAliasReq struct {
	*base
	collName string
	alias    string
}

type dropCollAliasReq struct {
	*base
	alias string
}

type alterCollAliasReq struct {
	*base
	collName string
	alias    string
}

func (r *createCollReq) Do(model.Interface) error {
	// TODO
	return nil
}

func (r *createCollReq) Decode(kvpairs map[string]interface{}) error {
	if err := r.base.Decode(kvpairs); err != nil {
		return err
	}
	if v, ok := kvpairs["CollectionMeta"]; ok {
		meta := &model.Collection{}
		encodededString := strings.Trim(v.(string), "\"")
		payload, err := base64.StdEncoding.DecodeString(encodededString)
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

func (r *dropCollReq) Do(model.Interface) error {
	// TODO
	return nil
}

func (r *dropCollReq) Decode(kvpairs map[string]interface{}) error {
	if err := r.base.Decode(kvpairs); err != nil {
		return err
	}
	v, ok := kvpairs["CollectionID"]
	if ok {
		r.collID = v.(int64)
	}
	return nil
}

func (r *createCollAliasReq) Do(metastore model.Interface) error {
	// TODO
	return nil
}

func (r *createCollAliasReq) Decode(kvpairs map[string]interface{}) error {
	if err := r.base.Decode(kvpairs); err != nil {
		return err
	}
	if name, ok := kvpairs["CollectionName"]; ok {
		r.collName = name.(string)
	}
	if alias, ok := kvpairs["Alias"]; ok {
		r.alias = alias.(string)
	}
	return nil
}

func (r *dropCollAliasReq) Do(metastore model.Interface) error {
	// TODO
	return nil
}

func (r *dropCollAliasReq) Decode(kvpairs map[string]interface{}) error {
	if err := r.base.Decode(kvpairs); err != nil {
		return err
	}
	if alias, ok := kvpairs["Alias"]; ok {
		r.alias = alias.(string)
	}
	return nil
}

func (r *alterCollAliasReq) Do(metastore model.Interface) error {
	// TODO
	return nil
}

func (r *alterCollAliasReq) Decode(kvpairs map[string]interface{}) error {
	if err := r.base.Decode(kvpairs); err != nil {
		return err
	}
	if alias, ok := kvpairs["Alias"]; ok {
		r.alias = alias.(string)
	}
	if collName, ok := kvpairs["CollectionName"]; ok {
		r.collName = collName.(string)
	}
	return nil
}
