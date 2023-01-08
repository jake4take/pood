package updateUserActionService

import (
	"pood/v2/app/models"
	"pood/v2/config"
)

func GetUserActionById(userAction models.UserAction) (models.UserAction, error) {
	err := config.Db.First(&userAction, userAction).Error
	return userAction, err
}

func UpdateUserActionPrivateStatus(userAction models.UserAction) error {
	err := config.Db.
		Model(models.UserAction{}).
		Where(models.UserAction{ID: userAction.ID}).
		Update("private", userAction.Private).
		Error

	if err != nil {
		return err
	}

	return nil
}
