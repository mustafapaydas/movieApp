package api

import (
	"github.com/gin-gonic/gin"
	"movieProject/service"
)

func GetMovieApi(group *gin.RouterGroup) {
	movie := group.Group("/movie")
	movie.GET("/", service.Paginator)

}
