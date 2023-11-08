package repository

import (
	"github.com/arisatriop/goerr"
	"github.com/arisatriop/goilerplate-mono/config"
	"github.com/arisatriop/goilerplate-mono/domain/example/delivery/api"
	"github.com/arisatriop/goilerplate-mono/domain/example/entity"
	"github.com/gin-gonic/gin"
)

type ExampleRepository interface {
	Create(ctx *gin.Context, request *api.ExampleCreateRequest) *goerr.Error
	Update(ctx *gin.Context, request *api.ExampleUpdateRequest) *goerr.Error
	Delete(ctx *gin.Context, id string) *goerr.Error
	FindById(ctx *gin.Context, id string) (*entity.Example, *goerr.Error)
	FindAll(ctx *gin.Context, request *api.ExampleReadRequest) ([]entity.Example, *goerr.Error)
	FindByUuid(ctx *gin.Context, uuid string) (*entity.Example, *goerr.Error)
	FindByCode(ctx *gin.Context, code string) (*entity.Example, *goerr.Error)
	Count(ctx *gin.Context, request *api.ExampleReadRequest) (int, *goerr.Error)
}

type ExampleRepositoryImpl struct {
	DB *config.DB
}

func (r *ExampleRepositoryImpl) Create(ctx *gin.Context, request *api.ExampleCreateRequest) *goerr.Error {
	panic("Not implement")
}

func (r *ExampleRepositoryImpl) Update(ctx *gin.Context, request *api.ExampleUpdateRequest) *goerr.Error {
	panic("Not implement")
}

func (r *ExampleRepositoryImpl) Delete(ctx *gin.Context, id string) *goerr.Error {
	panic("Not implement")
}

func (r *ExampleRepositoryImpl) FindById(ctx *gin.Context, id string) (*entity.Example, *goerr.Error) {
	panic("Not implement")
}

func (r *ExampleRepositoryImpl) FindAll(ctx *gin.Context, request *api.ExampleReadRequest) ([]entity.Example, *goerr.Error) {
	panic("Not implement")
}

func (r *ExampleRepositoryImpl) FindByUuid(ctx *gin.Context, uuid string) (*entity.Example, *goerr.Error) {
	panic("Not implement")
}

func (r *ExampleRepositoryImpl) FindByCode(ctx *gin.Context, code string) (*entity.Example, *goerr.Error) {
	panic("Not implement")
}

func (r *ExampleRepositoryImpl) Count(ctx *gin.Context, request *api.ExampleReadRequest) (int, *goerr.Error) {
	panic("Not implement")
}

func NewExampleRepository(db *config.DB) ExampleRepository {
	return &ExampleRepositoryImpl{
		DB: db,
	}
}
