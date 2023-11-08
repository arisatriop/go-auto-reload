package delivery

import (
	"github.com/arisatriop/goilerplate-mono/config"
	"github.com/arisatriop/goilerplate-mono/domain/example/delivery/api"
	"github.com/arisatriop/goilerplate-mono/domain/example/repository"
	"github.com/arisatriop/goilerplate-mono/domain/example/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Route struct{}

func (r *Route) Init(e *gin.Engine, validator *validator.Validate, db *config.DB, endpoint string) {
	request := api.NewExampleRequest()
	response := api.NewExampleResponse()
	repository := repository.NewExampleRepository(db)
	usecase := usecase.NewExampleUsecase(validator, repository)
	handler := NewExampleHandler(request, response, usecase)

	e.POST("api/examples", handler.Create())
	e.PUT("api/examples/:param", handler.Update())
	e.DELETE("api/examples/", handler.Delete())
	e.GET("api/examples/:param", handler.Find())
	e.GET("api/examples", handler.FindAll())

}
