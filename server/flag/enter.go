package flag

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"go.uber.org/zap"
	"os"
	"go_blog/server/global"
)

// 定义 CLI 标志，用于不同操作的命令行选项
var (
	sqlFlag = &cli.BoolFlag{
		Name:  "sql",
		Usage: "Initializes the srtucture of the MySQL database table.",
	}
	sqlExportFlag = &cli.BoolFlag{
		Name:  "sql-export",
		Usage: "Exports SQL data to a specified file.",
	}
	sqlImportFlag = &cli.StringFlag{
		Name:  "sql-import",
		Usage: "Imports SQL data from a specified file.",
	}
	esFlag = &cli.BoolFlag{
		Name:  "es",
		Usage: "Initializes the Elasticsearch index.",
	}
	esExportFlag = &cli.BoolFlag{
		Name:  "es-export",
		Usage: "Exports data from Elasticsearch to a specified file.",
	}
	esImportFlag = &cli.StringFlag{
		Name:  "es-import",
		Usage: "Imports data into Elasticsearch from a specified file.",
	}
	adminFlag = &cli.BoolFlag{
		Name:  "admin",
		Usage: "Creates an administrator using the name, email and address specified in the config.yaml file.",
	}
)

// Run 执行基于命令行标志的相应操作
// 该函数是 CLI 应用的核心处理函数，根据用户提供的命令行标志执行不同的管理操作
// 参数：
//   - c: CLI 上下文，包含命令行参数和标志信息
// 功能：
//   1. 检查是否设置了多个标志，确保只执行一个命令
//   2. 根据不同的标志选择执行相应的操作
//   3. 记录操作结果（成功或失败）
func Run(c *cli.Context) {
    // 检查是否设置了多个标志，确保一次只执行一个命令
    if c.NumFlags() > 1 {
        err := cli.NewExitError("Only one command can be specified", 1)
        global.Log.Error("Invalid command usage:", zap.Error(err))
        os.Exit(1)
    }

    // 根据不同的标志选择执行的操作
    switch {
    // 执行 SQL 表结构初始化
    case c.Bool(sqlFlag.Name):
        if err := SQL(); err != nil {
            global.Log.Error("Failed to create table structure:", zap.Error(err))
            return
        } else {
            global.Log.Info("Successfully created table structure")
        }
    // 执行 SQL 数据导出
    case c.Bool(sqlExportFlag.Name):
        if err := SQLExport(); err != nil {
            global.Log.Error("Failed to export SQL data:", zap.Error(err))
        } else {
            global.Log.Info("Successfully exported SQL data")
        }
    // 执行 SQL 数据导入
    case c.IsSet(sqlImportFlag.Name):
        if errs := SQLImport(c.String(sqlImportFlag.Name)); len(errs) > 0 {
            var combinedErrors string
            for _, err := range errs {
                combinedErrors += err.Error() + "\n"
            }
            err := errors.New(combinedErrors)
            global.Log.Error("Failed to import SQL data:", zap.Error(err))
        } else {
            global.Log.Info("Successfully imported SQL data")
        }
    // 执行 Elasticsearch 索引初始化
    case c.Bool(esFlag.Name):
        if err := Elasticsearch(); err != nil {
            global.Log.Error("Failed to create ES indices:", zap.Error(err))
        } else {
            global.Log.Info("Successfully created ES indices")
        }
    // 执行 Elasticsearch 数据导出
    case c.Bool(esExportFlag.Name):
        if err := ElasticsearchExport(); err != nil {
            global.Log.Error("Failed to export ES data:", zap.Error(err))
        } else {
            global.Log.Info("Successfully exported ES data")
        }
    // 执行 Elasticsearch 数据导入
    case c.IsSet(esImportFlag.Name):
        if num, err := ElasticsearchImport(c.String(esImportFlag.Name)); err != nil {
            global.Log.Error("Failed to import ES data:", zap.Error(err))
        } else {
            global.Log.Info(fmt.Sprintf("Successfully imported ES data, totaling %d records", num))
        }
    // 创建管理员用户
    case c.Bool(adminFlag.Name):
        if err := Admin(); err != nil {
            global.Log.Error("Failed to create an administrator:", zap.Error(err))
        } else {
            global.Log.Info("Successfully created an administrator")
        }
    // 处理未知命令
    default:
        err := cli.NewExitError("unknown command", 1)
        global.Log.Error(err.Error(), zap.Error(err))
    }
}

// NewApp 创建并配置一个新的 CLI 应用程序
// 该函数初始化一个 CLI 应用实例，设置应用名称、命令行标志和默认执行函数
// 返回值：
//   - *cli.App: 配置完成的 CLI 应用程序实例，可用于执行命令行操作
func NewApp() *cli.App {
    // 创建一个新的 CLI 应用实例
    app := cli.NewApp()
    
    // 设置应用名称为 "Go Blog"
    app.Name = "Go Blog"
    
    // 配置命令行标志，包括数据库操作、Elasticsearch 操作和管理员创建等功能
    app.Flags = []cli.Flag{
        sqlFlag,       // 初始化 MySQL 数据库表结构
        sqlExportFlag, // 导出 SQL 数据到文件
        sqlImportFlag, // 从文件导入 SQL 数据
        esFlag,        // 初始化 Elasticsearch 索引
        esExportFlag,  // 导出 ES 数据到文件
        esImportFlag,  // 从文件导入 ES 数据
        adminFlag,     // 创建管理员用户
    }
    
    // 设置默认执行函数为 Run，当执行命令时会调用该函数
    app.Action = Run
    
    // 返回配置完成的 CLI 应用实例
    return app
}


// InitFlag 初始化并运行 CLI 应用程序
// 该函数检查命令行参数，如果存在参数则创建并运行 CLI 应用，处理相应的管理操作
// 功能：
//   1. 检查是否存在命令行参数（os.Args > 1）
//   2. 如果存在参数，创建 CLI 应用并执行命令
//   3. 处理命令执行过程中的错误
//   4. 特殊处理帮助命令（-h 或 -help）
//   5. 执行完成后退出程序
// 注意：
//   - 此函数在检测到命令行参数时会调用 os.Exit(0) 退出程序
//   - 如果没有命令行参数，则函数正常返回，程序继续执行其他初始化操作
func InitFlag() {
    // 检查是否存在命令行参数（参数数量大于 1）
    if len(os.Args) > 1 {
        // 创建 CLI 应用实例
        app := NewApp()
        // 运行应用，处理命令行参数
        err := app.Run(os.Args)
        // 处理运行过程中的错误
        if err != nil {
            global.Log.Error("Application execution encountered an error:", zap.Error(err))
            os.Exit(1) // 出错时退出程序，返回错误码 1
        }
        // 命令执行完成后退出程序，返回成功码 0
        os.Exit(0)
    }
    // 如果没有命令行参数，函数正常返回，程序继续执行
}