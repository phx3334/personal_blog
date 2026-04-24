package flag

import (
	"fmt"
	"os"
	"os/exec"
	"go_blog/server/global"
	"time"
)

// SQLExport 导出 MySQL 数据库到 SQL 文件
// 该函数通过 Docker 容器执行 mysqldump 命令，将数据库结构和数据导出到以当前日期命名的 SQL 文件中
// 返回值：如果执行过程中出现错误，返回错误信息；否则返回 nil
func SQLExport() error {
    // 从全局配置中获取 MySQL 连接信息
    mysql := global.Config.Mysql

    // 生成备份文件名，格式为 mysql_年月日.sql
    timer := time.Now().Format("20060102")  // 获取当前日期，格式为 YYYYMMDD
    sqlPath := fmt.Sprintf("mysql_%s.sql", timer)  // 构建文件名

    // 构建 Docker 命令，在 MySQL 容器中执行 mysqldump 命令
    // 参数说明：
    // - docker exec: 在运行的容器中执行命令
    // - mysql: 目标容器名称
    // - mysqldump: MySQL 导出工具
    // - -u + 用户名: 指定数据库用户名
    // - -p + 密码: 指定数据库密码
    // - mysql.DBName: 指定要导出的数据库名称
    cmd := exec.Command("docker", "exec", "mysql", "mysqldump", "-u"+mysql.Username, "-p"+mysql.Password, mysql.DBName)

    // 创建输出文件，用于存储导出的 SQL 内容
    outFile, err := os.Create(sqlPath)
    if err != nil {
        return err  // 如果创建文件失败，返回错误
    }
    defer outFile.Close()  // 确保函数结束时关闭文件，避免资源泄露

    // 将命令的标准输出重定向到创建的文件
    cmd.Stdout = outFile
    // 执行命令，将数据库导出到文件中
    return cmd.Run()
}