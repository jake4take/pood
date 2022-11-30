package getStatsService

import (
	"gorm.io/gorm/clause"
	"pood/v2/app/models/userActionModel"
	"pood/v2/config"
)

func GetUserAction(action userActionModel.UserAction) (resp *userActionModel.UserAction) {
	err := config.Db.
		Preload(clause.Associations).
		First(&resp, action).Error
	if err != nil {
		return nil
	}

	return resp
}
