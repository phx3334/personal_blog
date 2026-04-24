package initialize

import (
	"go_blog/server/global"
	"os"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitGorm 初始化 GORM 数据库连接并返回数据库实例
// 该函数负责：
// 1. 从全局配置中获取 MySQL 配置
// 2. 建立数据库连接
// 3. 配置连接池参数
// 4. 处理连接过程中的错误
func InitGorm() *gorm.DB {
	// 从全局配置中获取 MySQL 配置信息
	mysqlCfg := global.Config.Mysql
	
	// 使用配置的 DSN (Data Source Name) 打开数据库连接
	// 同时设置 GORM 日志级别
	db, err := gorm.Open(mysql.Open(mysqlCfg.Dsn()), &gorm.Config{
		Logger: logger.Default.LogMode(mysqlCfg.LogLevel()),
	})
    // 处理数据库连接错误
	if err != nil {
		global.Log.Error("Failed to connect to database", zap.Error(err))
		os.Exit(1) // 连接失败时直接退出程序
	}
	

	// 获取底层的 *sql.DB 实例，用于配置连接池参数
	sqlDB, err := db.DB()
	if err != nil {
		global.Log.Error("Failed to get database connection", zap.Error(err))
		os.Exit(1) // 获取连接失败时直接退出程序
	}

	// 设置连接池的最大空闲连接数
	sqlDB.SetMaxIdleConns(mysqlCfg.MaxIdleConns)
	// 设置连接池的最大打开连接数
	sqlDB.SetMaxOpenConns(mysqlCfg.MaxOpenConns)

	// 返回初始化好的 GORM 数据库实例
	return db
}
