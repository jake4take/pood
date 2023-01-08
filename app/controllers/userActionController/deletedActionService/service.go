package deletedActionService

import (
	"pood/v2/app/models"
	"pood/v2/config"
)

func DeleteUserAction(userAction models.UserAction) error {
	err := config.Db.
		Model(models.UserAction{}).
		Where(userAction).
		Update("deleted", true).
		Error

	if err != nil {
		return err
	}

	return nil
}
