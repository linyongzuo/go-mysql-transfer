package endpoint

import (
	"bytes"
	"github.com/siddontang/go-mysql/canal"
	"github.com/siddontang/go-mysql/mysql"
	"go-mysql-transfer/model"
)

type MysqlEndpoint struct {
}

func newMysqlEndpoint() *MysqlEndpoint {
	return &MysqlEndpoint{}
}

func (m *MysqlEndpoint) Connect() error {
	return nil
}
func (m *MysqlEndpoint) Ping() error {
	return nil
}
func (m *MysqlEndpoint) Name() string {
	return ""
}
func (m *MysqlEndpoint) Consume(from mysql.Position, rows []*model.RowRequest) error {
	for _, row := range rows {
		// 暂时无视规则，直接写入
		sql := bytes.NewBufferString("")
		switch row.Action {
		case canal.UpdateAction:
		case canal.DeleteAction:
		case canal.InsertAction:
			sql.WriteString("insert into (")
			for index, field := range row.Table.Columns {
				sql.WriteString(field.Name)
				if index != len(row.Table.Columns)-1 {
					sql.WriteString("")
				} else {
					sql.WriteString(") values(")
				}
			}
		}
	}
	return nil
}
func (m *MysqlEndpoint) Stock([]*model.RowRequest) int64 {
	return 0
}
func (m *MysqlEndpoint) Close() {
	return
}
