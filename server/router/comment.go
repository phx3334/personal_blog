package router

import (
	"github.com/gin-gonic/gin"
	"go_blog/server/api"
)

type CommentRouter struct {
}

// InitCommentRouter 初始化评论相关路由
// 功能：为评论模块注册不同权限级别的路由
// 参数：
//   PrivateRouter: 私有路由组（需要登录）
//   PublicRouter: 公开路由组（无需登录）
//   AdminRouter: 管理员路由组（需要管理员权限）
func (c *CommentRouter) InitCommentRouter(PrivateRouter *gin.RouterGroup, PublicRouter *gin.RouterGroup, AdminRouter *gin.RouterGroup) {

	// 创建评论相关的路由子组
	commentPrivateRouter := PrivateRouter.Group("comment") // 私有评论路由组
	commentPublicRouter := PublicRouter.Group("comment")   // 公开评论路由组
	commentAdminRouter := AdminRouter.Group("comment")     // 管理员评论路由组

	// 获取评论 API 实例
	commentApi := api.ApiGroupApp.CommentApi

	// 私有路由组（需要登录）
	{
		commentPrivateRouter.POST("create", commentApi.CommentCreate) // 创建评论
		commentPrivateRouter.DELETE("delete", commentApi.CommentDelete) // 删除评论
		commentPrivateRouter.GET("info", commentApi.CommentInfo) // 获取评论详情
	}

	// 公开路由组（无需登录）
	{
		commentPublicRouter.GET(":article_id", commentApi.CommentInfoByArticleID) // 根据文章 ID 获取评论
		commentPublicRouter.GET("new", commentApi.CommentNew) // 获取最新评论
	}

	// 管理员路由组（需要管理员权限）
	{
		commentAdminRouter.GET("list", commentApi.CommentList) // 获取评论列表（管理员）
	}
}