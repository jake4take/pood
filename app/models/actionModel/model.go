package actionModel

import (
	"pood/v2/app/models/unitModel"
)

type Action struct {
	ID       uint           `json:"id" gorm:"primaryKey"`
	Name     string         `json:"name"`
	Type     uint           `json:"type" gorm:"default:1"`
	Subtype  uint           `json:"subtype"`
	UnitId   uint           `json:"unit" gorm:"default:3"`
	UnitInfo unitModel.Unit `json:"unit_info" gorm:"foreignKey:UnitId"`
}

type ActionCreateRequest struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Type    uint   `json:"type"`
	Subtype uint   `json:"subtype"`
	Unit    uint   `json:"unit"`
}

type FindActionByName struct {
	Name string `json:"name"`
}
