package logModel

import (
	"pood/v2/app/models/actionModel"
	"time"
)

type Log struct {
	Id           uint       `json:"id" gorm:"primaryKey"`
	UserActionId uint       `json:"user_action_id" `
	UserId       uint       `json:"user_id"`
	LogDate      time.Time  `json:"log_date"`
	Value        *string    `json:"value"`
	StartTime    *time.Time `json:"start_time"`
	EndTime      *time.Time `json:"end_time"`
	Count        *float64   `json:"count"`
	Description  *string    `json:"description"`
}

type CreateLogRequest struct {
	Value       *string    `json:"value"`
	StartTime   *time.Time `json:"start_start"`
	EndTime     *time.Time `json:"end_time"`
	Count       *float64   `json:"count"`
	Description *string    `json:"description"`
}

type GetStatsResponse struct {
	Action       actionModel.Action `json:"action"`
	UserActionId uint               `json:"user_action_id"`
	Count        int                `json:"count"`
	Stats        []Log              `json:"stats"`
}
