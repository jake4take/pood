package actionController

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"pood/v2/app/controllers/actionController/createMyActionService"
	"pood/v2/app/models"
	"pood/v2/app/services/tokenService"
	"pood/v2/config"
)

type ActionController struct{}

func NewActionController() *ActionController {
	return &ActionController{}
}

// CreateMyAction
// @Summary Создать новый action
// @Description Создать новый action и привязаться к нему по токену
// @Description - привязка по **id** или создание по остальным полям
// @Accept  json
// @Produce json
// @Tags    Actions
// @Param body body models.ActionCreateRequest true "body"
// @Success 201 {object} models.SuccessResponse
// @Failure 401 {object} models.FailedResponse
// @Router  /action [post]
// @Security ApiKeyAuth
func (ActionController) CreateMyAction(c *fiber.Ctx) error {
	var detail string

	user, err := tokenService.CheckToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	var reqAction *models.Action
	err = json.Unmarshal(c.Body(), &reqAction)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	switch {
	case reqAction.ID != 0:
		err, detail = createMyActionService.CreateUserActionByActionId(*user, reqAction)
	case reqAction.Name != "":
		err, detail = createMyActionService.CreateUserActionByActionName(*user, reqAction)
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "name or id is required",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"detail": detail,
	})
}

// FindActionByName
// @Summary Найти action по неполному совпадению имени
// @Accept  json
// @Produce json
// @Tags    Actions
// @Param name query string false "name"
// @Param template query boolean false "template"
// @Success 200 {array} models.Action
// @Failure 401 {object} models.FailedResponse
// @Router  /actions [get]
// @Security ApiKeyAuth
func (ActionController) FindActionByName(c *fiber.Ctx) error {
	_, err := tokenService.CheckToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	db := config.Db
	if name := c.Query("name", ""); name != "" {
		db = db.Where(fmt.Sprintf("name like '%%%s%%'", name))
	}

	if template := c.Query("template", ""); template == "false" {
		db = db.Where("template = false")
	} else if template == "true" {
		db = db.Where("template = true")
	}

	var actions []models.Action
	err = db.Find(&actions).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": actions,
	})
}
