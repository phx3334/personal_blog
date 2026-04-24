package router

import (
	"github.com/gin-gonic/gin"
	"go_blog/server/api"
)

type ImageRouter struct {
}

func (i *ImageRouter) InitImageRouter(AdminRouter *gin.RouterGroup) {
	ImageRouter := AdminRouter.Group("image")
	imageApi := api.ApiGroupApp.ImageApi
	{
		ImageRouter.POST("upload", imageApi.UploadImage) // 上传图片
		ImageRouter.DELETE("delete", imageApi.DeleteImage)
		ImageRouter.GET("list", imageApi.GetImageList)
	}
}