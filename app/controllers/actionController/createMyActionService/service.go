package createMyActionService

import (
	"errors"
	"fmt"
	"pood/v2/app/models"
	"pood/v2/config"
)

func CreateUserActionByActionId(user models.User, reqAction *models.Action) (error, string) {
	action, err := FindActionById(*reqAction)
	if err != nil {
		return err, ""
	}

	userAction := models.UserAction{UserId: user.ID, ActionId: action.ID}
	err = config.Db.First(&userAction, userAction).Error

	if err != nil {
		_, err = CreateUserAction(models.User{ID: user.ID}, models.Action{ID: action.ID})
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

func CreateUserActionByActionName(user models.User, reqAction *models.Action) (error, string) {
	action, err := FindActionByName(*reqAction)
	if err != nil {
		action, err = CreateAction(*reqAction)
		if err != nil {
			return err, ""
		}
	}

	var userAction *models.UserAction
	err = config.Db.First(&userAction, models.UserAction{UserId: user.ID, ActionId: action.ID}).Error
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

func FindActionByName(action models.Action) (*models.Action, error) {
	var resp models.Action

	err := config.Db.
		First(&resp, action).
		Error

	if err != nil {
		return nil, errors.New("action not found")
	}

	return &resp, nil
}

func FindActionById(action models.Action) (*models.Action, error) {
	var resp models.Action

	err := config.Db.First(&resp, models.Action{ID: action.ID}).Error

	if err != nil {
		return nil, errors.New(fmt.Sprintf("action not found (%d)", action.ID))
	}

	return &resp, nil
}

func CreateAction(action models.Action) (resp *models.Action, err error) {
	err = config.Db.FirstOrCreate(&resp, &action).Error
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func CreateUserAction(user models.User, action models.Action) (*models.UserAction, error) {
	var userAction models.UserAction
	err := config.Db.
		FirstOrCreate(&userAction, &models.UserAction{
			UserId:   user.ID,
			ActionId: action.ID,
		}).
		Error

	if err != nil {
		return nil, errors.New("user_action create error")
	}

	return &userAction, nil
}

func EnableUserActionAgain(userAction models.UserAction) error {
	err := config.Db.
		Model(models.UserAction{}).
		Where(userAction).
		Update("deleted", false).
		Error

	if err != nil {
		return err
	}

	return nil
}
