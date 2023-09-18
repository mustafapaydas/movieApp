package main

import (
	"github.com/gin-gonic/gin"
	"movieProject/api"
	"movieProject/config"
)

var r = gin.Default()

func init() {
	config.InitDbConfig()
	restApi := r.Group("/api")
	{
		api.AddRoute(restApi)
	}

}
func main() {
	r.Run(":8081")
}
