package env

import "os"

var (
	AppName    = "unknown"
	AppVersion = "v0.0.1"
)

func init() {
	// 应用名称 发布平台通过环境变量注入
	AppName = os.Getenv("APP_NAME")
	// 应用版本 发布平台通过环境变量注入 通常是发布的镜像名称
	AppVersion = os.Getenv("APP_VERSION")
}
