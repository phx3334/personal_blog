package initialize

import(
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
	"os"
	"go_blog/server/global"
)

// ConnectES 连接到 Elasticsearch 服务器并返回 Elasticsearch 客户端实例
// 该函数负责：
// 1. 从全局配置中获取 Elasticsearch 配置信息
// 2. 创建 Elasticsearch 客户端配置
// 3. 根据配置决定是否启用控制台日志输出
// 4. 创建 Elasticsearch 客户端实例
// 5. 处理客户端创建失败的情况
func ConnectEs() *elasticsearch.TypedClient {
    // 从全局配置中获取 Elasticsearch 配置信息
    esCfg := global.Config.ES
    
    // 创建 Elasticsearch 客户端配置
    cfg := elasticsearch.Config{
        Addresses: []string{esCfg.URL}, // Elasticsearch 服务器地址
        Username:  esCfg.Username,      // Elasticsearch 用户名（如果需要认证）
        Password:  esCfg.Password,      // Elasticsearch 密码（如果需要认证）
    }

    // 如果配置了控制台输出，则启用彩色日志记录
    if esCfg.IsConsolePrint {
        cfg.Logger = &elastictransport.ColorLogger{
            Output:            os.Stdout,          // 输出到标准输出
            EnableRequestBody: true,               // 启用请求体日志
            EnableResponseBody: true,              // 启用响应体日志
        }
    }

    // 创建 Elasticsearch 客户端实例
    client, err := elasticsearch.NewTypedClient(cfg)
    
    // 处理客户端创建失败的情况
    if err != nil {
        global.Log.Error("Failed to create elasticsearch client", zap.Error(err))
        os.Exit(1) // 创建客户端失败时直接退出程序
    }
    
    // 创建成功，返回 Elasticsearch 客户端实例
    return client
}

