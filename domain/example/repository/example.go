package repository

import (
	"github.com/arisatriop/go-auto-reload/config"
	"github.com/arisatriop/go-auto-reload/domain/example/delivery"
	"github.com/arisatriop/go-auto-reload/domain/example/entity"
	"github.com/arisatriop/goerr"
	"github.com/gin-gonic/gin"
)

type ExampleRepository interface {
	Create(ctx *gin.Context, request *delivery.ExampleCreateRequest) *goerr.Error
	Update(ctx *gin.Context, request *delivery.ExampleUpdateRequest) *goerr.Error
	Delete(ctx *gin.Context, id string) *goerr.Error
	FindById(ctx *gin.Context, id string) (*entity.Example, *goerr.Error)
	FindAll(ctx *gin.Context, request *delivery.ExampleReadRequest) ([]entity.Example, *goerr.Error)
	FindByUuid(ctx *gin.Context, uuid string) (*entity.Example, *goerr.Error)
	FindByCode(ctx *gin.Context, code string) (*entity.Example, *goerr.Error)
	Count(ctx *gin.Context, request *delivery.ExampleReadRequest) (int, *goerr.Error)
}

type ExampleRepositoryImpl struct {
	DB *config.DB
}

func (r *ExampleRepositoryImpl) Create(ctx *gin.Context, request *delivery.ExampleCreateRequest) *goerr.Error {
	panic("Not implement")
}

func (r *ExampleRepositoryImpl) Update(ctx *gin.Context, request *delivery.ExampleUpdateRequest) *goerr.Error {
	panic("Not implement")
}

func (r *ExampleRepositoryImpl) Delete(ctx *gin.Context, id string) *goerr.Error {
	panic("Not implement")
}

func (r *ExampleRepositoryImpl) FindById(ctx *gin.Context, id string) (*entity.Example, *goerr.Error) {
	panic("Not implement")
}

func (r *ExampleRepositoryImpl) FindAll(ctx *gin.Context, request *delivery.ExampleReadRequest) ([]entity.Example, *goerr.Error) {
	panic("Not implement")
}

func (r *ExampleRepositoryImpl) FindByUuid(ctx *gin.Context, uuid string) (*entity.Example, *goerr.Error) {
	panic("Not implement")
}

func (r *ExampleRepositoryImpl) FindByCode(ctx *gin.Context, code string) (*entity.Example, *goerr.Error) {
	panic("Not implement")
}

func (r *ExampleRepositoryImpl) Count(ctx *gin.Context, request *delivery.ExampleReadRequest) (int, *goerr.Error) {
	panic("Not implement")
}

func NewExampleRepository(db *config.DB) ExampleRepository {
	return &ExampleRepositoryImpl{
		DB: db,
	}
}
