package api

import "go_blog/server/service"

type ApiGroup struct {
	BaseApi
	UserApi
	ImageApi
	ArticleApi
	CommentApi
	FeedbackApi
	ConfigApi
	WebsiteApi
}

var ApiGroupApp = new(ApiGroup)

var baseService = service.ServiceGroupApp.BaseService
var userService = service.ServiceGroupApp.UserService
var jwtService = service.ServiceGroupApp.JwtService
var qqService = service.ServiceGroupApp.QQService
var imageService = service.ServiceGroupApp.ImageService
var articleService = service.ServiceGroupApp.ArticleService
var commentService = service.ServiceGroupApp.CommentService
var feedbackService = service.ServiceGroupApp.FeedbackService
var configService = service.ServiceGroupApp.ConfigService
var websiteService = service.ServiceGroupApp.WebsiteService
