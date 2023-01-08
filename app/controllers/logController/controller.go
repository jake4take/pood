package logController

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"pood/v2/app/controllers/logController/getMyLogsService"
	"pood/v2/app/controllers/logController/putLogByIdService"
	"pood/v2/app/models"
	"pood/v2/app/services/queryService"
	"pood/v2/app/services/tokenService"
	"pood/v2/config"
	"strconv"
)

type LogController struct{}

func NewLogController() *LogController {
	return &LogController{}
}

// GetMyLogs
// @Summary Получить мои действия
// @Description
// @Accept  json
// @Produce json
// @Tags    Logs
// @Success 200 {array} models.GetLogsResponse
// @Failure 400 {object} models.FailedResponse
// @Failure 401 {object} models.FailedResponse
// @Param order query string false "field[eq]"
// @Router  /logs/my [get]
// @Security ApiKeyAuth
func (LogController) GetMyLogs(c *fiber.Ctx) error {
	user, err := tokenService.CheckToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"detail": err.Error()})
	}
	queries := queryService.GetQueries(c)
	db := queryService.ConfigurationDbQuery(config.Db, queries)

	logs, err := getMyLogsService.GetMyLogs(db, *user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	data, err := json.Marshal(logs)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	var response models.GetLogsResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": response})
}

// PutLogById
// @Summary Изменить лог
// @Description
// @Accept  json
// @Produce json
// @Tags    Logs
// @Success 200 {array} models.GetLogsResponse
// @Failure 400 {object} models.FailedResponse
// @Failure 401 {object} models.FailedResponse
// @Param id path string true "id"
// @Param body body models.PutLogRequest true "body"
// @Router  /log/{id} [put]
// @Security ApiKeyAuth
func (LogController) PutLogById(c *fiber.Ctx) error {
	user, err := tokenService.CheckToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"detail": err.Error()})
	}

	id, err := strconv.Atoi(c.Params("id", "0"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}
	if id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": "id is required"})
	}

	var log models.Log
	err = config.Db.Where(models.Log{Id: uint(id), UserId: user.ID}).First(&log).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"detail": err.Error()})
	}

	var request models.PutLogRequest
	err = json.Unmarshal(c.Body(), &request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	log = models.Log{
		Id:          uint(id),
		Value:       request.Value,
		StartTime:   request.StartTime,
		EndTime:     request.EndTime,
		Count:       request.Count,
		Description: request.Description,
	}

	err = config.Db.Updates(&log).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	if request.FileIds != nil && len(*request.FileIds) > 0 {
		for _, fileId := range *request.FileIds {
			putLogByIdService.UpdateFilesLogId(uint(id), uint(fileId))
		}
	}

	err = config.Db.Preload("Files").Where(models.Log{Id: uint(id), UserId: user.ID}).First(&log).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"detail": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(log)
}
