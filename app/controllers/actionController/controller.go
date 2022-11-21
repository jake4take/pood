package actionController

import (
	"Pood/app/models/actionModel"
	"Pood/config"
)

type ActionController struct{}

func NewActionController() *ActionController {
	return &ActionController{}
}

func GetAction(action actionModel.Action) (resp *actionModel.Action) {
	err := config.Db.First(&resp, action).Error
	if err != nil {
		return nil
	}

	return resp
}
