package db

type ColumnInfo struct {
	ColumnName    string `gorm:"column:COLUMN_NAME"`
	DataType      string `gorm:"column:DATA_TYPE"`
	ColumnKey     string `gorm:"column:COLUMN_KEY"`
	IsNullable    string `gorm:"column:IS_NULLABLE"`
	ColumnComment string `gorm:"column:COLUMN_COMMENT"`
}
