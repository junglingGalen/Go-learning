package main

import (
	_ "fmt"
	h "net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context){
		c.JSON(h.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

