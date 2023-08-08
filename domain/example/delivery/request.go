package delivery

import (
	"log"

	"github.com/gin-gonic/gin"
)

const exampleRequestPath = "domain/example/delivery/request.go"

type ExampleReadRequest struct {
	Limit  string `validate:"required" json:"limit" form:"limit"`
	Offset string `validate:"required" json:"offset" form:"offset"`
	Search string `validate:"required" json:"example" form:"example"`
}

type ExampleCreateRequest struct {
	Example string `validate:"required" json:"example" form:"example"`
}

type ExampleUpdateRequest struct {
	Example string `validate:"required" json:"example" form:"example"`
}

type ExampleRequest interface {
	GetReadRequestPayload(ctx *gin.Context) (*ExampleReadRequest, error)
	GetCreateRequestPayload(ctx *gin.Context) (*ExampleCreateRequest, error)
	GetUpdateRequestPayload(ctx *gin.Context) (*ExampleUpdateRequest, error)
}

type ExampleRequestImpl struct{}

func NewExampleRequest() ExampleRequest {
	return &ExampleRequestImpl{}
}

func (request *ExampleRequestImpl) GetCreateRequestPayload(c *gin.Context) (*ExampleCreateRequest, error) {
	exampleCreateRequest := &ExampleCreateRequest{}
	err := c.Bind(exampleCreateRequest)
	if err != nil {
		log.Fatal(err.Error())
	}
	return exampleCreateRequest, err
}

func (request *ExampleRequestImpl) GetReadRequestPayload(c *gin.Context) (*ExampleReadRequest, error) {
	exampleReadRequest := &ExampleReadRequest{}
	err := c.Bind(exampleReadRequest)
	if err != nil {
		log.Fatal(err.Error())
	}
	return exampleReadRequest, nil
}

func (request *ExampleRequestImpl) GetUpdateRequestPayload(c *gin.Context) (*ExampleUpdateRequest, error) {
	exampleUpdateRequest := &ExampleUpdateRequest{}
	err := c.Bind(exampleUpdateRequest)
	if err != nil {
		log.Fatal(err.Error())
	}
	return exampleUpdateRequest, nil
}
