package router

import (
	"github.com/gofiber/fiber/v2"
	"pood/v2/app/controllers"
	"pood/v2/app/controllers/actionController"
	"pood/v2/app/controllers/fileController"
	"pood/v2/app/controllers/logController"
	"pood/v2/app/controllers/typeInfoController"
	"pood/v2/app/controllers/unitController"
	"pood/v2/app/controllers/userActionController"
	"pood/v2/app/controllers/userController"
)

func apiRouter(f fiber.Router) {
	appController := initializationController()

	action := f.Group("/action")
	actions := f.Group("/actions")
	{
		action.Post("", appController.ActionController.CreateMyAction)
		actions.Get("", appController.ActionController.FindActionByName)
	}

	userActions := f.Group("/userActions")
	{
		userActions.Get("/my", appController.UserActionController.GetMyUserActions)
		userActions.Get("/my/active", appController.UserActionController.ActiveUserActions)
	}

	userAction := f.Group("/userAction")
	{
		userAction.Delete("/:id", appController.UserActionController.DeleteAction)
		userAction.Post("/:id/done", appController.UserActionController.Done)
		userAction.Get("/:id/stats", appController.UserActionController.GetStats)
		userAction.Put("/:id/private", appController.UserActionController.UpdatePrivateUserAction)
	}

	typeInfo := f.Group("/typeInfo")
	units := f.Group("/unitInfo")
	{
		typeInfo.Get("/", appController.TypeInfoController.GetTypeInfo)
		units.Get("/", appController.UnitController.GetUnits)
	}

	user := f.Group("/user")
	{
		user.Get("/:id/actions", appController.UserController.GetUserActionByUser)
	}

	logs := f.Group("/logs")
	log := f.Group("/log")
	{
		logs.Get("/my", appController.LogController.GetMyLogs)
		log.Put("/:id", appController.LogController.PutLogById)
	}

	file := f.Group("/file")
	{
		file.Post("/", appController.FileController.CreateNewFile)
		file.Delete("/:id", appController.FileController.DeleteFile)
	}

	upload := f.Group("/uploads")
	{
		upload.Get("/:date/:name", appController.FileController.GetFile)
	}

}

func initializationController() controllers.Controller {
	return controllers.Controller{
		ActionController:     actionController.NewActionController(),
		UserActionController: userActionController.NewUserActionController(),
		UserController:       userController.NewUserController(),
		LogController:        logController.NewLogController(),
		TypeInfoController:   typeInfoController.NewTypeInfoController(),
		UnitController:       unitController.NewUnitController(),
		FileController:       fileController.NewActionController(),
	}
}
