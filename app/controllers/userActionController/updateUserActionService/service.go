package updateUserActionService

import (
	"pood/v2/app/models/userActionModel"
	"pood/v2/config"
)

func GetUserActionById(userAction userActionModel.UserAction) (userActionModel.UserAction, error) {
	err := config.Db.First(&userAction, userAction).Error
	return userAction, err
}

func UpdateUserActionPrivateStatus(userAction userActionModel.UserAction) error {
	err := config.Db.
		Model(userActionModel.UserAction{}).
		Where(userActionModel.UserAction{ID: userAction.ID}).
		Update("private", userAction.Private).
		Error

	if err != nil {
		return err
	}

	return nil
}
