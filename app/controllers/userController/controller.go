package userController

import (
	"Pood/app/models/tokenModel"
	"Pood/app/models/userModel"
	"Pood/config"
	"errors"
	"fmt"
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
