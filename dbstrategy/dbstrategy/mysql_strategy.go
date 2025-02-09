package dbstrategy

import (
	"fmt"
	"generate_dao/db"
	"generate_dao/dbstrategy/idbface"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLStrategy struct {
	db *gorm.DB
}

func (m *MySQLStrategy) ToUnifiedType(dataType string) string {
	switch dataType {
	case "tinyint", "smallint", "mediumint", "int":
		return "int"
	case "bigint":
		return "int64"
	case "float", "double", "decimal":
		return "float64"
	case "char", "varchar", "text", "tinytext", "mediumtext", "longtext":
		return "string"
	case "date", "datetime", "timestamp", "time":
		return "time"
	case "boolean", "bool":
		return "bool"
	case "json":
		return "json"
	default:
		return "string"
	}
}

func (m *MySQLStrategy) GetTables(specifiedTables []string) ([]string, error) {
	// 如果指定了 tables，则直接返回
	if len(specifiedTables) > 0 {
		return specifiedTables, nil
	}
	var dbName string
	m.db.Raw("SELECT DATABASE()").Scan(&dbName)
	// 否则返回所有表
	var tables []string
	m.db.Raw(`
		SELECT TABLE_NAME 
		FROM INFORMATION_SCHEMA.TABLES 
		WHERE TABLE_SCHEMA = ?
	`, dbName).Scan(&tables)
	return tables, nil
}

func (m *MySQLStrategy) GetColumns(tableName string) ([]db.ColumnInfo, error) {
	var dbName string
	m.db.Raw("SELECT DATABASE()").Scan(&dbName)

	var columns []db.ColumnInfo
	m.db.Raw(`
		SELECT 
			COLUMN_NAME,
			DATA_TYPE,
			COLUMN_KEY,
			IS_NULLABLE,
			COLUMN_COMMENT
		FROM INFORMATION_SCHEMA.COLUMNS 
		WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?
		ORDER BY ORDINAL_POSITION
	`, dbName, tableName).Scan(&columns)
	return columns, nil
}

func NewMySQLStrategy(dsn string) idbface.IDatabaseStrategy {
	var dialector gorm.Dialector
	dialector = mysql.Open(dsn)

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect dbstrategy: %w", err))
	}

	return &MySQLStrategy{db: db}
}
