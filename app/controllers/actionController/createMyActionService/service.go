package createMyActionService

import (
	"errors"
	"pood/v2/app/models/actionModel"
	"pood/v2/app/models/userActionModel"
	"pood/v2/app/models/userModel"
	"pood/v2/config"
)

func FindActionByName(action actionModel.Action) (*actionModel.Action, error) {
	var resp actionModel.Action

	err := config.Db.
		Where(actionModel.Action{Name: action.Name}).
		First(&resp).
		Error

	if err != nil {
		return nil, errors.New("action not found")
	}

	return &resp, nil
}

func CreateAction(action actionModel.Action) (resp *actionModel.Action, err error) {
	err = config.Db.FirstOrCreate(&resp, &action).Error
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func CreateUserAction(user userModel.User, action actionModel.Action) (*userActionModel.UserAction, error) {
	var userAction userActionModel.UserAction
	err := config.Db.
		FirstOrCreate(&userAction, &userActionModel.UserAction{
			UserId:   user.ID,
			ActionId: action.ID,
		}).
		Error

	if err != nil {
		return nil, errors.New("userAction create error")
	}

	return &userAction, nil
}

func EnableUserActionAgain(userAction userActionModel.UserAction) error {
	err := config.Db.
		Model(userActionModel.UserAction{}).
		Where(userAction).
		Update("deleted", false).
		Error

	if err != nil {
		return err
	}

	return nil
}
