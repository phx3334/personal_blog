package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// GenerateVerificationCode 生成指定长度的数字验证码
// 该函数用于生成随机的数字验证码，常用于邮箱验证、短信验证等场景
// 参数：
//   - length: int，验证码的长度（位数）
// 返回值：
//   - string，生成的数字验证码字符串，长度为指定的 length
// 实现原理：
//   1. 使用当前时间的纳秒数作为随机种子，确保每次生成的验证码不同
//   2. 创建一个新的随机数生成器
//   3. 生成一个 0 到 10^length - 1 之间的随机整数
//   4. 使用 fmt.Sprintf 格式化输出，确保返回的字符串长度固定为 length，不足的前面补零
// 注意：
//   - 该函数生成的验证码只包含数字，长度固定为指定的 length
func GenerateVerificationCode(length int) string {
	// 创建随机数生成器，使用当前时间的纳秒数作为种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成 0 到 10^length - 1 之间的随机整数，然后格式化为固定长度的字符串
	// %0*d 表示输出长度为 length 的数字，不足的前面补零
	return fmt.Sprintf("%0*d", length, r.Intn(int(math.Pow10(length))))
}