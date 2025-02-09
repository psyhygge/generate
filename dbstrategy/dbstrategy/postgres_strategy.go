package dbstrategy

import (
	"generate_dao/db"
	"generate_dao/dbstrategy/idbface"
)

type PostgresStrategy struct {
}

func (s *PostgresStrategy) ToUnifiedType(dataType string) string {
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

func (s *PostgresStrategy) GetTables(specifiedTables []string) ([]string, error) {

	return nil, nil
}

func (s *PostgresStrategy) GetColumns(tableName string) ([]db.ColumnInfo, error) {

	return nil, nil
}

func NewPostgresStrategy() idbface.IDatabaseStrategy {
	return &PostgresStrategy{}
}
