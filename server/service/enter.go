package service

type ServiceGroup struct {
	BaseService
	EsService
	JwtService
	GaodeService
	UserService
	QQService
	ImageService
	ArticleService
	CommentService
	FeedbackService
	ConfigService
	WebsiteService
	CalendarService
}

var ServiceGroupApp = new(ServiceGroup)
