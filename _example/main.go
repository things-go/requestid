package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/things-go/requestid"
)

func main() {
	router := gin.New()
	router.Use(requestid.RequestId())
	router.GET("/", func(c *gin.Context) {
		fmt.Println(requestid.FromRequestId(c.Request.Context()))
		fmt.Println(requestid.GetRequestId(c))
	})
	router.Run(":8080")
}
