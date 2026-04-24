package router

import(
	"github.com/gin-gonic/gin"
	"go_blog/server/api"
)

type FeedbackRouter struct {
}

func (f *FeedbackRouter) InitFeedbackRouter(privateGroup, publicGroup, adminGroup *gin.RouterGroup) {
	privateFeedbackGroup := privateGroup.Group("/feedback")
	publicFeedbackGroup := publicGroup.Group("/feedback")
	adminFeedbackGroup := adminGroup.Group("/feedback")

	feedbackApi := api.ApiGroupApp.FeedbackApi

	{
		privateFeedbackGroup.POST("create", feedbackApi.CreateFeedback)
		privateFeedbackGroup.POST("info", feedbackApi.FeedbackInfo)
	}

	{
		publicFeedbackGroup.GET("new", feedbackApi.FeedbackNew)
	}
	{
		adminFeedbackGroup.DELETE("delete", feedbackApi.DeleteFeedback)
		adminFeedbackGroup.PUT("reply", feedbackApi.ReplyFeedback)
		adminFeedbackGroup.GET("list", feedbackApi.FeedbackList)
	}
}
