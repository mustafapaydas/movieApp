package api

import (
	"github.com/gin-gonic/gin"
)

//type Group struct {
//	Group     *gin.RouterGroup
//	GroupName string
//}
//
//var v1 = gin.RouterGroup{}
//
//func AddRoute(group Group) {
//	v1.Group(fmt.Sprintf("/%s", group.GroupName))
//}
//func SetRoutes(group *gin.RouterGroup) {
//	return
//}

func AddRoute(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		GetMovieApi(v1)
	}
}
