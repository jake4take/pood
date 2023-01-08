package getMyLogsService

import (
	"gorm.io/gorm"
	"pood/v2/app/models"
)

func GetMyLogs(db *gorm.DB, user models.User) (resp []models.Log, err error) {
	err = db.
		Model(models.Log{}).
		Preload("Files", "delete_at IS NULL").
		Where(models.Log{UserId: user.ID}).
		Find(&resp).
		Error

	if err != nil {
		return resp, err
	}

	return resp, nil
}
