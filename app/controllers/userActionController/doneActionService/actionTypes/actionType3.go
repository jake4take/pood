package actionTypes

import (
	"errors"
	"pood/v2/app/models/logModel"
)

func CreateLogType3(log logModel.Log) (string, error) {
	if log.Count == nil {
		return "", errors.New("count is required parameter (type float)")
	}

	err := CreateLog(log)
	if err != nil {
		return "", err
	}

	return "counted", nil
}
