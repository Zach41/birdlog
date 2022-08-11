package log

import (
	"fmt"
	"testing"
	"time"

	"github.com/Zach41/birdlog/application/reqs"
	"github.com/stretchr/testify/assert"
)

type tso int64

// var _ reqs.DDLOperation =

func (t tso) TSO() int64 {
	return int64(t)
}

func (t tso) Operation() int {
	return reqs.InvalidMetaOperation
}

func (t tso) Time() time.Time {
	return time.Now()
}

func (t tso) Decode(map[string]interface{}) error {
	return nil
}

func TestLowerBound(t *testing.T) {
	list := make([]reqs.DDLOperation, 100)
	for i := 0; i < 100; i++ {
		list[i] = tso(i) * 2
	}
	searchT := func(i int, expected int) {
		txt := fmt.Sprintf("search %d", i)
		t.Run(txt, func(t *testing.T) {
			idx := lowerbound(list, int64(i))
			assert.Equal(t, expected, idx)
		})
	}
	searchT(6, 3)
	searchT(198, 99)
	searchT(0, 0)
	searchT(200, -1)
	searchT(197, 99)
}
