package config

//auth信息
type Auth struct {
	//是否允许匿名
	Anonymous bool
	//是否通过认证
	Authenticated bool
	//用户名
	User string
	//用户密码
	Password string
	//能否写文件
	Writable bool
	//能否下载文件
	DownLoadable bool
	//远程连接的地址
	Url string
}