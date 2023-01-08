package getMyActionService

import (
	"gorm.io/gorm"
	"pood/v2/app/models"
)

func GetUserActions(db *gorm.DB, user models.User) (actions []models.UserAction, err error) {
	err = db.Preload("Action").
		Preload("Action.UnitInfo").
		//Preload("Logs", func(db *gorm.DB) *gorm.DB {
		//	return db.Order("id DESC")
		//}).
		Where(models.UserAction{UserId: user.ID}).
		Find(&actions).Error
	if err != nil {
		return actions, err
	}

	return actions, nil
}
