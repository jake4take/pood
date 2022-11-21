package logModel

import "time"

type Log struct {
	Id       uint      `json:"id" gorm:"primaryKey"`
	ActionId uint      `json:"action_id"`
	UserId   uint      `json:"user_id"`
	LogDate  time.Time `json:"log_date"`
}

type CreateLogRequest struct {
	ActionId uint `json:"action_id"`
}
