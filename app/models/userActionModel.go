package models

import (
	"time"
)

type UserAction struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	UserId   uint    `json:"user_id"`
	User     *User   `json:"user"`
	ActionId uint    `json:"action_id"`
	Action   *Action `json:"action"`
	Deleted  bool    `json:"deleted" gorm:"default:0"`
	Logs     []Log   `json:"logs" gorm:"foreignKey:UserActionId"`
	Private  bool    `json:"private" gorm:"default:true"`
}

type MyActionsResponse struct {
	ID     uint `json:"id"`
	UserId uint `json:"-"`
	Action struct {
		ID       uint   `json:"id" gorm:"primaryKey"`
		Name     string `json:"name"`
		Type     uint   `json:"type"`
		Subtype  uint   `json:"subtype"`
		UnitInfo struct {
			Name string `json:"name"`
		} `json:"unit_info"`
	} `json:"action"`
	Deleted bool `json:"deleted"`
	Private bool `json:"private"`
}

type MyActiveActions struct {
	ID          uint      `json:"id"`
	Action      Action    `json:"action"`
	StartTime   time.Time `json:"start_time"`
	Description *string   `json:"description"`
}

type UpdateRequest struct {
	Private *bool `json:"private"`
}

type UserActionsResponse struct {
	ID     uint `json:"id"`
	UserId uint `json:"user_id"`
	Action struct {
		ID       uint   `json:"id" gorm:"primaryKey"`
		Name     string `json:"name"`
		Type     uint   `json:"type"`
		Subtype  uint   `json:"subtype"`
		UnitInfo struct {
			Name string `json:"name"`
		} `json:"unit_info"`
	} `json:"action"`
}
