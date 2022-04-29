package api

import (
	xconv "github.com/howcrazy/xconv"
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"helloWorld/pkg/model/example/dto/request"
	"helloWorld/pkg/model/example/dto/response"
	"helloWorld/pkg/model/example/entity"
	"helloWorld/pkg/service"
)

type IExampleApi interface {
	Check(*zap.Logger)
}
type ExampleApi struct {
	ExampleService service.IExampleService
	log            *zap.Logger
}

func NewExampleApi(s service.IExampleService, log *zap.Logger) ExampleApi {
	return ExampleApi{ExampleService: s, log: log}
}

// Check godoc
// @Summary Check endpoint of go example app
// @Description Check Service details
// @Tags ExampleApi
// @Accept json
// @Produce json
// @Param        account  body      request.ExampleReq  true  "ExampleReq"
// @Success 200 {object} response.ExampleResp
// @Router /api/v1/check [post]
func (p ExampleApi) Check(c echo.Context) error {
	var req request.ExampleReq
	if err := c.Bind(&req); err != nil {
		return err
	}
	if checkEntity, err := p.ExampleService.Check(req); err != nil {
		return err
	} else {
		return c.JSON(fasthttp.StatusOK, ofSlice(checkEntity))
	}
}

func ofSlice(examples []entity.Example) []response.ExampleResp {
	ls := make([]response.ExampleResp, len(examples))
	for i := 0; i < len(examples); i++ {
		xconv.Convert(examples[i], &ls[i])
	}
	return ls
}
