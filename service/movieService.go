package service

import (
	"github.com/gin-gonic/gin"
	"movieProject/common/response"
	"movieProject/config"
	"movieProject/data/Entity"
	"net/http"
	"strconv"
)

func Paginator(c *gin.Context) {
	var page int
	var pageSize int
	if c.Query("page") != "" {
		page, _ = strconv.Atoi(c.Query("page"))
	} else {
		page = 0
	}
	if c.Query("size") != "" {
		pageSize, _ = strconv.Atoi(c.Query("size"))
	} else {
		pageSize = 10
	}
	var movies []Entity.Movie
	var count int64
	offset := (page - 1) * pageSize
	config.GetDB().Table("tbl_movie").Count(&count)
	config.GetDB().Offset(offset).Limit(pageSize).Model(&Entity.Movie{}).Preload("MovieAwardRelations").Preload("Stars").Find(&movies)
	c.JSON(http.StatusOK, gin.H{"page": response.PageResponse{Count: int(count), Data: movies}})
}
