package actionTypes

import (
	"pood/v2/app/models/logModel"
	"pood/v2/config"
)

func CreateLog(log logModel.Log) error {
	err := config.Db.Create(&log).Error
	if err != nil {
		return err
	}

	return nil
}
