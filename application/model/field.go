package model

type DataType int32

const (
	DataType_None         DataType = 0
	DataType_Bool         DataType = 1
	DataType_Int8         DataType = 2
	DataType_Int16        DataType = 3
	DataType_Int32        DataType = 4
	DataType_Int64        DataType = 5
	DataType_Float        DataType = 10
	DataType_Double       DataType = 11
	DataType_String       DataType = 20
	DataType_VarChar      DataType = 21
	DataType_BinaryVector DataType = 100
	DataType_FloatVector  DataType = 101
)

type KeyValuePair struct {
	Key   string
	Value string
}

type Int64Tuple struct {
	Key, Value int64
}

type KeyDataPair struct {
	Key  string
	Data []byte
}

type ConsistencyLevel int32

const (
	ConsistencyLevel_Strong     ConsistencyLevel = 0
	ConsistencyLevel_Session    ConsistencyLevel = 1
	ConsistencyLevel_Bounded    ConsistencyLevel = 2
	ConsistencyLevel_Eventually ConsistencyLevel = 3
	ConsistencyLevel_Customized ConsistencyLevel = 4
)

type Field struct {
	FieldID      int64
	Name         string
	IsPrimaryKey bool
	Description  string
	DataType     DataType
	TypeParams   []*KeyValuePair
	IndexParams  []*KeyValuePair
	AutoID       bool
}

func (f *Field) Clone() *Field {
	// TODO
	return &Field{}
}
