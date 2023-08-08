package usecase

import (
	"github.com/arisatriop/go-auto-reload/domain/example/delivery"
	"github.com/arisatriop/go-auto-reload/domain/example/repository"
	"github.com/arisatriop/goerr"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ExampleUsecase interface {
	Create(ctx *gin.Context, request *delivery.ExampleCreateRequest) *goerr.Error
	Update(ctx *gin.Context, request *delivery.ExampleUpdateRequest) *goerr.Error
	Delete(ctx *gin.Context, id string) *goerr.Error
	FindById(ctx *gin.Context, id string) (*delivery.Response, *goerr.Error)
	FindAll(ctx *gin.Context, request *delivery.ExampleReadRequest) ([]delivery.Response, *goerr.Error)
}

type ExampleUsecaseImpl struct {
	Validator  *validator.Validate
	Repository repository.ExampleRepository
}

func (u *ExampleUsecaseImpl) Create(ctx *gin.Context, request *delivery.ExampleCreateRequest) *goerr.Error {
	panic("Not implement")
}

func (u *ExampleUsecaseImpl) Update(ctx *gin.Context, request *delivery.ExampleUpdateRequest) *goerr.Error {
	panic("Not implement")
}

func (u *ExampleUsecaseImpl) Delete(ctx *gin.Context, id string) *goerr.Error {
	panic("Not implement")
}

func (u *ExampleUsecaseImpl) FindById(ctx *gin.Context, id string) (*delivery.Response, *goerr.Error) {
	panic("Not implement")
}

func (u *ExampleUsecaseImpl) FindAll(ctx *gin.Context, request *delivery.ExampleReadRequest) ([]delivery.Response, *goerr.Error) {
	panic("Not implement")
}

func NewExampleUsecase(validator *validator.Validate, repository repository.ExampleRepository) ExampleUsecase {
	return &ExampleUsecaseImpl{
		Validator:  validator,
		Repository: repository,
	}
}
