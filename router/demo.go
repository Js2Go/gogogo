package router

import (
	"github.com/gin-gonic/gin"
	"gogogo/services"
)

func DemoRoute(r *gin.Engine) {
	demoGroup := r.Group("/demo")
	{
		demoGroup.GET("/get", services.DemoGet)
	}
}
