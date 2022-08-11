package log

import (
	"github.com/Zach41/birdlog/application/model"
	"github.com/Zach41/birdlog/application/reqs"
)

func findCollectionByName(store map[int64]interface{}, name string) *model.Collection {
	for _, v := range store {
		meta, ok := v.(*model.Collection)
		if !ok {
			continue
		}
		if meta.Name == name {
			return meta
		}
	}
	return nil
}

func rmElements[T comparable](list []T, es []T) []T {
	ret := make([]T, 0)
	rmset := make(map[T]struct{})
	for _, e := range es {
		rmset[e] = struct{}{}
	}
	for _, v := range list {
		if _, ok := rmset[v]; !ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func lowerbound(ops []reqs.DDLOperation, targetTso int64) int {
	l, r := 0, len(ops)-1
	idx := -1
	for l <= r {
		mid := (l + r) / 2
		if ops[mid].TSO() >= targetTso {
			idx = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return idx
}
