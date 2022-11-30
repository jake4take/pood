package actionTypes

import (
	"pood/v2/app/models/logModel"
	"pood/v2/app/models/userActionModel"
	"pood/v2/config"
)

func CreateLogType2(userAction userActionModel.UserAction, log logModel.Log) (string, error) {
	var logs []logModel.Log
	err := config.Db.
		Where("user_action_id = ?", userAction.ID).
		Where("start_time IS NOT NULL").
		Where("end_time IS NULL").
		Find(&logs).
		Error

	if err != nil {
		return "1", err
	}

	if len(logs) == 0 {
		log.StartTime = &log.LogDate
		err = CreateLog(log)
		return "new interval started", nil
	}

	for _, item := range logs {
		item.EndTime = &log.LogDate
		err = FinishLog(item)
		if err != nil {
			return "", err
		}
	}

	return "interval ended", nil
}

func FinishLog(log logModel.Log) error {
	err := config.Db.
		Model(logModel.Log{}).
		Where(logModel.Log{Id: log.Id}).
		Updates(&log).
		Error

	return err
}
