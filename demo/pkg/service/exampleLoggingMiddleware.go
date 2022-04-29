package service

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"go.uber.org/zap"
	"helloWorld/pkg/model/example/dto/request"
	"helloWorld/pkg/model/example/entity"
	"time"
)

type ExampleLoggingMiddleware struct {
	Next       IExampleService
	log        *zap.Logger
	consoleLog log.Logger
}

func (mw ExampleLoggingMiddleware) Check(trx request.ExampleReq) (output []entity.Example, err error) {
	defer func(begin time.Time) {
		_ = mw.consoleLog.Log(
			"method", "Check",
			"input", fmt.Sprintf("%#v", trx),
			"output", fmt.Sprintf("%#v", output),
			"err", err,
			"took", time.Since(begin))
	}(time.Now())
	output, err = mw.Next.Check(trx)
	mw.log.Debug(fmt.Sprintf("Request: %v\nResponse: %v", trx, output))
	return
}
