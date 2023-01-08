package actionTypes

import (
	"pood/v2/app/models"
	"pood/v2/config"
)

func CreateLog(log models.Log) (*models.Log, error) {
	err := config.Db.Create(&log).Error
	if err != nil {
		return nil, err
	}

	return &log, nil
}
