package Entity

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Name                string               `gorm:"not null;default: null" json:"name"`
	Year                int                  `gorm:"not null" json:"year"`
	Description         string               `json:"description"`
	MovieAwardRelations []MovieAwardRelation `json:"movieAwardRelations"`
	Stars               []Star               `gorm:"many2many:tbl_movie_star_relation" json:"stars"`
}

func (Movie) TableName() string {
	return "tbl_movie"
}
