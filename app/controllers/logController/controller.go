package logController

import (
	"Pood/app/controllers/actionController"
	"Pood/app/controllers/queryController"
	"Pood/app/controllers/tokenController"
	"Pood/app/controllers/userActionController"
	"Pood/app/models/actionModel"
	"Pood/app/models/logModel"
	"Pood/app/models/userActionModel"
	"Pood/config"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

type LogController struct{}

func NewLogController() *LogController {
	return &LogController{}
}

// Done
// @Summary Сделал action
// @Description Создать лог совершенного action по токену
// @Accept  json
// @Produce json
// @Tags    Actions
// @Param body body logModel.CreateLogRequest true "body"
// @Success 200 {object} swagModel.AccessCreateResponse
// @Router  /action/done [post]
// @Security ApiKeyAuth
func (LogController) Done(c *fiber.Ctx) error {
	user, err := tokenController.CheckToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	var request logModel.Log
	err = json.Unmarshal(c.Body(), &request)
	request.UserId = user.ID

	location := time.FixedZone("UTC-0", 0)
	request.LogDate = time.Now().In(location)

	err = userActionController.HaveUserAction(userActionModel.UserAction{UserId: user.ID, ActionId: request.ActionId})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "user not have action",
		})
	}

	if err = CreateLog(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"detail": "ok",
	})
}

// GetStats
// @Summary Получить статистику
// @Description Получить статистику по id action по токену
// @Accept  json
// @Produce json
// @Tags    Actions
// @Param id path string true "id"
// @Param filter[log_date][gte] query string false "date"
// @Param filter[log_date][lte] query string false "date"
// @Param order query string false "field[eq]"
// @Success 200 {object} []logModel.Log
// @Router  /action/{id}/stats [get]
// @Security ApiKeyAuth
func (LogController) GetStats(c *fiber.Ctx) error {
	user, err := tokenController.CheckToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	actionId, err := strconv.Atoi(c.Params("id", "0"))
	if err != nil || actionId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "action not found",
		})
	}

	queries := queryController.GetQueries(c)
	db := queryController.ConfigurationDbQuery(config.Db, queries)

	var response []logModel.Log
	db.Where(logModel.Log{ActionId: uint(actionId), UserId: user.ID}).
		Find(&response)

	action := actionController.GetAction(actionModel.Action{ID: uint(actionId)})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"action": action,
		"count":  len(response),
		"stats":  response,
	})
}

func CreateLog(log logModel.Log) error {
	err := config.Db.Create(&log).Error
	if err != nil {
		return err
	}

	return nil
}
