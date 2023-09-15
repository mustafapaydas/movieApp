package Entity

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Name        string `gorm:"column:name;not null"`
	Year        int
	Description string
}

func (Movie) TableName() string {
	return "tbl_movie"
}
