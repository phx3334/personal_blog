package router

import (
	"github.com/gin-gonic/gin"
	"go_blog/server/api"
	"go_blog/server/middleware"
)

type UserRouter struct {
}

// InitUserRouter 初始化用户相关路由，分别配置私有、公共、登录和管理员路由
// 参数：
//   - PrivateRouter: 私有路由组，需要 JWT 认证
//   - PublicRouter: 公共路由组，无需认证
//   - AdminRouter: 管理员路由组，需要 JWT 认证和管理员权限
func (u *UserRouter) InitUserRouter(PrivateRouter *gin.RouterGroup, PublicRouter *gin.RouterGroup, AdminRouter *gin.RouterGroup) {
	// 创建用户私有路由组，路径前缀为 /user
	// 需要 JWT 认证，用于处理用户个人相关操作
	userPrivateRouter := PrivateRouter.Group("user")
	
	// 创建用户公共路由组，路径前缀为 /user
	// 无需认证，用于处理公开的用户相关操作
	userPublicRouter := PublicRouter.Group("user")
	
	// 创建用户登录路由组，路径前缀为 /user
	// 无需认证，但应用登录记录中间件，用于处理登录相关操作
	userLoginRouter := PublicRouter.Group("user").Use(middleware.LoginRecord())
	
	// 创建用户管理员路由组，路径前缀为 /user
	// 需要 JWT 认证和管理员权限，用于处理用户管理操作
	userAdminRouter := AdminRouter.Group("user")
	
	// 获取用户 API 实例，用于处理具体的业务逻辑
	userApi := api.ApiGroupApp.UserApi
	
	// 私有路由：需要用户登录后才能访问
	{
		userPrivateRouter.POST("logout", userApi.Logout)           // 用户登出
		userPrivateRouter.PUT("resetPassword", userApi.UserResetPassword) // 重置密码
		userPrivateRouter.GET("info", userApi.UserInfo)           // 获取用户信息
		userPrivateRouter.PUT("changeInfo", userApi.UserChangeInfo) // 修改用户信息
		userPrivateRouter.GET("weather", userApi.UserWeather)       // 获取天气信息
		userPrivateRouter.GET("chart", userApi.UserChart)           // 获取用户数据图表
	}
	
	// 公共路由：无需登录即可访问
	{
		userPublicRouter.POST("forgotPassword", userApi.ForgotPassword) // 忘记密码
		userPublicRouter.GET("card", userApi.UserCard)               // 获取用户名片信息
	}
	
	// 登录相关路由：无需登录，但会记录登录行为
	{
		userLoginRouter.POST("register", userApi.Register) // 用户注册
		userLoginRouter.POST("login", userApi.Login)       // 用户登录
	}
	
	// 管理员路由：需要管理员权限才能访问
	{
		userAdminRouter.GET("list", userApi.UserList)         // 获取用户列表
		userAdminRouter.PUT("freeze", userApi.UserFreeze)     // 冻结用户
		userAdminRouter.PUT("unfreeze", userApi.UserUnfreeze) // 解冻用户
		userAdminRouter.GET("loginList", userApi.UserLoginList) // 获取用户登录记录
	}
}