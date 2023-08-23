package tracing

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func init() {

}

func Init(service string) (opentracing.Tracer, io.Closer) {
	cfg, err := config.FromEnv()

	cfg.ServiceName = service
	// 收集器类型 const代表 要么获取所有跟踪，要么不获取任何跟踪 全部接收 1，无接收 0
	// rateLimiting :选择每秒采样的跟踪数
	// probabilistic  百分比概率选择采样  param =0.1 代表10%的采样率
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1
	// 收集的时候是否开启日志
	cfg.Reporter.LogSpans = true

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("Error: cannot init jaeger:%v\n", err))
	}
	return tracer, closer
}
