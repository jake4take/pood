package router

import (
	"Pood/app/controllers"
	"Pood/app/controllers/actionController"
	"Pood/app/controllers/logController"
	"Pood/app/controllers/userActionController"
	"Pood/app/controllers/userController"
	"github.com/gofiber/fiber/v2"
)

func apiRouter(f fiber.Router) {
	appController := initializationController()

	userAction := f.Group("/actions")
	{
		userAction.Get("/my", appController.UserActionController.GetMyActions)
	}

	action := f.Group("/action")
	{
		action.Post("", appController.UserActionController.CreateMyAction)
		action.Post("/done", appController.LogController.Done)
		action.Get("/:id/stats", appController.LogController.GetStats)
		action.Delete("/:id", appController.UserActionController.DeleteAction)
	}

}

func initializationController() controllers.Controller {
	return controllers.Controller{
		ActionController:     actionController.NewActionController(),
		UserActionController: userActionController.NewUserActionController(),
		UserController:       userController.NewUserController(),
		LogController:        logController.NewLogController(),
	}
}
