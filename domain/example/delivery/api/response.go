package api

import "github.com/arisatriop/go-auto-reload/domain/example/entity"

type Response struct {
	Code    string `json:"code"`
	Example string `json:"example"`
	Uuid    string `json:"uuid"`
}

type ExampleResponse interface {
	ResponseFindAll(example []entity.Example) []Response
	ResponseFindById(example *entity.Example) *Response
	ResponseFindCode(example *entity.Example) *Response
	ResponseFindByUuid(example *entity.Example) *Response
}

type ExampleResponseImpl struct{}

func (r *ExampleResponseImpl) ResponseFindAll(example []entity.Example) []Response {
	panic("Not implement")
}

func (r *ExampleResponseImpl) ResponseFindById(example *entity.Example) *Response {
	panic("Not implement")
}

func (r *ExampleResponseImpl) ResponseFindCode(example *entity.Example) *Response {
	panic("Not implement")
}

func (r *ExampleResponseImpl) ResponseFindByUuid(example *entity.Example) *Response {
	panic("Not implement")
}

func NewExampleResponse() ExampleResponse {
	return &ExampleResponseImpl{}
}
