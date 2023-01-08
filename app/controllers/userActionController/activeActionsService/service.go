package activeActionsService

import (
	"gorm.io/gorm"
	"pood/v2/app/models"
)

func GetActiveUserActions(db *gorm.DB, user models.User) (actions []models.UserAction, err error) {
	err = db.Preload("Action").
		Preload("Logs", "start_time IS NOT NULL and end_time IS NULL").
		Where(models.UserAction{UserId: user.ID}).
		Find(&actions).Error
	if err != nil {
		return actions, err
	}

	return actions, nil
}
