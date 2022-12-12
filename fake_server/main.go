package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	os.Remove("mysock.sock")
	r := gin.Default()
	r.GET("/api/v1/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, []string{
			"Jack",
			"Marry",
			"Sandy",
		})
	})
	r.POST("/api/v1/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, gin.H{
			"id":   "ABC-111",
			"name": "Jack",
		})
	})
	r.RunUnix("mysock.sock")
}
