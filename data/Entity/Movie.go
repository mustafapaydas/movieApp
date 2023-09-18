package Entity

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Name                string `gorm:"column:name;not null"`
	Year                int
	Description         string
	MovieAwardRelations []MovieAwardRelation
	Stars               []Star `gorm:"many2many:tbl_movie_star_relation"`
}

func (Movie) TableName() string {
	return "tbl_movie"
}
