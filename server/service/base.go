package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_blog/server/utils"
	"time"
)

type BaseService struct {
}



// SendEmailVerificationCode 发送邮箱验证码
// 该方法用于生成验证码并发送到指定邮箱，同时将会话信息存储到用户会话中
// 参数：
//   - c: *gin.Context，Gin 上下文，用于获取和存储会话信息
//   - to: string，接收验证码的邮箱地址
// 返回值：
//   - error: 如果发送邮件过程中出现错误，返回错误信息；否则返回 nil
// 执行流程：
//   1. 生成 6 位数字验证码
//   2. 设置验证码过期时间为当前时间后 5 分钟
//   3. 将验证码、目标邮箱和过期时间存储到用户会话中
//   4. 构建邮件主题和 HTML 格式的邮件内容
//   5. 调用 utils.Email 函数发送邮件
//   6. 返回操作结果
func (baseService *BaseService) SendEmailVerificationCode(c *gin.Context, to string) error {
	// 生成 6 位数字验证码
	verificationCode := utils.GenerateVerificationCode(6)
	// 计算验证码过期时间（当前时间后 5 分钟）
	expireTime := time.Now().Add(5 * time.Minute).Unix()

	// 将验证码、验证邮箱、过期时间存入会话中
	// 会话存储用于后续验证用户输入的验证码是否正确
	session := sessions.Default(c)
	session.Set("verification_code", verificationCode) // 存储验证码
	session.Set("email", to)                         // 存储目标邮箱
	session.Set("expire_time", expireTime)           // 存储过期时间
	_ = session.Save()                               // 保存会话

	// 构建邮件主题
	subject := "您的邮箱验证码"
	// 构建 HTML 格式的邮件内容
	body := `这里是fzsirrr的个人博客,<br/>
<br/>
你正在注册该博客的账户！为了确保你的邮箱安全，请使用以下验证码进行验证：<br/>
<br/>
验证码：[<font color="blue"><u>` + verificationCode + `</u></font>]<br/>
还在浪费时间吗，该验证码在 5 分钟内有效，请尽快食用。<br/>
<br/>
如果你没有请求此验证码，请忽略此邮件。
<br/>
如有任何疑问，请联系：<br/>
fzsirrr的徒弟: tj <br/>
邮箱: 2877712419@qq.com<br/>
<br/>
期待你在该博客上留下的足迹，一起探索未知吧！<br/>
		
<br/>`
    
	// 发送邮件（忽略返回的错误，因为函数会统一返回错误）
	_ = utils.Email(to, subject, body)
    
	// 返回 nil 表示操作成功
	return nil
}