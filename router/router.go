package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhs007/cc-payment/controller/api"
)

// GinEngine -
var GinEngine *gin.Engine

func init() {
	GinEngine = gin.Default()
}

// SetRouter -
func SetRouter() {

	apiGroup := GinEngine.Group("/api")
	apiGroup.GET("/accounts", api.GetAccounts)
	apiGroup.POST("/pay", api.Pay)
	apiGroup.GET("/payments", api.GetPayments)
}
