package Entity

import "gorm.io/gorm"

type Award struct {
	gorm.Model
	Name                string
	Institution         string
	MovieAwardRelations []MovieAwardRelation
}

func (Award) TableName() string {
	return "tbl_award"
}
