package deletedActionService

import (
	"pood/v2/app/models/userActionModel"
	"pood/v2/config"
)

func DeleteUserAction(userAction userActionModel.UserAction) error {
	err := config.Db.
		Model(userActionModel.UserAction{}).
		Where(userAction).
		Update("deleted", true).
		Error

	if err != nil {
		return err
	}

	return nil
}
