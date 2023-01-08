package tokenService

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"pood/v2/app/controllers/userActionController/authorizationsService"
	"pood/v2/app/models"
	"pood/v2/config"
)

func CheckToken(c *fiber.Ctx) (*models.User, error) {
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

	user, err := authorizationsService.GetUserByToken(*token)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetToken(token string) (*models.Token, error) {
	var resp models.Token
	err := config.Db.
		Where(models.Token{Token: token}).
		First(&resp).
		Error

	if err != nil {
		return nil, errors.New("token not found")
	}

	return &resp, nil
}
