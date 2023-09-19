package logic

import (
	"fmt"
	"movieProject/data/Entity"
)

type MovieLogic struct {
	*AbstractLogic
}

var abstractLogic = AbstractLogic{}

func (m *MovieLogic) Create(movie Entity.Movie) (any, error) {
	if movie.ID != 0 {
		fmt.Println("\n\n\n\nccccccccccccccccccccccc\n\n\n\n")
	}
	return abstractLogic.Create(movie)
}
