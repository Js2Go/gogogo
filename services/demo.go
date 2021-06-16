package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gogogo/models"
	"gogogo/share"
	"net/http"
)

func DemoGet(ctx *gin.Context) {
	id := ctx.Query("id")
	stmt, err := share.DbConn.Prepare("select * from websites where id = ?")
	if err != nil {
		fmt.Println(err)
	}
	
	row := stmt.QueryRow(id)
	var ws models.Website
	row.Scan(&ws.ID, &ws.Name, &ws.Url, &ws.Alexa, &ws.Country)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": ws,
	})
	ctx.Next()
}
