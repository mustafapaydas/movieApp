package Entity

import "gorm.io/gorm"

type MovieAwardRelation struct {
	gorm.Model
	MovieID uint
	Movie   Movie
	AwardID uint
	Award   Award
	year    int `gorm:"not null"`
}

func (MovieAwardRelation) TableName() string {
	return "tbl_movie_award_relation"
}
