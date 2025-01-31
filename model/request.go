package model

import (
	"github.com/siddontang/go-mysql/schema"
	"sync"
)

var RowRequestPool = sync.Pool{
	New: func() interface{} {
		return new(RowRequest)
	},
}

type RowRequest struct {
	RuleKey   string
	Action    string
	Timestamp uint32
	Old       []interface{}
	Row       []interface{}
	Table     *schema.Table
}

type PosRequest struct {
	Name  string
	Pos   uint32
	Force bool
}

func BuildRowRequest() *RowRequest {
	return RowRequestPool.Get().(*RowRequest)
}

func ReleaseRowRequest(t *RowRequest) {
	RowRequestPool.Put(t)
}
