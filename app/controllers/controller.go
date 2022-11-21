package controllers

import (
	"Pood/app/controllers/actionController"
	"Pood/app/controllers/logController"
	"Pood/app/controllers/userActionController"
	"Pood/app/controllers/userController"
)

type Controller struct {
	UserActionController *userActionController.UserActionController
	ActionController     *actionController.ActionController
	UserController       *userController.UserController
	LogController        *logController.LogController
}
