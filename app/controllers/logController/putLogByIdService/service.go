package putLogByIdService

import (
	"fmt"
	"pood/v2/app/models"
	"pood/v2/config"
)

func UpdateFilesLogId(logId uint, fileId uint) {
	err := config.Db.Model(models.File{}).
		Where(models.File{ID: fileId}).
		Update("log_id", logId).
		Error

	if err != nil {
		fmt.Println(err.Error())
	}
}
