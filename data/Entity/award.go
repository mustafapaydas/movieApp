package Entity

import "gorm.io/gorm"

type Award struct {
	gorm.Model
	Name                string `gorm:"not null"`
	Institution         string
	MovieAwardRelations []MovieAwardRelation
}

func (Award) TableName() string {
	return "tbl_award"
}
