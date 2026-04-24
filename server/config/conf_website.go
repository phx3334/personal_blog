package config

// Website 网站信息
type Website struct {
	Logo                 string `json:"logo" yaml:"logo"`
	FullLogo             string `json:"full_logo" yaml:"full_logo"`
	Title                string `json:"title" yaml:"title"`                                   // 网站标题
	Slogan               string `json:"slogan" yaml:"slogan"`                                 // 网站标语
	SloganEn             string `json:"slogan_en" yaml:"slogan_en"`                           // 英文标语
	Description          string `json:"description" yaml:"description"`                       // 网站描述
	Version              string `json:"version" yaml:"version"`                               // 网站版本
	CreatedAt            string `json:"created_at" yaml:"created_at"`                         // 创建时间
	IcpFiling            string `json:"icp_filing" yaml:"icp_filing"`                         // ICP 备案
	PublicSecurityFiling string `json:"public_security_filing" yaml:"public_security_filing"` // 公安备案
	BilibiliURL          string `json:"bilibili_url" yaml:"bilibili_url"`                     // Bilibili 链接
	GiteeURL             string `json:"gitee_url" yaml:"gitee_url"`                           // Gitee 链接
	GithubURL            string `json:"github_url" yaml:"github_url"`                         // GitHub 链接
	Name                 string `json:"name" yaml:"name"`                                     // 昵称
	Job                  string `json:"job" yaml:"job"`                                       // 职业
	Address              string `json:"address" yaml:"address"`                               // 地址
	Email                string `json:"email" yaml:"email"`                                   // 邮箱
}