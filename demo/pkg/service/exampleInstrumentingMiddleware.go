package service

import (
	"fmt"
	"github.com/go-kit/kit/metrics"
	"helloWorld/pkg/model/example/dto/request"
	"helloWorld/pkg/model/example/entity"
	"time"
)

type ExampleInstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           IExampleService
}

func (mw ExampleInstrumentingMiddleware) Check(trx request.ExampleReq) (output []entity.Example, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Check", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	output, err = mw.Next.Check(trx)
	return
}
