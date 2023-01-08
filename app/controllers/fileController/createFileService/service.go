package createFileService

import (
	"fmt"
	"mime/multipart"
	"pood/v2/app/models"
	"pood/v2/config"
)

func CreateFile(req *multipart.FileHeader, path string) (resp *models.File, err error) {
	var file models.File
	file.Name = req.Filename
	file.Path = path

	err = config.Db.FirstOrCreate(&resp, &file).Error
	fmt.Println(resp.ID)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
