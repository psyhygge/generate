package langstrategy

import (
	"generate_dao/db"
	"generate_dao/dbstrategy"
	"generate_dao/languagestrategy/ilangface"
	"generate_dao/utils"
)

type JavaStrategy struct {
}

func (j *JavaStrategy) GetFileSuffix() string {
	return ".java"
}

func (j *JavaStrategy) MapDataType(unifiedType string) string {
	switch unifiedType {
	case "int":
		return "int"
	case "int64":
		return "long"
	case "float64":
		return "double"
	case "string":
		return "String"
	case "time":
		return "java.util.Date"
	case "bool":
		return "boolean"
	case "json":
		return "String"
	default:
		return "String"
	}
}

func (j *JavaStrategy) GetModelTemplateData() string {
	return dbstrategy.JavaModelTemplate
}

func (j *JavaStrategy) GetFields(columns []db.ColumnInfo, namingStyle string) []interface{} {
	var interfaceFields []interface{}

	for _, col := range columns {
		// 直接存储为 interface{}
		interfaceFields = append(interfaceFields, map[string]interface{}{
			"Name":    utils.ToCamelCase(col.ColumnName, namingStyle),
			"Type":    j.MapDataType(col.DataType),
			"Comment": col.ColumnComment,
		})
	}

	return interfaceFields
}

func NewJavaStrategy() ilangface.ILanguageStrategy {
	return &JavaStrategy{}
}
