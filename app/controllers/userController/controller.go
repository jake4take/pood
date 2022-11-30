package userController

import (
	"errors"
	"fmt"
	"pood/v2/app/models/tokenModel"
	"pood/v2/app/models/userModel"
	"pood/v2/config"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func GetUserByToken(token tokenModel.Token) (*userModel.User, error) {
	var user userModel.User
	err := config.Db.
		Where(userModel.User{ID: token.UserId}).
		First(&user).
		Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("user not found (id = %d)", token.UserId))
	}

	return &user, nil
}
