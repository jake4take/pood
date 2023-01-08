package doneActionService

import (
	"fmt"
	"pood/v2/app/controllers/userActionController/doneActionService/actionTypes"
	"pood/v2/app/models"
	"pood/v2/config"
	"time"
)

func HaveUserAction(userAction models.UserAction) (*models.UserAction, error) {
	var resp models.UserAction
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

func CheckLogType(userAction models.UserAction, log models.Log, filesId *[]int64) (detail string, err error) {
	location := time.FixedZone("UTC-0", 0)
	log.LogDate = time.Now().In(location)
	detail = "ok"

	var respLog *models.Log
	switch userAction.Action.Type {
	case 1:
		respLog, err = actionTypes.CreateLog(log)
		detail = "counter increased"
	case 2:
		respLog, detail, err = actionTypes.CreateLogType2(userAction, log)
	case 3:
		respLog, detail, err = actionTypes.CreateLogType3(log)
	}

	if respLog != nil && filesId != nil {
		for _, id := range *filesId {
			fmt.Println(respLog.Id)
			updateFilesLogId(respLog.Id, uint(id))
		}
	}

	return detail, err
}

func updateFilesLogId(logId uint, fileId uint) {
	err := config.Db.Model(models.File{}).
		Where(models.File{ID: fileId}).
		Update("log_id", logId).
		Error

	if err != nil {
		fmt.Println(err.Error())
	}
}
