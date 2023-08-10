package delivery

import (
	"github.com/arisatriop/go-auto-reload/domain/example/delivery/api"
	"github.com/arisatriop/go-auto-reload/domain/example/usecase"
	"github.com/arisatriop/golog"
	"github.com/gin-gonic/gin"
)

type ExampleHandler interface {
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
	Find() gin.HandlerFunc
	FindAll() gin.HandlerFunc
}

type ExampleHanlderImpl struct {
	Request  api.ExampleRequest
	Response api.ExampleResponse
	Usecase  usecase.ExampleUsecase
}

func (h *ExampleHanlderImpl) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		activity, user := "Create example", ""
		go golog.LogStart(activity, user)

		// do something

		golog.LogEnd(activity, user)
	}
}

func (h *ExampleHanlderImpl) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		activity, user := "Update example", ""
		go golog.LogStart(activity, user)

		// do something

		golog.LogEnd(activity, user)
	}
}

func (h *ExampleHanlderImpl) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		activity, user := "Delete example", ""
		go golog.LogStart(activity, user)

		// do something

		golog.LogEnd(activity, user)
	}
}

func (h *ExampleHanlderImpl) Find() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		activity, user := "Find example", ""
		go golog.LogStart(activity, user)

		// do something

		golog.LogEnd(activity, user)
	}
}

func (h *ExampleHanlderImpl) FindAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		activity, user := "Find example", ""
		go golog.LogStart(activity, user)

		// do something

		golog.LogEnd(activity, user)
	}
}

func NewExampleHandler(request api.ExampleRequest, response api.ExampleResponse, usecase usecase.ExampleUsecase) ExampleHandler {
	return &ExampleHanlderImpl{
		Request: request,
	}
}
