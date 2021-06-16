package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gogogo/config"
	"gogogo/middlewares"
	"gogogo/router"
	"gogogo/share"
	"gogogo/util"
)

func startDb() {
	db, err := util.CreateDbConn(config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_DB)
	if err != nil {
		fmt.Printf("init db failed, err: %v", err)
	}
	share.DbConn = db
}

func startServer() {
	r := gin.Default()
	r.Use(middlewares.Cors)
	router.DemoRoute(r)
	r.Run(fmt.Sprintf("%s:%s", config.SERVER_HOST, config.SERVER_PORT))
}

func Run() {
	startDb()
	startServer()
}
