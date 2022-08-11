package log

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParseLine(t *testing.T) {

	// input := `[2022/08/10 10:52:26.392 +08:00] [INFO] [log/meta_logger.go:128] [MetaLogInfo=true] [CollectionMeta="{\"TenantID\":\"\",\"CollectionID\":435185240287739905,\"Partitions\":[{\"PartitionID\":435185240287739906,\"PartitionName\":\"_default\",\"PartitionCreatedTimestamp\":435185240287739906,\"Extra\":null}],\"Name\":\"hello_milvus\",\"Description\":\"hello_milvus is the simplest demo to introduce the APIs\",\"AutoID\":false,\"Fields\":[{\"FieldID\":100,\"Name\":\"ID\",\"IsPrimaryKey\":true,\"Description\":\"\",\"DataType\":5,\"TypeParams\":null,\"IndexParams\":null,\"AutoID\":false},{\"FieldID\":101,\"Name\":\"random\",\"IsPrimaryKey\":false,\"Description\":\"\",\"DataType\":11,\"TypeParams\":null,\"IndexParams\":null,\"AutoID\":false},{\"FieldID\":102,\"Name\":\"embeddings\",\"IsPrimaryKey\":false,\"Description\":\"\",\"DataType\":101,\"TypeParams\":[{\"key\":\"dim\",\"value\":\"8\"}],\"IndexParams\":null,\"AutoID\":false},{\"FieldID\":0,\"Name\":\"RowID\",\"IsPrimaryKey\":false,\"Description\":\"row id\",\"DataType\":5,\"TypeParams\":null,\"IndexParams\":null,\"AutoID\":false},{\"FieldID\":1,\"Name\":\"Timestamp\",\"IsPrimaryKey\":false,\"Description\":\"time stamp\",\"DataType\":5,\"TypeParams\":null,\"IndexParams\":null,\"AutoID\":false}],\"FieldIDToIndexID\":[],\"VirtualChannelNames\":[\"by-dev-rootcoord-dml_0_435185240287739905v0\",\"by-dev-rootcoord-dml_1_435185240287739905v1\"],\"PhysicalChannelNames\":[\"by-dev-rootcoord-dml_0\",\"by-dev-rootcoord-dml_1\"],\"ShardsNum\":2,\"StartPositions\":[{\"key\":\"by-dev-rootcoord-dml_0\",\"data\":\"AQBYVasWCgY=\"},{\"key\":\"by-dev-rootcoord-dml_1\",\"data\":\"AgBYVasWCgY=\"}],\"CreateTime\":435185240287739906,\"ConsistencyLevel\":0,\"Aliases\":null,\"Extra\":null}"] [Operation=0] [TSO=435185240287739906]`
	// kv, err := parseLine(input)
	// assert.Nil(t, err)
	// expectedTime := dateparse.MustParse("2022/08/10 10:52:26.392 +08:00")
	// assert.Equal(t, expectedTime, kv[logTimeKey])
	// assert.Equal(t, 5, len(kv))
	// assert.Equal(t, "true", kv["MetaLogInfo"])
	// assert.Equal(t, int64(435185240287739906), kv["TSO"])
	// assert.Equal(t, int64(CreateCollection), kv["Operation"])
}

func TestReadLogFile(t *testing.T) {
	// filename, err := testTmpFile()
	// assert.Nil(t, err)

	// defer os.Remove(filename)
	// metas, err := readLogFile(filename)
	// assert.Nil(t, err)

	// assert.Equal(t, 1, len(metas))
	// collMeta := metas[0].(*collMeta)
	// assert.Equal(t, "hello_milvus", collMeta.Name)
	// assert.Equal(t, int64(435185865725775169), collMeta.CollectionID)
	// assert.Equal(t, 1, len(collMeta.Partitions))
	// assert.Equal(t, int64(435185865725775170), collMeta.Partitions[0].PartitionID)
	// assert.Equal(t, 5, len(collMeta.Fields))
	// log.Printf("Collection Meta: %#v", collMeta.Collection)
}

func testTmpFile() (string, error) {
	input := `[2022/08/10 11:32:12.274 +08:00] [INFO] [log/meta_logger.go:128] [MetaLogInfo=true] [CollectionMeta="eyJUZW5hbnRJRCI6IiIsIkNvbGxlY3Rpb25JRCI6NDM1MTg1ODY1NzI1Nzc1MTY5LCJQYXJ0aXRpb25zIjpbeyJQYXJ0aXRpb25JRCI6NDM1MTg1ODY1NzI1Nzc1MTcwLCJQYXJ0aXRpb25OYW1lIjoiX2RlZmF1bHQiLCJQYXJ0aXRpb25DcmVhdGVkVGltZXN0YW1wIjo0MzUxODU4NjU3MjU1NzUxNzIsIkV4dHJhIjpudWxsfV0sIk5hbWUiOiJoZWxsb19taWx2dXMiLCJEZXNjcmlwdGlvbiI6ImhlbGxvX21pbHZ1cyBpcyB0aGUgc2ltcGxlc3QgZGVtbyB0byBpbnRyb2R1Y2UgdGhlIEFQSXMiLCJBdXRvSUQiOmZhbHNlLCJGaWVsZHMiOlt7IkZpZWxkSUQiOjEwMCwiTmFtZSI6IklEIiwiSXNQcmltYXJ5S2V5Ijp0cnVlLCJEZXNjcmlwdGlvbiI6IiIsIkRhdGFUeXBlIjo1LCJUeXBlUGFyYW1zIjpudWxsLCJJbmRleFBhcmFtcyI6bnVsbCwiQXV0b0lEIjpmYWxzZX0seyJGaWVsZElEIjoxMDEsIk5hbWUiOiJyYW5kb20iLCJJc1ByaW1hcnlLZXkiOmZhbHNlLCJEZXNjcmlwdGlvbiI6IiIsIkRhdGFUeXBlIjoxMSwiVHlwZVBhcmFtcyI6bnVsbCwiSW5kZXhQYXJhbXMiOm51bGwsIkF1dG9JRCI6ZmFsc2V9LHsiRmllbGRJRCI6MTAyLCJOYW1lIjoiZW1iZWRkaW5ncyIsIklzUHJpbWFyeUtleSI6ZmFsc2UsIkRlc2NyaXB0aW9uIjoiIiwiRGF0YVR5cGUiOjEwMSwiVHlwZVBhcmFtcyI6W3sia2V5IjoiZGltIiwidmFsdWUiOiI4In1dLCJJbmRleFBhcmFtcyI6bnVsbCwiQXV0b0lEIjpmYWxzZX0seyJGaWVsZElEIjowLCJOYW1lIjoiUm93SUQiLCJJc1ByaW1hcnlLZXkiOmZhbHNlLCJEZXNjcmlwdGlvbiI6InJvdyBpZCIsIkRhdGFUeXBlIjo1LCJUeXBlUGFyYW1zIjpudWxsLCJJbmRleFBhcmFtcyI6bnVsbCwiQXV0b0lEIjpmYWxzZX0seyJGaWVsZElEIjoxLCJOYW1lIjoiVGltZXN0YW1wIiwiSXNQcmltYXJ5S2V5IjpmYWxzZSwiRGVzY3JpcHRpb24iOiJ0aW1lIHN0YW1wIiwiRGF0YVR5cGUiOjUsIlR5cGVQYXJhbXMiOm51bGwsIkluZGV4UGFyYW1zIjpudWxsLCJBdXRvSUQiOmZhbHNlfV0sIkZpZWxkSURUb0luZGV4SUQiOltdLCJWaXJ0dWFsQ2hhbm5lbE5hbWVzIjpbImJ5LWRldi1yb290Y29vcmQtZG1sXzBfNDM1MTg1ODY1NzI1Nzc1MTY5djAiLCJieS1kZXYtcm9vdGNvb3JkLWRtbF8xXzQzNTE4NTg2NTcyNTc3NTE2OXYxIl0sIlBoeXNpY2FsQ2hhbm5lbE5hbWVzIjpbImJ5LWRldi1yb290Y29vcmQtZG1sXzAiLCJieS1kZXYtcm9vdGNvb3JkLWRtbF8xIl0sIlNoYXJkc051bSI6MiwiU3RhcnRQb3NpdGlvbnMiOlt7ImtleSI6ImJ5LWRldi1yb290Y29vcmQtZG1sXzAiLCJkYXRhIjoiQVFDUW5Ud1hDZ1k9In0seyJrZXkiOiJieS1kZXYtcm9vdGNvb3JkLWRtbF8xIiwiZGF0YSI6IkFnQ1FuVHdYQ2dZPSJ9XSwiQ3JlYXRlVGltZSI6NDM1MTg1ODY1NzI1NTc1MTcyLCJDb25zaXN0ZW5jeUxldmVsIjowLCJBbGlhc2VzIjpudWxsLCJFeHRyYSI6bnVsbH0="] [Operation=0] [TSO=435185865725575172]`
	tmpDir := os.TempDir()

	file, err := ioutil.TempFile(tmpDir, "ut_log_test")
	defer file.Close()
	if err != nil {
		return "", err
	}
	_, err = file.WriteString(input)
	if err != nil {
		return "", err
	}
	return file.Name(), nil
}
