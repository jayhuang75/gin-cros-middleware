# gin-cros-middleware
## What is Cross-domain?
Cross-domain more information: 
[https://en.wikipedia.org/wiki/Cross-domain_solution](https://en.wikipedia.org/wiki/Cross-domain_solution)

## How to use this?
#### Install package
```bash
$ go get github.com/jayhuang75/gin-cros-middleware
```

#### In your gin application main.go, import the package
```go
import (
    "github.com/jayhuang75/gin-cros-middleware"
)
```

#### Use the middleware
```go
app := gin.Default()

app.Use(cros.CORSMiddleware())
```