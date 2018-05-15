# gin-cross-middleware
[![Build Status](https://travis-ci.org/jayhuang75/gin-cross-middleware.svg?branch=master)](https://travis-ci.org/jayhuang75/gin-cross-middleware)
[![codecov](https://codecov.io/gh/jayhuang75/gin-cross-middleware/branch/master/graph/badge.svg)](https://codecov.io/gh/jayhuang75/gin-cross-middleware)
[![Go Report Card](https://goreportcard.com/badge/github.com/jayhuang75/gin-cross-middleware)](https://goreportcard.com/report/github.com/jayhuang75/gin-cross-middleware)

## What is Cross-domain?
Cross-domain more information: 
[https://en.wikipedia.org/wiki/Cross-domain_solution](https://en.wikipedia.org/wiki/Cross-domain_solution)

## How to use this?
#### Install package
```bash
$ go get github.com/jayhuang75/gin-cross-middleware
```

#### In your gin application main.go, import the package
```go
import (
    "github.com/jayhuang75/gin-cross-middleware"
)
```

#### Use the middleware
```go
app := gin.Default()

app.Use(cross.Middleware())
```