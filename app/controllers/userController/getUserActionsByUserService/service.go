package getUserActionsByUserService

import (
	"pood/v2/app/models"
	"pood/v2/config"
)

func FindUserById(user models.User) (*models.User, error) {
	err := config.Db.Find(&user, user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserActions(user models.User) (actions []models.UserAction, err error) {
	err = config.Db.Preload("Action").
		Preload("Action.UnitInfo").
		Where(models.UserAction{UserId: user.ID}).
		Where("deleted = false").
		Where("private = false").
		Find(&actions).Error
	if err != nil {
		return actions, err
	}

	return actions, nil
}
