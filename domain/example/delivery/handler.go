package delivery

import (
	"github.com/arisatriop/go-auto-reload/domain/example/usecase"
	"github.com/gin-gonic/gin"
)

type ExampleHandler interface {
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
	FindById() gin.HandlerFunc
	FindAll() gin.HandlerFunc
	FindByUuid() gin.HandlerFunc
	FindByCode() gin.HandlerFunc
}

type ExampleHanlderImpl struct {
	Usecase usecase.ExampleUsecase
}

func (h *ExampleHanlderImpl) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		panic("Not implement")
	}
}

func (h *ExampleHanlderImpl) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		panic("Not implement")
	}
}

func (h *ExampleHanlderImpl) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		panic("Not implement")
	}
}

func (h *ExampleHanlderImpl) FindById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		panic("Not implement")
	}
}

func (h *ExampleHanlderImpl) FindAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		panic("Not implement")
	}
}

func (h *ExampleHanlderImpl) FindByUuid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		panic("Not implement")
	}
}

func (h *ExampleHanlderImpl) FindByCode() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		panic("Not implement")
	}
}
