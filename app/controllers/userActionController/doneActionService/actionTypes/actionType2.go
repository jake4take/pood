package actionTypes

import (
	"pood/v2/app/models"
	"pood/v2/config"
)

func CreateLogType2(userAction models.UserAction, log models.Log) (*models.Log, string, error) {
	var logs []models.Log
	err := config.Db.
		Where("user_action_id = ?", userAction.ID).
		Where("start_time IS NOT NULL").
		Where("end_time IS NULL").
		Find(&logs).
		Error

	if err != nil {
		return nil, "1", err
	}

	if len(logs) == 0 {
		log.StartTime = &log.LogDate
		respLog, _ := CreateLog(log)
		return respLog, "new interval started", nil
	}

	for _, item := range logs {
		item.EndTime = &log.LogDate
		err = FinishLog(item)
		if err != nil {
			return nil, "", err
		}
	}

	return &logs[0], "interval ended", nil
}

func FinishLog(log models.Log) error {
	err := config.Db.
		Model(models.Log{}).
		Where(models.Log{Id: log.Id}).
		Updates(&log).
		Error

	return err
}
