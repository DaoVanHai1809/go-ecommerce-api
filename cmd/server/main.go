package main

import (
	"go-ecommerce-backend-api/internal/routers"
)

func main() {
  var r = routers.NewRouter()
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
