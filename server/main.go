package main

import (
	"go_blog/server/core"
	"go_blog/server/flag"
	"go_blog/server/global"
	"go_blog/server/initialize"
)
func main() {
	global.Config = core.InitConf()
	global.Log = core.InitLogger()
	initialize.OtherInit()
	global.DB = initialize.InitGorm()
	global.Redis = initialize.ConnectRedis()
	global.ESClient = initialize.ConnectEs()

	// 显式关闭所有连接
    defer func() {
        // 关闭 Redis 连接
        global.Redis.Close()
        
        // 关闭 MySQL 连接
        if sqlDB, err := global.DB.DB(); err == nil {
            sqlDB.Close()
        }
    }()

	
	flag.InitFlag()
    
	// 初始化定时任务
	initialize.InitCron()

	core.RunServer()
}
