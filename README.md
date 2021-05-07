# requestid

[![GoDoc](https://godoc.org/github.com/things-go/requestid?status.svg)](https://godoc.org/github.com/things-go/requestid)

requestid is an requestId(traceId) middleware for [Gin](https://github.com/gin-gonic/gin)

## Format 
    `hostname-pid-initrandvalue-sequence`

## Installation

```bash
    go get github.com/things-go/requestid
```

## Simple Example

[embedmd]:# (_example/main.go go)
```go
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
```

## License

This project is under MIT License. See the [LICENSE](LICENSE) file for the full license text.
