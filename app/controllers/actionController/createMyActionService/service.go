package createMyActionService

import (
	"errors"
	"fmt"
	"pood/v2/app/models/actionModel"
	"pood/v2/app/models/userActionModel"
	"pood/v2/app/models/userModel"
	"pood/v2/config"
)

func CreateUserActionByActionId(user userModel.User, reqAction *actionModel.Action) (error, string) {
	action, err := FindActionById(*reqAction)
	if err != nil {
		return err, ""
	}

	userAction := userActionModel.UserAction{UserId: user.ID, ActionId: action.ID}
	err = config.Db.First(&userAction, userAction).Error

	if err != nil {
		_, err = CreateUserAction(userModel.User{ID: user.ID}, actionModel.Action{ID: action.ID})
		if err != nil {
			return err, ""
		}
		return nil, "action added"
	}

	if userAction.Deleted {
		err = EnableUserActionAgain(userAction)
		if err != nil {
			return err, ""
		}
		return errors.New("ok"), "action restored"
	}
	return errors.New("user has this action"), ""
}

func CreateUserActionByActionName(user userModel.User, reqAction *actionModel.Action) (error, string) {
	action, err := FindActionByName(*reqAction)
	if err != nil {
		action, err = CreateAction(*reqAction)
		if err != nil {
			return err, ""
		}
	}

	var userAction *userActionModel.UserAction
	err = config.Db.First(&userAction, userActionModel.UserAction{UserId: user.ID, ActionId: action.ID}).Error
	if err != nil {
		userAction, err = CreateUserAction(user, *action)
		if err != nil {
			return err, ""
		}
		return nil, "user_action added"
	}

	if userAction.Deleted {
		err = EnableUserActionAgain(*userAction)
		if err != nil {
			return err, ""
		}
	}

	return errors.New("user has this action"), ""
}

func FindActionByName(action actionModel.Action) (*actionModel.Action, error) {
	var resp actionModel.Action

	err := config.Db.
		First(&resp, action).
		Error

	if err != nil {
		return nil, errors.New("action not found")
	}

	return &resp, nil
}

func FindActionById(action actionModel.Action) (*actionModel.Action, error) {
	var resp actionModel.Action

	err := config.Db.First(&resp, actionModel.Action{ID: action.ID}).Error

	if err != nil {
		return nil, errors.New(fmt.Sprintf("action not found (%d)", action.ID))
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
		return nil, errors.New("user_action create error")
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
