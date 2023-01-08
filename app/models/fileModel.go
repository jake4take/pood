package models

import (
	"time"
)

type File struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Path     string  `json:"path"`
	Name     string  `json:"name"`
	CreateAt string  `json:"create_at"`
	LogId    *uint   `json:"log_id"`
	DeleteAt *string `json:"delete_at"`
}

func (f *File) BeforeCreate() (err error) {
	f.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	return
}

type CreateFileRequest struct {
	File string `json:"file"`
}

type CreateFileResponse struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Path     string `json:"path"`
	Name     string `json:"name"`
	CreateAt string `json:"create_at"`
}
