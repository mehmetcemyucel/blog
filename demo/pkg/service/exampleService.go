package service

import (
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"helloWorld/pkg/model/example/dto/request"
	"helloWorld/pkg/model/example/entity"
	"helloWorld/pkg/repository"
)

type IExampleService interface {
	Check(trx request.ExampleReq) ([]entity.Example, error)
}

type ExampleService struct {
	ExampleRepository repository.IExampleRepository
	log               *zap.Logger
}

func NewExampleService(p repository.IExampleRepository, logger *zap.Logger) IExampleService {
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "example_group",
		Subsystem: "example_service",
		Name:      "request_count",
		Help:      "Number of requests recieved",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "example_group",
		Subsystem: "example_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of request in microseconds",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "example_group",
		Subsystem: "example_service",
		Name:      "count_result",
		Help:      "The result of each count method",
	}, []string{})

	var exampleService IExampleService
	exampleService = &ExampleService{p, logger}
	//exampleService = ExampleLoggingMiddleware{exampleService, logger, log.NewLogfmtLogger(os.Stderr)}
	exampleService = ExampleInstrumentingMiddleware{requestCount, requestLatency, countResult, exampleService}
	return exampleService
}

func (s *ExampleService) Check(trx request.ExampleReq) ([]entity.Example, error) {
	if examples, err := s.ExampleRepository.ActiveExamples(trx.Name); err != nil {
		s.log.Error("examples cannot loaded, ", zap.Error(err))
		return nil, err
	} else {
		return examples, err
	}
}
