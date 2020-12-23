package tool

import (
	"github.com/micro/go-micro/v2/server"
	limiter "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
)

var QPS int = 100

//设置限流器参数
func setQPS(qps int) {
	QPS = qps
}

//获取限流器
func getLimiter() *server.HandlerWrapper {
	limite := limiter.NewHandlerWrapper(QPS)
	return &limite
}
