package api

import (
	"github.com/gin-gonic/gin"
	"movieProject/service"
)

func getMovieApi(group *gin.RouterGroup) {
	var movieService = service.MovieService{}
	movie := group.Group("/movie")
	movie.GET("/", movieService.Paginator)
	movie.POST("/", movieService.Create)

}
