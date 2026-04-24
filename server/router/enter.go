package router

type RouterGroup struct {
	BaseRouter
	UserRouter
	ImageRouter
	ArticleRouter
	ConfigRouter
	FeedbackRouter
	CommentRouter
	WebsiteRouter
}

var RouterGroupApp = new(RouterGroup)
