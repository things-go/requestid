# requestid

[![GoDoc](https://godoc.org/github.com/things-go/requestid?status.svg)](https://godoc.org/github.com/things-go/requestid)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/things-go/requestid?tab=doc)
[![codecov](https://codecov.io/gh/things-go/requestid/branch/master/graph/badge.svg)](https://codecov.io/gh/things-go/requestid)
![Action Status](https://github.com/things-go/requestid/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/things-go/requestid)](https://goreportcard.com/report/github.com/things-go/requestid)
[![Licence](https://img.shields.io/github/license/things-go/requestid)](https://raw.githubusercontent.com/things-go/requestid/master/LICENSE)
[![Tag](https://img.shields.io/github/v/tag/things-go/requestid)](https://github.com/thinkgos/requestid/tags)


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
	router.Use(requestid.RequestId())
	router.GET("/", func(c *gin.Context) {
		fmt.Println(requestid.FromRequestId(c.Request.Context()))
		fmt.Println(requestid.GetRequestId(c))
	})
	router.Run(":8080")
}
```

## License

This project is under MIT License. See the [LICENSE](LICENSE) file for the full license text.
