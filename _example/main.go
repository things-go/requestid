package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/things-go/requestid"
)

func main() {
	router := gin.New()
	router.Use(requestid.RequestID())
	router.GET("/", func(c *gin.Context) {
		fmt.Println(requestid.FromRequestID(c.Request.Context()))
		fmt.Println(requestid.GetRequestID(c))
	})
	router.Run(":8080")
}
