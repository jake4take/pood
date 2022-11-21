package tokenController

import (
	"Pood/app/controllers/userController"
	"Pood/app/models/tokenModel"
	"Pood/app/models/userModel"
	"Pood/config"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func CheckToken(c *fiber.Ctx) (*userModel.User, error) {
	fullToken := c.GetReqHeaders()["Authorization"]

	if fullToken == "" {
		return nil, errors.New("who are you? auth failed")
	}

	//tokenString := strings.Fields(fullToken)
	//if len(tokenString) != 2 {
	//	return nil, errors.New("invalid token")
	//}

	token, err := GetToken(fullToken)
	if err != nil {
		return nil, err
	}

	user, err := userController.GetUserByToken(*token)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetToken(token string) (*tokenModel.Token, error) {
	var resp tokenModel.Token
	err := config.Db.
		Where(tokenModel.Token{Token: token}).
		First(&resp).
		Error

	if err != nil {
		return nil, errors.New("token not found")
	}

	return &resp, nil
}
