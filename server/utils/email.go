package utils

import (
	"crypto/tls"
	"fmt"
	"go_blog/server/global"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

// Email 发送电子邮件
func Email(To, subject string, body string) error {
	to := strings.Split(To, ",") // 将收件人邮箱地址按逗号拆分成多个地址
	return send(to, subject, body)
}

// send 发送电子邮件
// 该函数用于向指定收件人发送邮件，支持带昵称的发件人地址和 SSL 加密
// 参数：
//   - to: []string，收件人邮箱地址列表
//   - subject: string，邮件主题
//   - body: string，邮件内容（HTML 格式）
// 返回值：
//   - error: 如果发送过程中出现错误，返回错误信息；否则返回 nil
// 执行流程：
//   1. 获取全局配置中的邮件设置（发件人、昵称、密码、服务器地址、端口、是否使用 SSL）
//   2. 使用 PlainAuth 创建 SMTP 认证信息
//   3. 创建新的电子邮件对象
//   4. 设置发件人地址（支持昵称格式："昵称 <邮箱>"）
//   5. 设置收件人、主题和邮件内容
//   6. 构建邮件服务器地址（host:port）
//   7. 根据配置选择是否使用 SSL 发送邮件
//   8. 返回发送结果（成功或错误）
// 注意：
//   - 该函数依赖全局配置中的邮件设置，需要确保配置正确
//   - 邮件内容应为 HTML 格式，会被转换为 []byte 类型
func send(to []string, subject string, body string) error {
	// 获取全局配置中的邮件设置
	emailCfg := global.Config.Email // 获取全局配置中的邮件设置

	from := emailCfg.From
	nickname := emailCfg.Nickname
	secret := emailCfg.Secret
	host := emailCfg.Host
	port := emailCfg.Port
	isTLS := emailCfg.IsTLS

	// 使用 PlainAuth 创建认证信息
	// 第一个参数为空字符串，表示不需要身份验证标识
	// 第二个参数为发件人邮箱
	// 第三个参数为邮箱密码/密钥
	// 第四个参数为 SMTP 服务器地址
	auth := smtp.PlainAuth("", from, secret, host)

	// 创建新的电子邮件对象
	e := email.NewEmail()
	if nickname != "" {
		// 如果设置了昵称，则格式化发件人地址为 "昵称 <邮箱>"
		e.From = fmt.Sprintf("%s <%s>", nickname, from)
	} else {
		// 否则直接使用发件人邮箱
		e.From = from
	}

	// 设置收件人、主题和邮件内容
	e.To = to            // 收件人邮箱地址列表
	e.Subject = subject  // 邮件主题
	e.HTML = []byte(body) // 邮件内容（HTML 格式）

	// 定义错误变量
	var err error
	// 构建邮件服务器的地址，格式为 host:port
	hostAddr := fmt.Sprintf("%s:%d", host, port)

	// 根据配置的是否使用 TLS 来选择邮件发送方法
	if isTLS {
		// 使用带 TLS 的邮件发送，配置服务器名称
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		// 使用普通的邮件发送
		err = e.Send(hostAddr, auth)
	}

	return err
}