package api

import (
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		getMovieApi(v1)
	}
}
