package idbface

import (
	"generate_dao/db"
)

type IDatabaseStrategy interface {
	// GetTables 获取数据库表名
	GetTables(specifiedTables []string) ([]string, error)
	// GetColumns 获取数据库表字段信息
	GetColumns(tableName string) ([]db.ColumnInfo, error)
	// ToUnifiedType 将数据库字段类型转换为统一类型
	ToUnifiedType(dataType string) string
}
