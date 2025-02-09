package utils

import (
	"strings"
	"unicode"
)

func ToCamelCase(s string, namingStyle string) string {
	parts := strings.Split(s, "_")
	for i := 0; i < len(parts); i++ {
		// 特殊处理 created_at 和 updated_at
		if parts[i] == "create" && i+1 < len(parts) && parts[i+1] == "at" {
			parts[i] = "Created"
			parts[i+1] = "At"
			i++
			continue
		}
		if parts[i] == "update" && i+1 < len(parts) && parts[i+1] == "at" {
			parts[i] = "Updated"
			parts[i+1] = "At"
			i++
			continue
		}
		parts[i] = strings.Title(parts[i])
	}

	// 根据命名风格调整
	result := strings.Join(parts, "")
	if namingStyle == "camelCase" && len(result) > 0 {
		result = strings.ToLower(string(result[0])) + result[1:]
	}
	return result
}

func ToJSONTag(columnName string) string {
	return strings.ToLower(columnName)
}

func ToSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}
