package authorizationsService

import (
	"errors"
	"fmt"
	"pood/v2/app/models"
	"pood/v2/config"
)

func GetUserByToken(token models.Token) (*models.User, error) {
	var user models.User
	err := config.Db.
		Where(models.User{ID: token.UserId}).
		First(&user).
		Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("user not found (id = %d)", token.UserId))
	}

	return &user, nil
}
