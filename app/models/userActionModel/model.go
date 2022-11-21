package userActionModel

import (
	"Pood/app/models/actionModel"
	"Pood/app/models/logModel"
	"Pood/app/models/userModel"
)

type UserAction struct {
	ID       uint                `json:"id" gorm:"primaryKey"`
	UserId   uint                `json:"_"`
	User     *userModel.User     `json:"user"`
	ActionId uint                `json:"_"`
	Action   *actionModel.Action `json:"action"`
	History  []*logModel.Log     `json:"history" gorm:"foreignKey:ActionId"`
	Deleted  bool                `json:"deleted" gorm:"default:0"`
}

type MyActionsResponse []struct {
	Action actionModel.Action `json:"action"`
}
