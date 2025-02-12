package main

import (
	"database/sql"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"generate_dao/core"
	"generate_dao/dbstrategy/dbstrategy"
	"generate_dao/languagestrategy/langstrategy"
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// 创建 Fyne 应用
	myApp := app.New()
	myWindow := myApp.NewWindow("Database Connection Test")

	// 配置结构体
	var config core.Config

	// 1. DSN 输入框
	dsnEntry := widget.NewEntry()
	dsnEntry.SetPlaceHolder("Enter DSN here...")

	// 2. 数据库类型选择
	dbTypeSelect := widget.NewSelect([]string{"mysql", "postgreSQL", "SQLite"}, func(value string) {
		config.Database.Type = value
		// 根据选择的数据库类型，自动生成示例 DSN
		switch value {
		case "mysql":
			dsnEntry.SetText("root:password@tcp(localhost:3306)/your_database?charset=utf8mb4&parseTime=True&loc=Local")
		case "postgres":
			dsnEntry.SetText("user=postgres password=yourpassword dbname=yourdatabase sslmode=disable")
		case "SQLite":
			dsnEntry.SetText("your_database.db")
		}
	})

	// 3. 确认按钮，点击后尝试连接数据库
	confirmButton := widget.NewButton("Confirm", func() {
		// 获取用户输入的 DSN
		dsn := dsnEntry.Text

		// 根据选中的数据库类型进行连接
		var db *sql.DB
		var err error
		switch config.Database.Type {
		case "mysql":
			db, err = sql.Open("mysql", dsn)
		case "postgres":
			db, err = sql.Open("postgres", dsn)
		case "SQLite":
			db, err = sql.Open("sqlite3", dsn)
		default:
			log.Println("Invalid database type selected.")
			err = fmt.Errorf("invalid database type selected")
		}

		// 检查数据库连接是否成功
		if err != nil {
			// 连接失败，提示错误
			log.Println("Error connecting to database:", err)
			widget.NewLabel("Connection failed, please check your DSN and try again.")
			dialog.ShowError(err, myWindow)
			return
		}

		// 尝试连接数据库
		if err := db.Ping(); err != nil {
			// 连接失败，提示错误
			log.Println("Ping failed:", err)
			widget.NewLabel("Connection failed, please check your DSN and try again.")
			dialog.ShowError(err, myWindow)
			return
		}

		// 连接成功，显示成功消息并进入下一步
		log.Println("Database connected successfully!")
		config.Database.DSN = dsn
		// 创建数据库策略
		dbFactory := &dbstrategy.DatabaseStrategyFactory{}
		dbStrategy, err := dbFactory.CreateStrategy(config.Database.Type, config.Database.DSN)
		if err != nil {
			panic(err)
		}

		// 3. 语言选择
		langSelect := widget.NewSelect([]string{"go", "java"}, func(value string) {
			config.Output.Language = value
		})

		// 4. 包名输入框
		packageNameEntry := widget.NewEntry()
		packageNameEntry.SetPlaceHolder("Enter Package Name...")
		packageNameEntry.OnChanged = func(s string) {
			config.Output.PackageName = s
		}

		// 5. 文件模型选择 (Mapper 或 Entity)
		fileModelSelect := widget.NewSelect([]string{"entity", "mapper"}, func(value string) {
			config.Output.FileModel = value
		})

		// 6. ModelsDir 输入框
		modelsDirEntry := widget.NewEntry()
		modelsDirEntry.SetPlaceHolder("Enter Models Directory...")
		modelsDirEntry.OnChanged = func(s string) {
			config.Output.ModelsDir = s
		}

		// 7. 命名风格选择 (PascalCase 或 camelCase)
		namingStyleSelect := widget.NewSelect([]string{"PascalCase", "camelCase"}, func(value string) {
			config.Output.NamingStyle = value
		})

		// 8. 文件命名风格选择 (PascalCase, camelCase, snake_case)
		fileNamingStyleSelect := widget.NewSelect([]string{"PascalCase", "camelCase", "snake_case"}, func(value string) {
			config.Output.FileNamingStyle = value
		})

		// 9. 表选择框 (假设从数据库中获取的表)
		// 示例：提供一个简单的表格选择框，实际中可以通过数据库查询动态加载
		tables, _ := dbStrategy.GetTables(nil)
		tablesSelect := widget.NewCheckGroup(tables, func(value []string) {
			config.Output.Tables = value
		})

		// 10. 确认按钮，点击后生成代码
		confirmButton2 := widget.NewButton("Generate Code", func() {
			// 生成代码逻辑（调用后端 API 或本地生成）
			log.Printf("Config: %+v\n", config)
			// TODO: 调用后端生成代码的逻辑
			fmt.Println("Code generation started...")

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

			// 创建执行器
			executer := &core.Executer{Generator: generator}

			// 生成代码
			if err := executer.Generator.Generate(&config); err != nil {
				panic(fmt.Errorf("failed to generate code: %w", err))
			}

			fmt.Println("Code generation completed!")
		})

		myWindow.SetContent(container.NewVBox(
			widget.NewLabel("Database connected successfully!"),
			// 进入下一步的逻辑，显示后续的配置项
			widget.NewLabel("选择编程语言(Select Language):"),
			langSelect,
			widget.NewLabel("输入包名(Enter Package Name):"),
			packageNameEntry,
			widget.NewLabel("选择生成文件模型(Select File Model):"),
			fileModelSelect,
			widget.NewLabel("输入生成文件目录(Enter Models Directory):"),
			modelsDirEntry,
			widget.NewLabel("选择类名命名方式(Select Naming Style):"),
			namingStyleSelect,
			widget.NewLabel("选择文件命名方式(Select File Naming Style):"),
			fileNamingStyleSelect,
			widget.NewLabel("勾选需要生成的表(Select Tables):"),
			tablesSelect,
			confirmButton2,
		))
	})

	// 布局，显示数据库类型选择和 DSN 输入框
	content := container.NewVBox(
		widget.NewLabel("选择数据库类型(Select Database Type):"),
		dbTypeSelect,
		widget.NewLabel("输入DSN连接(Enter DSN):"),
		dsnEntry,
		confirmButton,
	)

	// 设置窗口内容
	myWindow.SetContent(content)

	// 设置窗口大小并运行
	myWindow.Resize(fyne.NewSize(600, 200))
	myWindow.ShowAndRun()
}
