package activeActionsService

import (
	"gorm.io/gorm"
	"pood/v2/app/models/userActionModel"
	"pood/v2/app/models/userModel"
)

func GetActiveUserActions(db *gorm.DB, user userModel.User) (actions []userActionModel.UserAction, err error) {
	err = db.Preload("Action").
		Preload("Logs", "start_time IS NOT NULL and end_time IS NULL").
		Where(userActionModel.UserAction{UserId: user.ID}).
		Find(&actions).Error
	if err != nil {
		return actions, err
	}

	return actions, nil
}
