package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Cors(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, HEAD")
	ctx.Next()
}
