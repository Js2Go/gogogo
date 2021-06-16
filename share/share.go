package share

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

var (
	DbConn *sql.DB
	Router *gin.Engine
)
