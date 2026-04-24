package initialize

import(
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"os"
	"go_blog/server/global"
)

// ConnectRedis 连接到 Redis 服务器并返回 Redis 客户端实例
// 该函数负责：
// 1. 从全局配置中获取 Redis 配置信息
// 2. 创建 Redis 客户端并设置连接参数
// 3. 测试与 Redis 服务器的连接
// 4. 处理连接失败的情况
func ConnectRedis() *redis.Client {
    // 从全局配置中获取 Redis 配置信息
    redisCfg := global.Config.Redis
    
    // 创建 Redis 客户端实例，设置连接参数
    client := redis.NewClient(&redis.Options{
        Addr:     redisCfg.Address, // Redis 服务器地址和端口
        Password: redisCfg.Password, // Redis 服务器密码（如果有）
        DB:       redisCfg.DB,       // 使用的 Redis 数据库编号
    })
    
    // 测试与 Redis 服务器的连接，发送 Ping 命令
    _, err := client.Ping().Result()
    
    // 处理连接失败的情况
    if err != nil {
        global.Log.Error("Failed to connect to redis", zap.Error(err))
        os.Exit(1) // 连接失败时直接退出程序
    }
    
    // 连接成功，返回 Redis 客户端实例
    return client
}