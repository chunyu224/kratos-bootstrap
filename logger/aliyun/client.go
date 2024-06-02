package aliyun

import (
	conf "github.com/chunyu224/kratos-bootstrap/api/gen/go/conf/v1"
	aliyunLogger "github.com/go-kratos/kratos/contrib/log/aliyun/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// NewLogger 创建一个新的日志记录器 - Aliyun
func NewLogger(cfg *conf.Logger) log.Logger {
	wrapped := aliyunLogger.NewAliyunLog(
		aliyunLogger.WithProject(cfg.Aliyun.Project),
		aliyunLogger.WithEndpoint(cfg.Aliyun.Endpoint),
		aliyunLogger.WithAccessKey(cfg.Aliyun.AccessKey),
		aliyunLogger.WithAccessSecret(cfg.Aliyun.AccessSecret),
	)
	return wrapped
}
