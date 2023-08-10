package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseBody struct {
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseBodyWithCount struct {
	Message interface{} `json:"message"`
	Count   interface{} `json:"count"`
	Limit   interface{} `json:"limit"`
	Offset  interface{} `json:"offset"`
	Data    interface{} `json:"data,omitempty"`
}

func StatusOK(ctx *gin.Context, data interface{}, m ...interface{}) {

	res := ResponseBody{
		Message: "sukses",
		Data:    data,
	}
	if len(m) > 0 {
		res.Message = m[0]
	}

	ctx.JSON(http.StatusOK, res)
}

func StatusOKwithCount(ctx *gin.Context, data *ResponseBodyWithCount, m ...interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func StatusNotOKV2(ctx *gin.Context, err *Error) {
	res := ResponseBody{
		Message: err.Message,
		Data:    err.Data,
	}
	ctx.JSON(err.Code, res)
}
