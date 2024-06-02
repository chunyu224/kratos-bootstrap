package fluent

import (
	conf "github.com/chunyu224/kratos-bootstrap/api/gen/go/conf/v1"
	fluentLogger "github.com/go-kratos/kratos/contrib/log/fluent/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// NewLogger 创建一个新的日志记录器 - Fluent
func NewLogger(cfg *conf.Logger) log.Logger {
	wrapped, err := fluentLogger.NewLogger(cfg.Fluent.Endpoint)
	if err != nil {
		panic("create fluent logger failed")
		return nil
	}
	return wrapped
}
