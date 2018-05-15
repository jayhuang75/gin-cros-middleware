# gin-cross-middleware

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