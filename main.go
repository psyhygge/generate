package main

import (
	"fmt"
	"generate_dao/core"
	"generate_dao/dbstrategy/dbstrategy"
	"generate_dao/languagestrategy/langstrategy"
)

func main() {
	// 加载配置
	config := core.LoadConfig()
	// 创建数据库策略
	dbFactory := &dbstrategy.DatabaseStrategyFactory{}
	dbStrategy, err := dbFactory.CreateStrategy(config.Database.Type, config.Database.DSN)
	if err != nil {
		panic(err)
	}

	// 创建语言策略
	langFactory := &langstrategy.LanguageStrategyFactory{}
	langStrategy, err := langFactory.CreateStrategy(config.Output.Language)
	if err != nil {
		panic(err)
	}

	// 创建代码生成器
	generator := &core.CodeGenerator{
		DbStrategy:   dbStrategy,
		LangStrategy: langStrategy,
	}

	// 生成代码
	if err := generator.Generate(config); err != nil {
		panic(fmt.Errorf("failed to generate code: %w", err))
	}

	fmt.Println("Code generation completed!")
}
