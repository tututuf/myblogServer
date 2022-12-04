package main

import (
	"fmt"
	apiRouters "myblog/routers/api"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, I am Service")
	r := gin.Default()
	apiRouters.SetAdminRouter(r)

	r.Run(":8888")
}
