package userController

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"pood/v2/app/controllers/userController/getUserActionsByUserService"
	"pood/v2/app/models/userActionModel"
	"pood/v2/app/models/userModel"
	"pood/v2/app/services/tokenService"
	"strconv"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

// GetUserActionByUser
// @Summary Получить userActions юзера (private=false)
// @Description Получить userActions юзера (private=false) по id
// @Accept  json
// @Produce json
// @Tags    Users
// @Success 200 {array} userActionModel.UserActionsResponse
// @Failure 400 {object} defaultModel.FailedResponse
// @Failure 401 {object} defaultModel.FailedResponse
// @Param id path string true "id"
// @Router  /user/{id}/actions [get]
// @Security ApiKeyAuth
func (UserController) GetUserActionByUser(c *fiber.Ctx) error {
	_, err := tokenService.CheckToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"detail": err.Error()})
	}

	reqId, err := strconv.Atoi(c.Params("id", "0"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}
	if reqId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": "parameter id is incorrect"})
	}

	friend, err := getUserActionsByUserService.FindUserById(userModel.User{ID: uint(reqId)})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	friendAction, err := getUserActionsByUserService.GetUserActions(*friend)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	data, err := json.Marshal(friendAction)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	var response []userActionModel.UserActionsResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": response})
}
