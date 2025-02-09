package langstrategy

import (
	"generate_dao/db"
	"generate_dao/dbstrategy"
	"generate_dao/languagestrategy/ilangface"
	"generate_dao/utils"
	"strings"
)

type GoStrategy struct {
}

type Field struct {
	Name       string
	Type       string
	JSONTag    string
	GormTag    string
	Comment    string
	IsPrimary  bool
	IsNullable bool
}

func buildGormTag(col db.ColumnInfo) string {
	tags := []string{"column:" + col.ColumnName}
	if col.ColumnKey == "PRI" {
		tags = append(tags, "primary_key")
	}
	if col.IsNullable == "NO" {
		tags = append(tags, "not null")
	}
	return strings.Join(tags, ";")
}

func (g *GoStrategy) GetFields(columns []db.ColumnInfo, namingStyle string) []interface{} {
	var interfaceFields []interface{}

	for _, col := range columns {
		// 直接存储为 interface{}
		interfaceFields = append(interfaceFields, map[string]interface{}{
			"Name":       utils.ToCamelCase(col.ColumnName, namingStyle),
			"Type":       g.MapDataType(col.DataType),
			"JSONTag":    utils.ToJSONTag(col.ColumnName),
			"GormTag":    buildGormTag(col),
			"Comment":    col.ColumnComment,
			"IsPrimary":  col.ColumnKey == "PRI",
			"IsNullable": col.IsNullable == "YES",
		})
	}

	return interfaceFields
}

func (g *GoStrategy) GetFileSuffix() string {
	return ".go"
}

func (g *GoStrategy) MapDataType(unifiedType string) string {
	switch unifiedType {
	case "int":
		return "int"
	case "int64":
		return "int64"
	case "float64":
		return "float64"
	case "string":
		return "string"
	case "time":
		return "time.Time"
	case "bool":
		return "bool"
	case "json":
		return "datatypes.JSON"
	default:
		return "string"
	}
}

func (g *GoStrategy) GetModelTemplateData() string {
	return dbstrategy.GoModelTemplate
}

func NewGoStrategy() ilangface.ILanguageStrategy {
	return &GoStrategy{}
}
