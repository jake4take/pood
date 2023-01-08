package models

import (
	"time"
)

type Log struct {
	Id           uint       `json:"id" gorm:"primaryKey"`
	UserActionId uint       `json:"user_action_id" `
	UserAction   UserAction `json:"user_action"`
	UserId       uint       `json:"user_id"`
	LogDate      time.Time  `json:"log_date"`
	Value        *string    `json:"value"`
	StartTime    *time.Time `json:"start_time"`
	EndTime      *time.Time `json:"end_time"`
	Count        *float64   `json:"count"`
	Description  *string    `json:"description"`
	Files        []File     `json:"files"`
}

type CreateLogRequest struct {
	Value       *string    `json:"value"`
	StartTime   *time.Time `json:"start_start"`
	EndTime     *time.Time `json:"end_time"`
	Count       *float64   `json:"count"`
	Description *string    `json:"description"`
	FileIds     *[]int64   `json:"file_ids"`
}

type GetStatsResponse struct {
	Action       Action `json:"action"`
	UserActionId uint   `json:"user_action_id"`
	Count        int    `json:"count"`
	Stats        []Log  `json:"stats"`
}

type GetLogsResponse []struct {
	Id           uint       `json:"id"`
	UserActionId uint       `json:"user_action_id"`
	LogDate      time.Time  `json:"log_date"`
	Value        *string    `json:"value"`
	StartTime    *time.Time `json:"start_time"`
	EndTime      *time.Time `json:"end_time"`
	Count        *float64   `json:"count"`
	Description  *string    `json:"description"`
	Files        []File     `json:"files"`
}

type PutLogRequest struct {
	Value       *string    `json:"value"`
	StartTime   *time.Time `json:"start_time"`
	EndTime     *time.Time `json:"end_time"`
	Count       *float64   `json:"count"`
	Description *string    `json:"description"`
	FileIds     *[]int64   `json:"file_ids"`
}
