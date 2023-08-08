package delivery

import "github.com/arisatriop/go-auto-reload/domain/example/entity"

type Response struct {
	Code    string `json:"code"`
	Example string `json:"example"`
	Uuid    string `json:"uuid"`
}

type ExampleResponse interface {
	ResponseFindAll(ex []entity.Example) []Response
	ResponseFindById(ex *entity.Example) *Response
	ResponseFindCode(ex *entity.Example) *Response
	ResponseFindByUuid(ex *entity.Example) *Response
}

type ExampleResponseImpl struct{}

func (r *ExampleResponseImpl) ResponseFindAll(ex []entity.Example) []Response {
	panic("Not implement")
}

func (r *ExampleResponseImpl) ResponseFindById(ex *entity.Example) *Response {
	panic("Not implement")
}

func (r *ExampleResponseImpl) ResponseFindCode(ex *entity.Example) *Response {
	panic("Not implement")
}

func (r *ExampleResponseImpl) ResponseFindByUuid(ex *entity.Example) *Response {
	panic("Not implement")
}
