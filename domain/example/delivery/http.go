package delivery

import (
	"github.com/arisatriop/go-auto-reload/config"
	"github.com/arisatriop/go-auto-reload/domain/example/delivery/api"
	"github.com/arisatriop/go-auto-reload/domain/example/repository"
	"github.com/arisatriop/go-auto-reload/domain/example/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Route struct{}

func (r *Route) Init(e *gin.Engine, string, validator *validator.Validate, db *config.DB) {
	request := api.NewExampleRequest()
	response := api.NewExampleResponse()
	repository := repository.NewExampleRepository(db)
	usecase := usecase.NewExampleUsecase(validator, repository)
	handler := NewExampleHandler(request, response, usecase)

	e.POST("api/example", handler.Create())
	e.PUT("api/examples/:param", handler.Update())
	e.DELETE("api/examples/", handler.Delete())
	e.GET("api/examples/:param", handler.Find())
	e.GET("api/examples", handler.FindAll())

}
