package router

import (
	"go_blog/server/api"

	"github.com/gin-gonic/gin"
)

type BaseRouter struct {
}

// InitBaseRouter 初始化基础路由
// 该方法用于注册系统的基础 API 路由，包括验证码、邮箱验证和第三方登录等功能
// 参数：
//   - Router: 父级路由组，基础路由会注册到该路由组下
// 功能：
//   1. 创建名为 "base" 的子路由组
//   2. 获取 BaseApi 实例，用于处理基础 API 请求
//   3. 注册以下路由：
//      - POST /base/captcha: 生成验证码
//      - POST /base/sendEmailVerificationCode: 发送邮箱验证码
//      - GET /base/qqLoginURL: 获取 QQ 登录 URL
func (b *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	// 创建基础路由组，所有基础 API 都在 /base 路径下
	baseRouter := Router.Group("base")
	// 获取 BaseApi 实例，用于处理基础 API 请求
	baseApi := api.ApiGroupApp.BaseApi
	{
		// 注册验证码生成路由
		baseRouter.POST("captcha", baseApi.Captcha)
		// 注册邮箱验证码发送路由
		baseRouter.POST("sendEmailVerificationCode", baseApi.SendEmailVerificationCode)
		// 注册 QQ 登录 URL 获取路由
		baseRouter.GET("qqLoginURL", baseApi.QQLoginURL)
	}
}
