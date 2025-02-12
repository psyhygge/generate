package core

import (
	"fmt"
	"generate_dao/db"
	"generate_dao/dbstrategy/idbface"
	"generate_dao/languagestrategy/ilangface"
	"generate_dao/utils"
	"os"
	"text/template"
)

type CodeGenerator struct {
	DbStrategy   idbface.IDatabaseStrategy
	LangStrategy ilangface.ILanguageStrategy
}

func (cg *CodeGenerator) Generate(config *Config) error {
	// 获取需要生成的表
	tables, err := cg.DbStrategy.GetTables(config.Output.Tables)
	if err != nil {
		return fmt.Errorf("failed to get tables: %w", err)
	}
	if len(tables) == 0 {
		return fmt.Errorf("no tables found or specified")
	}

	// 生成模型
	for _, table := range tables {
		columns, err := cg.DbStrategy.GetColumns(table)
		if err != nil {
			return fmt.Errorf("failed to get columns for table %s: %w", table, err)
		}
		cg.generateModelFile(config, table, columns)
	}

	return nil
}

type ModelTemplateData struct {
	PackageName string
	StructName  string
	TableName   string
	Fields      []interface{}
}

func getFileName(fileNamingStyle, tableName string) string {
	switch fileNamingStyle {
	case "camelCase":
		return utils.ToCamelCase(tableName, "camelCase")
	case "PascalCase":
		return utils.ToCamelCase(tableName, "PascalCase")
	case "snakeCase":
		return utils.ToSnakeCase(tableName)
	default:
		return tableName
	}
}

func getFileModel(fileModel string) string {
	switch fileModel {
	case "entity":
		return ""
	default:
		return utils.ToCamelCase(fileModel, "PascalCase")
	}
}

func (cg *CodeGenerator) generateModelFile(config *Config, tableName string, columns []db.ColumnInfo) {

	model := cg.LangStrategy.GetModelTemplateData(config.Output.FileModel)
	if model == "" {
		fmt.Println("model is empty")
		return
	}

	data := ModelTemplateData{
		PackageName: config.Output.PackageName,
		StructName:  utils.ToCamelCase(tableName, "PascalCase"),
		TableName:   tableName,
		Fields:      cg.LangStrategy.GetFields(columns, config.Output.NamingStyle),
	}

	os.MkdirAll(config.Output.ModelsDir+"/"+config.Output.FileModel, os.ModePerm)
	fileName := fmt.Sprintf("%s/%s/%s%s%s", config.Output.ModelsDir, config.Output.FileModel, getFileName(config.Output.FileNamingStyle, tableName), getFileModel(config.Output.FileModel), cg.LangStrategy.GetFileSuffix())
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tmpl := template.Must(template.New("model").Parse(model))
	err = tmpl.Execute(file, data)
	if err != nil {
		panic(fmt.Errorf("failed to generate model: %w", err))
	}

	fmt.Printf("Generated model for table: %s -> %s\n", tableName, fileName)
}
