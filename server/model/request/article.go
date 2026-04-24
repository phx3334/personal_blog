package request

type CreateArticle struct {
	Cover string `json:"cover" binding:"required"` // 文章封面图片URL
	Title string `json:"title" binding:"required"` // 文章标题
	Category string `json:"category" binding:"required"` // 文章分类
	Tags []string `json:"tags" binding:"required"` // 文章标签列表
	Content string `json:"content" binding:"required"` // 文章内容
	Abstract string `json:"abstract" binding:"max=200"` // 文章摘要
}

type UpdateArticle struct {
	ID string `json:"id" binding:"required"` // 文章ID
	Cover string `json:"cover" binding:"required"` // 文章封面图片URL
	Title string `json:"title" binding:"required"` // 文章标题
	Category string `json:"category" binding:"required"` // 文章分类
	Tags []string `json:"tags" binding:"required"` // 文章标签列表
	Content string `json:"content" binding:"required"` // 文章内容
	Abstract string `json:"abstract" binding:"max=200"` // 文章摘要
}

type ArticleInfoByID struct {
	ID string `uri:"id" binding:"required"` // 文章ID
}


type ArticleSearch struct {
	Query    string `json:"query" form:"query" uri:"query" binding:"max=50"`
	Category string `json:"category" form:"category" uri:"category"`
	Tag      string `json:"tag" form:"tag" uri:"tag"`
	Sort     string `json:"sort" form:"sort" uri:"sort"`
	Order    string `json:"order" form:"order" uri:"order" binding:"required"`
	PageInfo
}

type ArticleDelete struct{
	IDs []string `json:"ids" binding:"required"` // 文章ID列表
}

type ArticleList struct {
	Title *string `json:"title" form:"title" ` // 文章标题
	Category *string `json:"category" form:"category" ` // 文章分类
	Abstract *string `json:"abstract" form:"abstract" ` // 文章摘要
	PageInfo
}

type ArticleLike struct{
	UserID    uint   `json:"-"`
	ArticleID string `json:"id" binding:"required"` // 文章ID
}

type ArticleLikesList struct {
	PageInfo
	UserID uint `json:"-"`
	
}