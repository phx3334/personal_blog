package router

import (
	"github.com/gin-gonic/gin"
	"go_blog/server/api"
)

type ImageRouter struct {
}

func (i *ImageRouter) InitImageRouter(AdminRouter *gin.RouterGroup, PrivateRouter *gin.RouterGroup) {
	imageApi := api.ApiGroupApp.ImageApi

	// 私有路由：普通用户上传图片（用于修改头像）
	privateImageRouter := PrivateRouter.Group("image")
	{
		privateImageRouter.POST("upload", imageApi.UploadImage) // 上传图片
	}

	// 管理员路由：图片管理功能
	adminImageRouter := AdminRouter.Group("image")
	{
		adminImageRouter.DELETE("delete", imageApi.DeleteImage)
		adminImageRouter.GET("list", imageApi.GetImageList)
	}
}
