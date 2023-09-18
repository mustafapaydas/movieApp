package Entity

import "gorm.io/gorm"

type Award struct {
	gorm.Model
	Name                string               `gorm:"not null" json:"name"`
	Institution         string               `json:"institution"`
	MovieAwardRelations []MovieAwardRelation `json:"movieAwardRelations"`
}

func (Award) TableName() string {
	return "tbl_award"
}
