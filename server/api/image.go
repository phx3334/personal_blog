package api

import (
	"go_blog/server/global"
	"go_blog/server/model/request"
	"go_blog/server/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ImageApi struct {

}

func (i *ImageApi) UploadImage(c *gin.Context) {
	_,header,err := c.Request.FormFile("image")
	if err!= nil {
		global.Log.Error("err.Error()", zap.Any("err", err))
		response.FailWithMessage("Fail to upload image", c)
		return
	}

	url,err := imageService.UploadImage(header)
	if err != nil {
		global.Log.Error("Fail to upload image", zap.Any("err", err))
		response.FailWithMessage("Fail to upload image", c)
		return
	}
	response.OkWithData(response.ImageUpload{
		Url: url,
		OssType: global.Config.System.OssType,
	},c)
}

func (i *ImageApi) DeleteImage(c *gin.Context){
	var req request.DeleteImageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("Invalid request data", c)
		return
	}
	err := imageService.DeleteImage(req)
	if(err != nil) {
		global.Log.Error("Fail to delete image", zap.Any("err", err))
		response.FailWithMessage("Fail to delete image", c)
		return
	}
	response.OkWithMessage("Success to delete image", c)
}

func (i *ImageApi) GetImageList(c *gin.Context) {
	var req request.ImageList
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage("Invalid request data", c)
		return
	}
	imageList, total, err := imageService.GetImageList(req)
	if err != nil {
		global.Log.Error("Fail to get image list", zap.Any("err", err))
		response.FailWithMessage("Fail to get image list", c)
		return
	}
	response.OkWithData(response.PageResult{
        List: imageList,
		Total: total,
	},c)	
}