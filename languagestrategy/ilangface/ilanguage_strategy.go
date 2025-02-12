package ilangface

import "generate_dao/db"

type ILanguageStrategy interface {
	// MapDataType 将统一数据类型映射为具体的语言类型
	MapDataType(unifiedType string) string
	// GetModelTemplateData 获取模板数据
	GetModelTemplateData(fileModel string) string
	// GetFileSuffix 获取文件后缀
	GetFileSuffix() string
	// GetFields 获取字段 columns 列信息 namingStyle 命名风格
	GetFields(columns []db.ColumnInfo, namingStyle string) []interface{}
}
