package getUserActionsByUserService

import (
	"pood/v2/app/models/userActionModel"
	"pood/v2/app/models/userModel"
	"pood/v2/config"
)

func FindUserById(user userModel.User) (*userModel.User, error) {
	err := config.Db.Find(&user, user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserActions(user userModel.User) (actions []userActionModel.UserAction, err error) {
	err = config.Db.Preload("Action").
		Preload("Action.UnitInfo").
		Where(userActionModel.UserAction{UserId: user.ID}).
		Where("deleted = false").
		Where("private = false").
		Find(&actions).Error
	if err != nil {
		return actions, err
	}

	return actions, nil
}
