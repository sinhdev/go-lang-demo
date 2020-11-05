package main

import (
	"gin-api-crud/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	controllers.SetupItemRouter(r)
	r.Run(":8080")
}
