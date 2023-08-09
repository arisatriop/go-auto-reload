package routes

import (
	"github.com/arisatriop/go-auto-reload/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	example "github.com/arisatriop/go-auto-reload/domain/example/delivery"
)

type RouteGroup struct {
	Public *gin.RouterGroup
}

func InitRoute(gin *gin.Engine) {
	db := config.GetDBConnection()
	validator := validator.New()

	example := example.Route{}
	example.Init(gin, validator, db, "/example")
}
