package api

import (
	"github.com/gin-gonic/gin"
	"movieProject/service"
)

func getMovieApi(group *gin.RouterGroup) {
	movie := group.Group("/movie")
	movie.GET("/", service.Paginator)
	//movie.POST("/", service.)

}
