package router

import(
	"github.com/gin-gonic/gin"
	"go_blog/server/api"
)

type ArticleRouter struct {

}

func (a *ArticleRouter) InitArticleRouter(privateGroup *gin.RouterGroup, publicGroup *gin.RouterGroup, adminGroup *gin.RouterGroup) {
	 privateArticleRouter := privateGroup.Group("article")
	 publicArticleRouter := publicGroup.Group("article")
	 adminArticleRouter := adminGroup.Group("article")

	 articleApi := api.ApiGroupApp.ArticleApi
	{
		privateArticleRouter.POST("like", articleApi.ArticleLike) // 点赞文章
		privateArticleRouter.GET("islike", articleApi.ArticleIsLike) // 是否点赞
		privateArticleRouter.GET("likelist", articleApi.ArticleLikesList) // 获取点赞列表
	}

	{
		publicArticleRouter.GET(":id", articleApi.ArticleInfoByID)
		publicArticleRouter.GET("search", articleApi.ArticleSearch)
		publicArticleRouter.GET("category", articleApi.ArticleCategory)
		publicArticleRouter.GET("tags", articleApi.ArticleTags)
	}

	{
		adminArticleRouter.POST("create", articleApi.ArticleCreate)
		adminArticleRouter.DELETE("delete", articleApi.ArticleDelete)
		adminArticleRouter.PUT("update", articleApi.ArticleUpdate)
		adminArticleRouter.GET("list", articleApi.ArticleList)
	}
}
