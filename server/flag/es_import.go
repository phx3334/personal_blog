package flag

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/bulk"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
	"os"
	"go_blog/server/global"
	"go_blog/server/model/elasticsearch"
	"go_blog/server/model/other"
	"go_blog/server/service"
)

// ElasticsearchImport 从指定的 JSON 文件导入数据到 Elasticsearch 索引
// 该函数读取 JSON 文件，反序列化数据，创建 Elasticsearch 索引，然后批量导入数据
// 参数：
//   - jsonPath: JSON 文件的路径，包含要导入的 Elasticsearch 数据
// 返回值：
//   - int: 成功导入的数据总条数
//   - error: 如果执行过程中出现错误，返回错误信息；否则返回 nil
func ElasticsearchImport(jsonPath string) (int, error) {
	// 读取指定路径的 JSON 文件
	byteData, err := os.ReadFile(jsonPath)
	if err != nil {
		return 0, err  // 如果读取文件失败，返回错误
	}

	// 反序列化 JSON 数据到 ESIndexResponse 结构体
	var response other.ESIndexResponse
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		return 0, err  // 如果反序列化失败，返回错误
	}

	// 创建 Elasticsearch 索引
	esService := service.ServiceGroupApp.EsService
	indexExists, err := esService.IndexExists(elasticsearch.ArticleIndex())
	if err != nil {
		return 0, err  // 如果检查索引失败，返回错误
	}
	if indexExists {
		// 如果索引已存在，删除它
		if err := esService.IndexDelete(elasticsearch.ArticleIndex()); err != nil {
			return 0, err  // 如果删除索引失败，返回错误
		}
	}	// 创建新的索引，使用预定义的映射结构
	err = esService.IndexCreate(elasticsearch.ArticleIndex(), elasticsearch.ArticleMapping())
	if err != nil {
		return 0, err  // 如果创建索引失败，返回错误
	}

	// 构建批量请求数据
	var request bulk.Request
	for _, data := range response.Data {
		// 为每条数据创建索引操作，指定文档的 ID
		request = append(request, types.OperationContainer{Index: &types.IndexOperation{Id_: data.ID}})
		// 添加文档数据到请求
		request = append(request, data.Doc)
	}

	// 使用 Elasticsearch 客户端执行批量操作
	_, err = global.ESClient.Bulk().
		Request(&request).                   // 提交请求数据
		Index(elasticsearch.ArticleIndex()). // 指定索引名称
		Refresh(refresh.True).               // 强制刷新索引以使文档立即可见
		Do(context.TODO())                   // 执行请求
	if err != nil {
		return 0, err  // 如果批量操作失败，返回错误
	}

	// 返回导入的数据总条数
	total := len(response.Data)
	return total, nil
}