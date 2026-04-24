package flag

import (
	"bufio"
	"fmt"
	"os"
	"go_blog/server/model/elasticsearch"
	"go_blog/server/service"
)

// Elasticsearch 初始化 Elasticsearch 索引
// 该函数检查文章索引是否存在，若存在则询问用户是否删除并重建，最后创建索引
// 返回值：如果执行过程中出现错误，返回错误信息；否则返回 nil
func Elasticsearch() error {
    // 获取 Elasticsearch 服务实例
    esService := service.ServiceGroupApp.EsService

    // 检查文章索引是否已存在
    indexExists, err := esService.IndexExists(elasticsearch.ArticleIndex())
    if err != nil {
        return err  // 如果检查失败，返回错误
    }

    if indexExists {
        // 打印提示信息，询问用户是否删除并重建索引
        fmt.Println("The index already exists. Do you want to delete the data and recreate the index? (y/n)")

        // 读取用户输入
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan()
        input := scanner.Text()

        switch input {
        case "y":
            // 如果用户输入 y，删除索引
            fmt.Println("Proceeding to delete the data and recreate the index...")
            if err := esService.IndexDelete(elasticsearch.ArticleIndex()); err != nil {
                return err  // 如果删除失败，返回错误
            }
        case "n":
            // 如果用户输入 n，退出程序
            fmt.Println("Exiting the program.")
            os.Exit(0)
        default:
            // 如果用户输入无效，提示重新输入并递归调用函数
            fmt.Println("Invalid input. Please enter 'y' to delete and recreate the index, or 'n' to exit.")
            return Elasticsearch() // 递归调用，重新输入
        }
    }

    // 创建文章索引，使用预定义的映射结构
    return esService.IndexCreate(elasticsearch.ArticleIndex(), elasticsearch.ArticleMapping())
}