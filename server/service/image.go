package service

import(
	"gorm.io/gorm"
	"mime/multipart"
	"go_blog/server/global"
	"go_blog/server/model/appTypes"
	"go_blog/server/model/database"
	"go_blog/server/model/other"
	"go_blog/server/model/request"
	"go_blog/server/utils"
	"go_blog/server/utils/upload"
)

type ImageService struct {

}

func (imageService *ImageService) UploadImage(header *multipart.FileHeader) (string, error) {
	oss := upload.NewOss()
	url, filename,err := oss.UploadImage(header)
	if err != nil {
		return "", err
	}
	return url, global.DB.Create(&database.Image{
		Name:     filename,
		URL:      url,
		Category: appTypes.Null,
		Storage:  global.Config.System.Storage(),
	}).Error
}

func (imageService *ImageService) DeleteImage(req request.DeleteImageReq) error {
	if(len(req.IDs) == 0) {
		return nil
	}
	 var Images []database.Image
	 err := global.DB.Where("id IN ?", req.IDs).Find(&Images).Error
	 if err != nil {
		return err
	 }
	 for _, image := range Images {
		if err := global.DB.Transaction(func(tx *gorm.DB) error {
			oss := upload.NewOssWithStorage(image.Storage)
			if err := global.DB.Delete(&image).Error; err != nil {
				return err
			}
			return oss.DeleteImage(image.Name)
		}); err != nil {
			return err
		}
	}
	return nil
}

func (imageService *ImageService) GetImageList(info request.ImageList) (list interface{}, total int64, err error) {
	db := global.DB

	if info.Name != nil {
		db = db.Where("name LIKE ?", "%"+*info.Name+"%")
	}

	if info.Category != nil {
		category := appTypes.ToCategory(*info.Category)
		db = db.Where("category = ?", category)
	}

	if info.Storage != nil {
		storage := appTypes.ToStorage(*info.Storage)
		db = db.Where("storage = ?", storage)
	}

	option := other.MySQLOption{
		PageInfo: info.PageInfo,
		Where:    db,
	}
	return utils.MySQLPagination(&database.Image{}, option)
}