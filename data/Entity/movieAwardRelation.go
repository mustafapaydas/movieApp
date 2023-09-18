package Entity

import "gorm.io/gorm"

type MovieAwardRelation struct {
	gorm.Model
	MovieID uint  `json:"movieID"`
	Movie   Movie `json:"movie"`
	AwardID uint  `json:"awardID"`
	Award   Award `json:"award"`
	year    int   `gorm:"not null" json:"year"`
}

func (MovieAwardRelation) TableName() string {
	return "tbl_movie_award_relation"
}
