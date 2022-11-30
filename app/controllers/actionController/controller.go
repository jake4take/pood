package actionController

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"pood/v2/app/controllers/actionController/createMyActionService"
	"pood/v2/app/models/actionModel"
	"pood/v2/app/services/tokenService"
)

type ActionController struct{}

func NewActionController() *ActionController {
	return &ActionController{}
}

var defaultActionType uint = 3

// CreateMyAction
// @Summary Создать новый action
// @Description Создать новый action и привязаться к нему по токену
// @Accept  json
// @Produce json
// @Tags    Actions
// @Param body body actionModel.ActionCreateRequest true "body"
// @Success 201 {object} defaultModel.SuccessResponse
// @Failure 401 {object} defaultModel.FailedResponse
// @Router  /action [post]
// @Security ApiKeyAuth
func (ActionController) CreateMyAction(c *fiber.Ctx) error {
	user, err := tokenService.CheckToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	var reqAction *actionModel.Action
	err = json.Unmarshal(c.Body(), &reqAction)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	if reqAction.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": errors.New("incorrect request"),
		})
	}

	if reqAction.Unit == 0 {
		reqAction.Unit = defaultActionType
	}

	action, err := createMyActionService.FindActionByName(*reqAction)
	if err != nil {
		action, err = createMyActionService.CreateAction(*reqAction)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"detail": err.Error(),
			})
		}
	}

	userAction, err := createMyActionService.CreateUserAction(*user, *action)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	if userAction.Deleted {
		err = createMyActionService.EnableUserActionAgain(*userAction)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"detail": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"detail": "action created",
	})
}
