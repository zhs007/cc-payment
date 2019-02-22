package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhs007/cc-payment/controller/api"
)

// Router -
var Router *gin.Engine

func init() {
	Router = gin.Default()
}

// SetRouter -
func SetRouter() {

	apiGroup := Router.Group("/api")
	apiGroup.GET("/accounts", api.GetAccounts)
}
