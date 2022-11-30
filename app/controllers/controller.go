package controllers

import (
	"pood/v2/app/controllers/actionController"
	"pood/v2/app/controllers/logController"
	"pood/v2/app/controllers/typeInfoController"
	"pood/v2/app/controllers/unitController"
	"pood/v2/app/controllers/userActionController"
	"pood/v2/app/controllers/userController"
)

type Controller struct {
	UserActionController *userActionController.UserActionController
	ActionController     *actionController.ActionController
	UserController       *userController.UserController
	LogController        *logController.LogController
	TypeInfoController   *typeInfoController.TypeInfoController
	UnitController       *unitController.UnitController
}
