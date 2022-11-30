package actionModel

import (
	"gorm.io/gorm"
	"pood/v2/app/models/unitModel"
)

type Action struct {
	ID       uint           `json:"id" gorm:"primaryKey"`
	Name     string         `json:"name"`
	Type     uint           `json:"type"`
	Subtype  uint           `json:"subtype"`
	Unit     uint           `json:"unit"`
	UnitInfo unitModel.Unit `json:"unit_info" gorm:"foreignKey:Unit"`
}

func (a *Action) BeforeCreate(tx *gorm.DB) (err error) {
	if a.Type == 0 {
		a.Type = 1
	}
	return
}

type ActionCreateRequest struct {
	Name    string `json:"name"`
	Type    uint   `json:"type"`
	Subtype uint   `json:"subtype"`
	Unit    uint   `json:"unit"`
}
