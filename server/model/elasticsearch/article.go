package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Article 文章表
type Article struct {
	CreatedAt string `json:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at"` // 更新时间

	Cover    string   `json:"cover"`    // 文章封面
	Title    string   `json:"title"`    // 文章标题
	Keyword  string   `json:"keyword"`  // 文章标题-关键字
	Category string   `json:"category"` // 文章类别
	Tags     []string `json:"tags"`     // 文章标签
	Abstract string   `json:"abstract"` // 文章简介
	Content  string   `json:"content"`  // 文章内容

	Views    int `json:"views"`    // 浏览量
	Comments int `json:"comments"` // 评论量
	Likes    int `json:"likes"`    // 收藏量

	ID string `json:"id"` // 文章ID
}

// ArticleIndex 文章 ES 索引
func ArticleIndex() string {
	return "article_index"
}

// ArticleMapping 定义 Elasticsearch 中文章索引的映射结构
// 作用：告诉 Elasticsearch 如何索引和存储文章的各个字段，优化搜索和排序性能
// 返回值：
//   *types.TypeMapping: 文章索引的映射定义
func ArticleMapping() *types.TypeMapping {
	return &types.TypeMapping{
		Properties: map[string]types.Property{
			// 文章创建时间，日期类型，格式为 "年-月-日 时:分:秒"
			"created_at": types.DateProperty{NullValue: nil, Format: func(s string) *string { return &s }("yyyy-MM-dd HH:mm:ss")},
			// 文章更新时间，日期类型，格式为 "年-月-日 时:分:秒"
			"updated_at": types.DateProperty{NullValue: nil, Format: func(s string) *string { return &s }("yyyy-MM-dd HH:mm:ss")},
			// 封面图片路径，文本类型，支持全文搜索
			"cover":      types.TextProperty{},
			// 文章标题，文本类型，支持全文搜索
			"title":      types.TextProperty{},
			// 文章关键词，关键字类型，用于精确匹配
			"keyword":    types.KeywordProperty{},
			// 文章分类，关键字类型，用于精确匹配
			"category":   types.KeywordProperty{},
			// 文章标签，关键字类型数组，用于精确匹配多个标签
			"tags":       []types.KeywordProperty{},
			// 文章摘要，文本类型，支持全文搜索
			"abstract":   types.TextProperty{},
			// 文章内容，文本类型，支持全文搜索
			"content":    types.TextProperty{},
			// 文章浏览量，整数类型，用于排序和统计
			"views":      types.IntegerNumberProperty{},
			// 文章评论数，整数类型，用于排序和统计
			"comments":   types.IntegerNumberProperty{},
			// 文章点赞数，整数类型，用于排序和统计
			"likes":      types.IntegerNumberProperty{},
		},
	}
}