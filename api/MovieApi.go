package api

import (
	"github.com/gin-gonic/gin"
	"movieProject/Service"
)

func GetMovieApi(group *gin.RouterGroup) {
	movie := group.Group("/movie")
	movie.GET("/", Service.Paginator)

}
