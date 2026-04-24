package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"go_blog/server/global"
	"strings"
	"time"
)

// GinLogger 是一个 Gin 中间件，用于记录请求日志。
// 该中间件会在每次请求结束后，使用 Zap 日志记录请求信息。
// 通过此中间件，可以方便地追踪每个请求的情况以及性能。
// GinLogger 是一个 Gin 中间件函数，用于记录 HTTP 请求的详细信息
// 该中间件会在请求处理前后记录相关信息，包括请求方法、路径、状态码、耗时等
// 返回值：
//   - gin.HandlerFunc：Gin 框架的中间件函数类型，用于注册到路由中
// 执行流程：
//   1. 记录请求开始时间
//   2. 获取请求的路径和查询参数
//   3. 执行后续的处理（调用 c.Next()）
//   4. 计算请求处理的耗时
//   5. 使用 Zap 日志库记录请求的详细信息
// 日志字段说明：
//   - status：HTTP 响应状态码
//   - method：HTTP 请求方法（GET、POST、PUT 等）
//   - path：请求路径
//   - query：URL 查询参数
//   - ip：客户端 IP 地址
//   - user-agent：客户端浏览器或设备信息
//   - errors：请求处理过程中的错误信息（如果有）
//   - cost：请求处理的耗时
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求开始时间，用于计算处理耗时
		start := time.Now()

		// 获取请求的路径和查询参数，用于日志记录
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// 继续执行后续的处理，包括其他中间件和路由处理函数
		c.Next()

		// 计算请求处理的耗时
		cost := time.Since(start)

		// 使用 Zap 记录请求日志，包含详细的请求信息
		global.Log.Info(path,
			// 记录响应状态码
			zap.Int("status", c.Writer.Status()),
			// 记录请求方法
			zap.String("method", c.Request.Method),
			// 记录请求路径
			zap.String("path", path),
			// 记录查询参数
			zap.String("query", query),
			// 记录客户端 IP 地址
			zap.String("ip", c.ClientIP()),
			// 记录 User-Agent 信息，用于识别客户端类型
			zap.String("user-agent", c.Request.UserAgent()),
			// 记录错误信息（如果有），只记录私有错误
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			// 记录请求处理的耗时
			zap.Duration("cost", cost),
		)
	}
}
// GinRecovery 是一个 Gin 中间件，用于捕获和处理请求中的 panic 错误。
// 该中间件的主要作用是确保服务在遇到未处理的异常时不会崩溃，并通过日志系统提供详细的错误追踪。
// GinRecovery 是一个 Gin 中间件函数，用于捕获并处理请求处理过程中的 panic
// 该中间件可以防止服务器因未捕获的 panic 而崩溃，同时记录详细的错误信息
// 参数：
//   - stack: bool，是否记录完整的堆栈信息
// 返回值：
//   - gin.HandlerFunc：Gin 框架的中间件函数类型，用于注册到路由中
// 执行流程：
//   1. 使用 defer 确保 panic 被捕获，并且处理函数会在 panic 后执行
//   2. 检查是否发生了 panic 错误
//   3. 处理特殊情况：连接被断开的错误（如 broken pipe）
//   4. 获取请求信息，用于日志记录
//   5. 根据错误类型和 stack 参数记录不同级别的日志
//   6. 对于非 broken pipe 错误，返回 500 错误状态码
//   7. 继续执行后续的请求处理
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 使用 defer 确保 panic 被捕获，并且处理函数会在 panic 后执行
		defer func() {
			// 检查是否发生了 panic 错误
			if err := recover(); err != nil {
				// 检查是否是连接被断开的问题（如 broken pipe），这些错误不需要记录堆栈信息
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				// 获取请求信息，包括请求体等
				httpRequest, _ := httputil.DumpRequest(c.Request, false)

				// 如果是 broken pipe 错误，则只记录错误信息，不记录堆栈信息
				if brokenPipe {
					global.Log.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// 由于连接断开，不能再向客户端写入状态码
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()                // 中止请求处理
					return
				}

				// 如果是其他类型的 panic，根据 `stack` 参数决定是否记录堆栈信息
				if stack {
					// 记录详细的错误和堆栈信息
					global.Log.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					// 只记录错误信息，不记录堆栈
					global.Log.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				// 返回 500 错误状态码，表示服务器内部错误
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		// 继续执行后续的请求处理
		c.Next()
	}
}