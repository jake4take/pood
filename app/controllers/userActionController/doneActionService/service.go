package doneActionService

import (
	"pood/v2/app/controllers/userActionController/doneActionService/actionTypes"
	"pood/v2/app/models/logModel"
	"pood/v2/app/models/userActionModel"
	"pood/v2/config"
	"time"
)

func HaveUserAction(userAction userActionModel.UserAction) (*userActionModel.UserAction, error) {
	var resp userActionModel.UserAction
	userAction.Deleted = false

	err := config.Db.
		Preload("Action").
		Preload("Logs").
		Where(&userAction).
		First(&resp).
		Error
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func CheckLogType(userAction userActionModel.UserAction, log logModel.Log) (detail string, err error) {
	location := time.FixedZone("UTC-0", 0)
	log.LogDate = time.Now().In(location)
	detail = "ok"

	switch userAction.Action.Type {
	case 1:
		err = actionTypes.CreateLog(log)
		detail = "counter increased"
	case 2:
		detail, err = actionTypes.CreateLogType2(userAction, log)
	case 3:
		detail, err = actionTypes.CreateLogType3(log)
	}

	return detail, err
}
