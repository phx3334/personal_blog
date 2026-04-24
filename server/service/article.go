package service

import (
	"context"
	"errors"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
    "github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptlanguage"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"gorm.io/gorm"
	"go_blog/server/global"
	"go_blog/server/model/appTypes"
	"go_blog/server/model/database"
	"go_blog/server/model/elasticsearch"
	"go_blog/server/model/other"
	"go_blog/server/model/request"
	"go_blog/server/utils"
	"time"
	"strconv"
)


type ArticleService struct {
}

// ArticleInfoByID 根据文章 ID 获取文章信息
// 参数:
//   id: 文章的唯一标识
// 返回值:
//   elasticsearch.Article: 文章信息
//   error: 错误信息
func (articleService *ArticleService) ArticleInfoByID(id string) (elasticsearch.Article, error) {
    // 异步更新浏览量，使用 goroutine 避免阻塞主流程
    go func() {
        // 创建文章浏览量服务实例
        articleView := articleService.NewArticleView()
        // 更新指定文章的浏览量，忽略错误（不影响主流程）
        _ = articleView.Set(id)
    }()
    // 调用 Get 方法获取文章详细信息并返回
    return articleService.Get(id)
}

// ArticleSearch 文章搜索方法
// 参数:
//   info: 搜索条件，包含查询关键词、标签、类别、排序方式等
// 返回值:
//   interface{}: 搜索结果列表
//   int64: 符合条件的总记录数
//   error: 错误信息
func (articleService *ArticleService) ArticleSearch(info request.ArticleSearch) (interface{}, int64, error) {
    // 创建 Elasticsearch 搜索请求
    req := &search.Request{
        Query: &types.Query{},
    }

    // 构建布尔查询
    boolQuery := &types.BoolQuery{}

    // 根据查询字段查询（标题、关键词、摘要、内容）
    if info.Query != "" {
        boolQuery.Should = []types.Query{
            {Match: map[string]types.MatchQuery{"title": {Query: info.Query}}},
            {Match: map[string]types.MatchQuery{"keyword": {Query: info.Query}}},
            {Match: map[string]types.MatchQuery{"abstract": {Query: info.Query}}},
            {Match: map[string]types.MatchQuery{"content": {Query: info.Query}}},
        }
    }

    // 根据标签筛选
    if info.Tag != "" {
        boolQuery.Must = []types.Query{
            {Match: map[string]types.MatchQuery{"tags": {Query: info.Tag}}},
        }
    }

    // 根据类别筛选
    if info.Category != "" {
        boolQuery.Filter = []types.Query{
            {Term: map[string]types.TermQuery{"category": {Value: info.Category}}},
        }
    }

    // 如果有查询条件，则使用 Bool 查询，否则使用 MatchAll 查询
    if boolQuery.Should != nil || boolQuery.Must != nil || boolQuery.Filter != nil {
        req.Query.Bool = boolQuery
    } else {
        req.Query.MatchAll = &types.MatchAllQuery{}
    }

    // 设置排序字段
    if info.Sort != "" {
        var sortField string
        switch info.Sort {
        case "time":
            sortField = "created_at" // 按创建时间排序
        case "view":
            sortField = "views"      // 按浏览量排序
        case "comment":
            sortField = "comments"   // 按评论数排序
        case "like":
            sortField = "likes"      // 按点赞数排序
        default:
            sortField = "created_at" // 默认按创建时间排序
        }

        // 设置排序顺序
        var order sortorder.SortOrder
        if info.Order != "asc" {
            order = sortorder.Desc // 降序
        } else {
            order = sortorder.Asc  // 升序
        }

        // 应用排序设置
        req.Sort = []types.SortCombinations{
            types.SortOptions{
                SortOptions: map[string]types.FieldSort{
                    sortField: {Order: &order},
                },
            },
        }
    }

    // 构建 Elasticsearch 分页选项
    option := other.EsOption{
        PageInfo:       info.PageInfo,       // 分页信息
        Index:          elasticsearch.ArticleIndex(), // 文章索引名
        Request:        req,                 // 搜索请求
        SourceIncludes: []string{"created_at", "cover", "title", "abstract", "category", "tags", "views", "comments", "likes"}, // 只返回指定字段
    }
    // 执行 Elasticsearch 分页搜索并返回结果
    return utils.EsPagination(context.TODO(), option)
}

// ArticleCategory 获取所有文章分类
// 返回值:
//   []database.ArticleCategory: 文章分类列表
//   error: 错误信息
func (articleService *ArticleService) ArticleCategory() ([]database.ArticleCategory, error) {
    // 定义分类列表变量
    var category []database.ArticleCategory
    // 从数据库中查询所有分类记录
    if err := global.DB.Find(&category).Error; err != nil {
        // 查询失败时返回错误
        return nil, err
    }
    // 查询成功时返回分类列表
    return category, nil
}
// ArticleTags 获取所有文章标签
// 返回值:
//   []database.ArticleTag: 文章标签列表
//   error: 错误信息
func (articleService *ArticleService) ArticleTags() ([]database.ArticleTag, error) {
    // 定义标签列表变量
    var tags []database.ArticleTag
    // 从数据库中查询所有标签记录
    if err := global.DB.Find(&tags).Error; err != nil {
        // 查询失败时返回错误
        return nil, err
    }
    // 查询成功时返回标签列表
    return tags, nil
}

// ArticleLike 文章点赞/取消点赞方法
// 功能：实现文章的点赞和取消点赞切换功能，同时更新数据库和 Elasticsearch 中的点赞数
// 参数:
//   req: 点赞请求信息，包含用户 ID 和文章 ID
// 返回值:
//   error: 操作过程中的错误信息
func (articleService *ArticleService) ArticleLike(req request.ArticleLike) error {
	// 开启数据库事务，确保操作的原子性
	return global.DB.Transaction(func(tx *gorm.DB) error {
		var al database.ArticleLike // 用于存储点赞记录
		var num int                 // 用于标记点赞(1)或取消点赞(-1)

		// 检查用户是否已点赞该文章
		if errors.Is(tx.Where("user_id = ? AND article_id = ?", req.UserID, req.ArticleID).First(&al).Error, gorm.ErrRecordNotFound) {
			// 用户未点赞，创建点赞记录
			if err := tx.Create(&database.ArticleLike{UserID: req.UserID, ArticleID: req.ArticleID}).Error; err != nil {
				return err // 创建失败，返回错误
			}
			num = 1 // 标记为点赞操作
		} else { // 用户已点赞，取消点赞
			// 删除已有的点赞记录
			if err := tx.Delete(&al).Error; err != nil {
				return err // 删除失败，返回错误
			}
			num = -1 // 标记为取消点赞操作
		}

		// 更新 Elasticsearch 中的文章点赞数
		source := "ctx._source.likes += " + strconv.Itoa(num) // 构建更新脚本
		script := types.Script{
			Source: &source,                        // 脚本内容
			Lang:   &scriptlanguage.Painless,       // 脚本语言：Painless
		}
		// 执行 Elasticsearch 更新操作，更新指定文章的点赞数
		_, err := global.ESClient.Update(elasticsearch.ArticleIndex(), req.ArticleID).Script(&script).Do(context.TODO())
		return err // 返回更新结果的错误信息
	})
}
// ArticleIsLike 检查用户是否点赞了该文章
// 参数:
//   req: 点赞请求信息，包含用户 ID 和文章 ID
// 返回值:
//   bool: 是否点赞
//   error: 错误信息
func (articleService *ArticleService) ArticleIsLike(req request.ArticleLike) (bool, error) {
	return !errors.Is(global.DB.Where("user_id = ? AND article_id = ?", req.UserID, req.ArticleID).First(&database.ArticleLike{}).Error, gorm.ErrRecordNotFound), nil
}
// ArticleLikesList 获取用户点赞的文章列表
// 参数:
//   info: 点赞列表请求信息，包含用户 ID 和分页信息
// 返回值:
//   interface{}: 点赞文章列表
//   int64: 点赞文章总数
//   error: 错误信息
func (articleService *ArticleService) ArticleLikesList(info request.ArticleLikesList) (interface{}, int64, error) {
	db := global.DB.Where("user_id = ?", info.UserID)
	option := other.MySQLOption{
		PageInfo: info.PageInfo,
		Where:    db,
	}

	l, total, err := utils.MySQLPagination(&database.ArticleLike{}, option)
	if err != nil {
		return nil, 0, err
	}
	var list []struct {
		Id_     string                `json:"_id"`
		Source_ elasticsearch.Article `json:"_source"`
	}

	for _, articleLike := range l {
		article, err := articleService.Get(articleLike.ArticleID)
		if err != nil {
			return nil, 0, err
		}
		article.UpdatedAt = ""
		article.Keyword = ""
		article.Content = ""
		list = append(list, struct {
			Id_     string                `json:"_id"`
			Source_ elasticsearch.Article `json:"_source"`
		}{
			Id_:     articleLike.ArticleID,
			Source_: article,
		})
	}
	return list, total, nil
}

// ArticleCreate 创建文章方法
// 功能：创建新文章，并更新相关的分类、标签和图片数据
// 参数:
//   req: 创建文章的请求信息，包含标题、封面、分类、标签、摘要、内容等
// 返回值:
//   error: 操作过程中的错误信息
func (articleService *ArticleService) ArticleCreate(req request.CreateArticle) error {
	// 检查文章是否已存在（通过标题判断）
	b, err := articleService.Exits(req.Title)
	if err != nil {
		return err // 检查过程中出错，返回错误
	}
	if b {
		return errors.New("the article already exists") // 文章已存在，返回错误
	}

	// 获取当前时间，用于设置文章的创建和更新时间
	now := time.Now().Format("2006-01-02 15:04:05")
	
	// 构建文章对象
	articleToCreate := elasticsearch.Article{
		CreatedAt: now,      // 创建时间
		UpdatedAt: now,      // 更新时间
		Cover:     req.Cover, // 文章封面
		Title:     req.Title, // 文章标题
		Keyword:   req.Title, // 关键词（使用标题作为关键词）
		Category:  req.Category, // 文章分类
		Tags:      req.Tags, // 文章标签
		Abstract:  req.Abstract, // 文章摘要
		Content:   req.Content, // 文章内容
	}

	// 开启数据库事务，确保操作的原子性
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 同时更新文章类别表中的数据（增加计数）
		if err := articleService.UpdateCategoryCount(tx, "", articleToCreate.Category); err != nil {
			return err // 更新类别计数失败，返回错误
		}

		// 同时更新文章标签表中的数据（增加计数）
		if err := articleService.UpdateTagsCount(tx, []string{}, articleToCreate.Tags); err != nil {
			return err // 更新标签计数失败，返回错误
		}

		// 同时更新图片表中的图片类别（将封面图片标记为封面类型）
		if err := utils.ChangeImagesCategory(tx, []string{articleToCreate.Cover}, appTypes.Cover); err != nil {
			return err // 更新封面图片类别失败，返回错误
		}

		// 查找文章内容中的插图
		illustrations, err := utils.FindIllustrations(articleToCreate.Content)
		if err != nil {
			return err // 查找插图失败，返回错误
		}

		// 更新插图的类别（标记为插图类型）
		if err := utils.ChangeImagesCategory(tx, illustrations, appTypes.Illustration); err != nil {
			return err // 更新插图类别失败，返回错误
		}

		// 执行实际的文章创建操作
		return articleService.Create(&articleToCreate)
	})
}


// ArticleDelete 删除文章方法
// 功能：删除指定的文章，并更新相关的分类、标签和图片数据
// 参数:
//   req: 删除请求信息，包含要删除的文章 ID 列表
// 返回值:
//   error: 操作过程中的错误信息
func (articleService *ArticleService) ArticleDelete(req request.ArticleDelete) error {
	// 检查是否有要删除的文章 ID，如果没有则直接返回
	if len(req.IDs) == 0 {
		return nil
	}

	// 开启数据库事务，确保操作的原子性
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 遍历要删除的文章 ID 列表
		for _, id := range req.IDs {
			// 获取要删除的文章信息
			articleToDelete, err := articleService.Get(id)
			if err != nil {
				return err // 获取文章失败，返回错误
			}

			// 同时更新文章类别表中的数据（减少计数）
			if err := articleService.UpdateCategoryCount(tx, articleToDelete.Category, ""); err != nil {
				return err // 更新类别计数失败，返回错误
			}

			// 同时更新文章标签表中的数据（减少计数）
			if err := articleService.UpdateTagsCount(tx, articleToDelete.Tags, []string{}); err != nil {
				return err // 更新标签计数失败，返回错误
			}

			// 同时更新图片表中的图片类别（将封面图片类别重置）
			if err := utils.InitImagesCategory(tx, []string{articleToDelete.Cover}); err != nil {
				return err // 更新封面图片类别失败，返回错误
			}

			// 查找文章内容中的插图
			illustrations, err := utils.FindIllustrations(articleToDelete.Content)
			if err != nil {
				return err // 查找插图失败，返回错误
			}

			// 更新插图的类别（重置为默认类别）
			if err := utils.InitImagesCategory(tx, illustrations); err != nil {
				return err // 更新插图类别失败，返回错误
			}
			comments ,err := ServiceGroupApp.CommentService.CommentInfoByArticleID(request.CommentInfoByArticleID{ArticleID: articleToDelete.ID})
			if err != nil {
				return err // 获取评论失败，返回错误
			}
			for _, comment := range comments {
			if err := ServiceGroupApp.CommentService.DeleteCommentAndChildren(tx, comment.ID); err != nil {
				return err // 删除评论失败，返回错误
			}
		}

		}
		
		// 执行实际的文章删除操作
		return articleService.Delete(req.IDs)
	})
}

// ArticleUpdate 更新文章方法
// 功能：更新文章信息，并同步更新相关的分类、标签和图片数据
// 参数:
//   req: 更新文章的请求信息，包含文章 ID、标题、封面、分类、标签、摘要、内容等
// 返回值:
//   error: 操作过程中的错误信息
func (articleService *ArticleService) ArticleUpdate(req request.UpdateArticle) error {
	// 获取当前时间，用于设置文章的更新时间
	now := time.Now().Format("2006-01-02 15:04:05")
	
	// 构建要更新的文章数据结构体
	articleToUpdate := struct {
		UpdatedAt string   `json:"updated_at"` // 更新时间
		Cover     string   `json:"cover"`      // 文章封面
		Title     string   `json:"title"`      // 文章标题
		Keyword   string   `json:"keyword"`    // 关键词（使用标题作为关键词）
		Category  string   `json:"category"`   // 文章分类
		Tags      []string `json:"tags"`       // 文章标签
		Abstract  string   `json:"abstract"`   // 文章摘要
		Content   string   `json:"content"`    // 文章内容
	}{
		UpdatedAt: now,      // 设置更新时间为当前时间
		Cover:     req.Cover, // 新的封面
		Title:     req.Title, // 新的标题
		Keyword:   req.Title, // 关键词使用新标题
		Category:  req.Category, // 新的分类
		Tags:      req.Tags, // 新的标签
		Abstract:  req.Abstract, // 新的摘要
		Content:   req.Content, // 新的内容
	}
	
	// 开启数据库事务，确保操作的原子性
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 获取旧文章信息，用于对比和更新相关数据
		oldArticle, err := articleService.Get(req.ID)
		if err != nil {
			return err // 获取旧文章失败，返回错误
		}

		// 同时更新文章类别表中的数据（减少旧分类计数，增加新分类计数）
		if err := articleService.UpdateCategoryCount(tx, oldArticle.Category, articleToUpdate.Category); err != nil {
			return err // 更新类别计数失败，返回错误
		}

		// 同时更新文章标签表中的数据（减少旧标签计数，增加新标签计数）
		if err := articleService.UpdateTagsCount(tx, oldArticle.Tags, articleToUpdate.Tags); err != nil {
			return err // 更新标签计数失败，返回错误
		}

		// 同时更新图片表中的图片类别（如果封面有变化）
		if articleToUpdate.Cover != oldArticle.Cover {
			// 将旧封面图片类别重置
			if err := utils.InitImagesCategory(tx, []string{oldArticle.Cover}); err != nil {
				return err // 重置旧封面类别失败，返回错误
			}
			// 将新封面图片类别设置为封面类型
			if err := utils.ChangeImagesCategory(tx, []string{articleToUpdate.Cover}, appTypes.Cover); err != nil {
				return err // 设置新封面类别失败，返回错误
			}
		}
		
		// 查找旧文章内容中的插图
		oldIllustrations, err := utils.FindIllustrations(oldArticle.Content)
		if err != nil {
			return err // 查找旧插图失败，返回错误
		}
		
		// 查找新文章内容中的插图
		newIllustrations, err := utils.FindIllustrations(articleToUpdate.Content)
		if err != nil {
			return err // 查找新插图失败，返回错误
		}
		
		// 计算插图的变化（新增和删除的插图）
		addedIllustrations, removedIllustrations := utils.DiffArrays(oldIllustrations, newIllustrations)
		
		// 将删除的插图类别重置
		if err := utils.InitImagesCategory(tx, removedIllustrations); err != nil {
			return err // 重置删除插图类别失败，返回错误
		}
		
		// 将新增的插图类别设置为插图类型
		if err := utils.ChangeImagesCategory(tx, addedIllustrations, appTypes.Illustration); err != nil {
			return err // 设置新增插图类别失败，返回错误
		}

		// 执行实际的文章更新操作
		return articleService.Update(req.ID, articleToUpdate)
	})
}

// ArticleList 获取文章列表方法
// 功能：根据条件查询文章列表，支持分页和筛选
// 参数:
//   info: 查询条件，包含标题、简介、类别等筛选条件和分页信息
// 返回值:
//   list: 文章列表数据
//   total: 符合条件的总记录数
//   err: 操作过程中的错误信息
func (articleService *ArticleService) ArticleList(info request.ArticleList) (list interface{}, total int64, err error) {
	// 创建 Elasticsearch 搜索请求
	req := &search.Request{
		Query: &types.Query{},
	}

	// 构建布尔查询
	boolQuery := &types.BoolQuery{}

	// 根据标题查询（如果提供了标题条件）
	if info.Title != nil {
		boolQuery.Must = append(boolQuery.Must, types.Query{
			Match: map[string]types.MatchQuery{
				"title": {Query: *info.Title}, // 匹配文章标题
			},
		})
	}

	// 根据简介查询（如果提供了简介条件）
	if info.Abstract != nil {
		boolQuery.Must = append(boolQuery.Must, types.Query{
			Match: map[string]types.MatchQuery{
				"abstract": {Query: *info.Abstract}, // 匹配文章简介
			},
		})
	}

	// 根据类别筛选（如果提供了类别条件）
	if info.Category != nil {
		boolQuery.Filter = []types.Query{
			{
				Term: map[string]types.TermQuery{
					"category": {Value: info.Category}, // 精确匹配文章类别
				},
			},
		}
	}

	// 根据条件执行查询
	if boolQuery.Must != nil || boolQuery.Filter != nil {
		// 如果有查询条件，使用布尔查询
		req.Query.Bool = boolQuery
	} else {
		// 如果没有查询条件，使用匹配所有文档的查询
		req.Query.MatchAll = &types.MatchAllQuery{}
		// 默认按创建时间降序排序（最新的文章在前）
		req.Sort = []types.SortCombinations{
			types.SortOptions{
				SortOptions: map[string]types.FieldSort{
					"created_at": {Order: &sortorder.Desc},
				},
			},
		}
	}

	// 构建 Elasticsearch 分页选项
	option := other.EsOption{
		PageInfo: info.PageInfo,       // 分页信息
		Index:    elasticsearch.ArticleIndex(), // 文章索引名
		Request:  req,                 // 搜索请求
	}
	// 执行 Elasticsearch 分页搜索并返回结果
	return utils.EsPagination(context.TODO(), option)
}