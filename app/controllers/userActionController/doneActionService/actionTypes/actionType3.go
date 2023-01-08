package actionTypes

import (
	"errors"
	"pood/v2/app/models"
)

func CreateLogType3(log models.Log) (*models.Log, string, error) {
	if log.Count == nil {
		return nil, "", errors.New("count is required parameter (type float)")
	}

	resp, err := CreateLog(log)
	if err != nil {
		return nil, "", err
	}

	return resp, "counted", nil
}
