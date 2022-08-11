package log

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Zach41/birdlog/application/model"
	"github.com/Zach41/birdlog/application/reqs"
	"github.com/araddon/dateparse"
)

type MetaStore struct {
	collections map[int64]*model.Collection
	indexes     map[int64]*model.Index
	segments    map[int64]*model.Segment
	ops         []reqs.DDLOperation
}

func Open(filenames ...string) (*MetaStore, error) {
	info := &MetaStore{
		ops: []reqs.DDLOperation{},
	}
	for _, filename := range filenames {
		metaList, err := readLogFile(filename)
		if err != nil {
			return nil, err
		}
		info.ops = append(info.ops, metaList...)
	}
	sort.Slice(info.ops, func(i, j int) bool {
		return info.ops[i].TSO() < info.ops[j].TSO()
	})
	return info, nil
}

func (l *MetaStore) Close() {
	// TODO
}

func (l *MetaStore) Recover(tso int64) error {
	end := lowerbound(l.ops, tso)
	if end == -1 {
		end = len(l.ops)
	}
	ops := l.ops[:end]
	for _, op := range ops {
		switch op.Operation() {
		}
	}
	return nil
}

func readLogFile(filename string) ([]reqs.DDLOperation, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ddlReqs := make([]reqs.DDLOperation, 0)
	reader := bufio.NewScanner(f)
	for reader.Scan() {
		line := reader.Text()
		if skipLine(line) {
			continue
		}
		kvs, err := parseLine(line)
		if err != nil {
			log.Fatalf("parse line failed, %v\nLine: %s", err, line)
		}
		op, ok := kvs["Operation"]
		if !ok {
			continue
		}
		ddlReq := reqs.NewRequest(op.(int))
		if ddlReq == nil {
			return nil, fmt.Errorf("undefined ddl operation %v", op)
		}
		if err := ddlReq.Decode(kvs); err != nil {
			log.Printf("request: %s, err: %v", line, err)
			return nil, fmt.Errorf("parse request failed")
		}
		ddlReqs = append(ddlReqs, ddlReq)
	}
	return ddlReqs, nil
}

func parseLine(line string) (map[string]interface{}, error) {
	tokens := strings.Split(strings.Trim(line, "[]"), "] [")
	timeToken := tokens[0]
	t, err := dateparse.ParseAny(timeToken)
	if err != nil {
		return nil, err
	}

	retkv := make(map[string]interface{})
	retkv[reqs.LogTimeKey] = t
	for _, token := range tokens {
		kv := strings.SplitN(token, "=", 2)
		if len(kv) != 2 || len(kv[0]) <= 0 || len(kv[1]) <= 0 {
			continue
		}
		ival, err := strconv.ParseInt(kv[1], 10, 64)
		if err != nil {
			retkv[kv[0]] = kv[1]
		} else {
			if kv[0] == "Operation" {
				retkv[kv[0]] = int(ival)
			} else {
				retkv[kv[0]] = ival
			}
		}
	}
	return retkv, nil
}

func skipLine(line string) bool {
	return !strings.Contains(line, "[MetaLogInfo=true]")
}
