package service

import (
	"helloWorld/pkg/model/example/dto/request"
	"helloWorld/pkg/model/example/entity"
	"helloWorld/pkg/repository"
	"testing"
)

var service IExampleService

var repoActiveExamplesMock func(name string) ([]entity.Example, error)

type exampleRepositoryMock struct{}

func (u exampleRepositoryMock) ActiveExamples(name string) ([]entity.Example, error) {
	return repoActiveExamplesMock(name)
}

func init() {
	var repo repository.IExampleRepository
	repo = exampleRepositoryMock{}
	service = &ExampleService{repo, nil}
}

func TestCheck_whenGetRecordsWithoutError_thenReturnSlice(t *testing.T) {
	repoActiveExamplesMock = func(name string) ([]entity.Example, error) {
		return make([]entity.Example, 1), nil
	}
	resp, _ := service.Check(request.ExampleReq{})
	if resp == nil {
		t.Errorf("response should not be nil")
	}
}
