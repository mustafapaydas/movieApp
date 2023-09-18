package Entity

import (
	"gorm.io/gorm"
)

type gender string

const (
	Male   gender = "Male"
	Female gender = "Female"
)

type Star struct {
	gorm.Model
	Name     string `gorm:"not null"`
	LastName string `gorm:"not null"`
	Gender   gender
	Age      int
	Movies   []Movie `gorm:"many2many:tbl_movie_star_relation"`
}

func (Star) TableName() string {
	return "tbl_star"
}
