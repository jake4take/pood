package getMyActionService

import (
	"gorm.io/gorm"
	"pood/v2/app/models/userActionModel"
	"pood/v2/app/models/userModel"
)

func GetUserActions(db *gorm.DB, user userModel.User) (actions []userActionModel.UserAction, err error) {
	err = db.Preload("Action").
		Preload("Action.UnitInfo").
		Preload("Logs").
		Where(userActionModel.UserAction{UserId: user.ID}).
		Find(&actions).Error
	if err != nil {
		return actions, err
	}

	return actions, nil
}
