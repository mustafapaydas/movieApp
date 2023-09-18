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
	Name     string  `gorm:"not null" json:"firstname"`
	LastName string  `gorm:"not null" json:"lastName"`
	Gender   gender  `json:"gender"`
	Age      int     `json:"age"`
	Movies   []Movie `gorm:"many2many:tbl_movie_star_relation" json:"movies"`
}

func (Star) TableName() string {
	return "tbl_star"
}
