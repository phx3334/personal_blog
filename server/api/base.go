package api

import (
	"go_blog/server/global"
	"go_blog/server/model/request"
	"go_blog/server/model/response"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type BaseApi struct {
}

// store 验证码存储对象，用于管理验证码的生成和验证
// 类型：base64Captcha.Store 接口实现
// 作用：
//   1. 存储验证码的正确答案，关联唯一的 CaptchaID
//   2. 提供验证码验证功能，检查用户输入是否正确
//   3. 自动管理验证码过期，默认有效期为 10 分钟
//   4. 支持线程安全的并发访问
//
// 实现：
//   - 使用 base64Captcha 包的默认内存存储实现
//   - 适合单服务器部署场景
//   - 对于分布式部署，需要使用 Redis 等分布式存储
//
// 工作流程：
//   1. 生成验证码时，captcha.Generate() 会将正确答案存储到 store 中
//   2. 验证验证码时，通过 store.Verify(CaptchaID, 用户输入, true) 验证
//   3. true 参数表示验证后删除该验证码，防止重复使用
//
// 使用场景：
//   - 登录、注册时的图形验证码验证
//   - 发送邮箱验证码前的图形验证码验证
//   - 其他需要防止机器人攻击的场景
var store = base64Captcha.DefaultMemStore




// Captcha 生成数字验证码
// 该方法用于生成图形验证码，包括创建验证码驱动、生成验证码图片和返回验证码信息
// 参数：
//   - c: *gin.Context，Gin 上下文，包含请求和响应信息
// 执行流程：
//   1. 根据配置创建数字验证码驱动，设置验证码的高度、宽度、长度等参数
//   2. 使用驱动和存储创建验证码对象
//   3. 生成验证码，获取验证码ID、base64编码的图片和错误信息
//   4. 处理错误：如果生成失败，记录错误日志并返回失败响应
//   5. 成功时，返回包含验证码ID和图片的成功响应
func (baseApi *BaseApi) Captcha(c *gin.Context) {
	// 创建数字验证码的驱动，使用配置中的参数
	// 参数说明：
	//   - global.Config.Captcha.Height: 验证码图片高度
	//   - global.Config.Captcha.Width: 验证码图片宽度
	//   - global.Config.Captcha.Length: 验证码字符长度
	//   - global.Config.Captcha.MaxSkew: 验证码字符最大倾斜度
	//   - global.Config.Captcha.DotCount: 验证码图片上的干扰点数量
	driver := base64Captcha.NewDriverDigit(
		global.Config.Captcha.Height,
		global.Config.Captcha.Width,
		global.Config.Captcha.Length,
		global.Config.Captcha.MaxSkew,
		global.Config.Captcha.DotCount,
	)



	// 创建验证码对象，使用驱动和内存存储
	// store 是全局变量，在包初始化时定义为 base64Captcha.DefaultMemStore
	captcha := base64Captcha.NewCaptcha(driver, store)

	// 生成验证码
	// 返回值：
	//   - id: 验证码唯一标识符
	//   - b64s: 验证码图片的 base64 编码
	//   - answer: 验证码的正确答案（此处未使用）
	//   - err: 错误信息
	id, b64s, _, err := captcha.Generate()

	// 处理生成验证码时的错误
	if err != nil {
		// 记录错误日志
		global.Log.Error("Failed to generate captcha:", zap.Error(err))
		// 返回失败响应
		response.FailWithMessage("Failed to generate captcha", c)
		return
	}
	
	// 生成成功，返回包含验证码信息的成功响应
	// response.Captcha 结构体包含验证码ID和图片路径
	response.OkWithData(response.Captcha{
		CaptchaID: id,  // 验证码唯一标识符，用于后续验证
		PicPath:   b64s, // 验证码图片的 base64 编码，前端可以直接显示
	}, c)
}



// SendEmailVerificationCode 发送邮箱验证码
// 该方法用于处理用户请求发送邮箱验证码的逻辑，包括验证图形验证码和发送邮箱验证码
// 参数：
//   - c: *gin.Context，Gin 上下文，包含请求和响应信息
// 执行流程：
//   1. 从请求体中绑定 JSON 数据到 request.SendEmailVerificationCode 结构体
//   2. 验证请求数据绑定是否成功
//   3. 验证图形验证码是否正确
//   4. 如果验证码正确，调用服务层发送邮箱验证码
//   5. 根据操作结果返回成功或失败的响应
func (baseApi *BaseApi) SendEmailVerificationCode(c *gin.Context) {
	// 定义请求结构体，用于绑定 JSON 数据
	var req request.SendEmailVerificationCode
	// 从请求体中绑定 JSON 数据到结构体
	err := c.ShouldBindJSON(&req)
	// 如果绑定失败，返回错误信息
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 验证图形验证码是否正确
	if store.Verify(req.CaptchaID, req.Captcha, true) {
		// 调用服务层发送邮箱验证码
		err = baseService.SendEmailVerificationCode(c, req.Email)
		// 如果发送失败，记录错误日志并返回失败响应
		if err != nil {
			global.Log.Error("Failed to send email:", zap.Error(err))
			response.FailWithMessage("Failed to send email", c)
			return
		}
		// 发送成功，返回成功响应
		response.OkWithMessage("Successfully sent email", c)
		return
	}
	// 验证码不正确，返回错误信息
	response.FailWithMessage("Incorrect verification code", c)
}
// QQLoginURL 返回 QQ 登录链接
func (baseApi *BaseApi) QQLoginURL(c *gin.Context) {
	url := global.Config.QQ.QQLoginURL()
	response.OkWithData(url, c)
}
