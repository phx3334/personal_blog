package initialize

import (
	"go_blog/server/global"
	"go_blog/server/router"	
	"net/http"
	"go_blog/server/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化 Gin 路由引擎，配置中间件和路由组
// 该函数负责：
// 1. 设置 Gin 运行模式
// 2. 配置全局中间件（日志、恢复、会话）
// 3. 注册静态文件服务
// 4. 创建并配置路由组（公共、私有、管理员）
// 5. 初始化各模块路由
// 返回值：
//   - *gin.Engine: 配置完成的 Gin 引擎实例，可用于启动 HTTP 服务器
func InitRouter() *gin.Engine {
	// 设置 Gin 运行模式（如 debug、release 等），从全局配置中读取
	// 运行模式会影响 Gin 的日志输出和错误处理行为
	gin.SetMode(global.Config.System.Env)
	
	// 创建默认的 Gin 引擎，包含内置的 Logger 和 Recovery 中间件
	// Logger 中间件：记录请求日志
	// Recovery 中间件：捕获 panic 并恢复，返回 500 错误
	Router := gin.New()

	// 应用自定义的日志中间件和恢复中间件
	// GinLogger: 自定义请求日志记录，支持结构化日志
	// GinRecovery(true): 自定义 panic 恢复，true 表示打印堆栈信息
	Router.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	

	// 创建基于 cookie 的会话存储，使用配置中的密钥
	// global.Config.System.SessionsSecret: 用于加密会话数据的密钥
	var store = cookie.NewStore([]byte(global.Config.System.SessionsSecret))
	
	// 注册会话中间件，会话名称为 "sessionid"
	// 后续可通过 sessions.Default(c) 获取会话实例，用于存储临时数据（如验证码）
	Router.Use(sessions.Sessions("sessionid", store))
	
	// 注册静态文件服务，使用配置中指定的路径
	// 第一个参数: URL 路径前缀（如 /upload）
	// 第二个参数: 本地文件系统目录，指定静态文件的存储位置
	// 用于提供图片、CSS、JavaScript 等静态文件的访问
	Router.StaticFS(global.Config.Upload.Path, http.Dir(global.Config.Upload.Path))

	// 初始化路由组实例，用于管理所有路由
	// router.RouterGroupApp: 全局路由组实例，包含所有路由初始化方法
	routerGroup := router.RouterGroupApp
	
	// 创建三个路由组，均使用配置的路由前缀
	// global.Config.System.RouterPrefix: 路由前缀，如 /api
	
	// 1. 公共路由组：无需认证即可访问（如登录、注册、验证码等）
    publicGroup := Router.Group(global.Config.System.RouterPrefix)
	
	// 2. 私有路由组：需要 JWT 认证（如用户个人信息、发布文章等）
    privateGroup := Router.Group(global.Config.System.RouterPrefix)
	privateGroup.Use(middleware.JWTAuth()) // 对私有路由组应用 JWT 认证中间件
	
	// 3. 管理员路由组：需要 JWT 认证和管理员权限（如用户管理、系统设置等）
	adminGroup := Router.Group(global.Config.System.RouterPrefix)
	adminGroup.Use(middleware.JWTAuth()).Use(middleware.AdminAuth()) // 应用 JWT 认证和管理员权限中间件
	
	// 初始化基础路由（公共路由）
	// 包含登录、注册、验证码、邮箱验证等无需认证的路由
	{
		routerGroup.InitBaseRouter(publicGroup)
	} 
	
	// 初始化用户相关路由
	{
		routerGroup.InitUserRouter(privateGroup, publicGroup, adminGroup)
		routerGroup.InitArticleRouter(privateGroup, publicGroup, adminGroup)
		routerGroup.InitFeedbackRouter(privateGroup, publicGroup, adminGroup)
		routerGroup.InitCommentRouter(privateGroup, publicGroup, adminGroup)
	}
   
	{
      routerGroup.InitImageRouter(adminGroup)
	  routerGroup.InitConfigRouter(adminGroup)
	  routerGroup.InitWebsiteRouter(adminGroup,publicGroup)
	}
	// 返回配置完成的 Gin 引擎，用于启动 HTTP 服务器
	return Router
}