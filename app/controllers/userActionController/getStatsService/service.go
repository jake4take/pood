package getStatsService

import (
	"gorm.io/gorm/clause"
	"pood/v2/app/models"
	"pood/v2/config"
)

func GetUserAction(action models.UserAction) (resp *models.UserAction, err error) {
	err = config.Db.
		Preload(clause.Associations).
		First(&resp, action).Error
	if err != nil {
		return nil, err
	}

	return resp, nil
}
