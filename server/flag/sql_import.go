package flag

import (
	"os"
	"go_blog/server/global"
	"strings"
)

// SQLImport 从指定的 SQL 文件导入数据到数据库
// 该函数读取 SQL 文件内容，按分号分割为单个 SQL 语句，然后逐个执行这些语句
// 参数：
//   - sqlPath: SQL 文件的路径
// 返回值：
//   - errs: 执行过程中遇到的错误列表，如果没有错误则返回 nil
func SQLImport(sqlPath string) (errs []error) {
    // 读取 SQL 文件内容
    byteData, err := os.ReadFile(sqlPath)
    if err != nil {
        // 如果读取文件失败，将错误添加到错误列表并返回
        return append(errs, err)
    }
    
    // 按分号分割 SQL 语句，得到单个 SQL 语句的列表
    sqlList := strings.Split(string(byteData), ";")
    
    // 遍历执行每个 SQL 语句
    for _, sql := range sqlList {
        // 去除字符串开头和结尾的空白符，避免空语句
        sql = strings.TrimSpace(sql)
        if sql == "" {
            // 跳过空语句
            continue
        }
        
        // 执行 SQL 语句
        err = global.DB.Exec(sql).Error
        if err != nil {
            // 如果执行失败，将错误添加到错误列表并继续执行下一条语句
            errs = append(errs, err)
            continue
        }
    }
    
    // 返回错误列表（如果没有错误则返回 nil）
    return nil
}