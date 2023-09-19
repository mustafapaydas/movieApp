package service

import (
	"github.com/gin-gonic/gin"
	"movieProject/common/response"
	"movieProject/config"
	"movieProject/data/Entity"
	"movieProject/logic"
	"net/http"
	"strconv"
)

type MovieService struct {
	*AbstractService
}

func getMovieLogic() logic.MovieLogic {
	return logic.MovieLogic{}
}

var _abstractService = AbstractService{}

func (m *MovieService) Paginator(c *gin.Context) {

	var page int
	var pageSize int
	page, _ = strconv.Atoi(c.DefaultQuery("page", "0"))
	pageSize, _ = strconv.Atoi(c.DefaultQuery("size", "10"))

	var movies []Entity.Movie
	var count int64
	offset := (page - 1) * pageSize
	config.GetDB().Table("tbl_movie").Count(&count)
	config.GetDB().Offset(offset).Limit(pageSize).Model(&Entity.Movie{}).Preload("MovieAwardRelations").Preload("Stars").Find(&movies)
	c.JSON(http.StatusOK, gin.H{"page": response.PageResponse{Count: int(count), Data: movies}})
}

func (m *MovieService) Create(c *gin.Context) {

	movie := Entity.Movie{}
	err := c.BindJSON(&movie)
	res := _abstractService.Create(&movie)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if res.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, res.Error)
		return
	}
	c.JSON(http.StatusCreated, &res)
}
