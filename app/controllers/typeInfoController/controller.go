package typeInfoController

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"pood/v2/app/models"
	"pood/v2/app/services/tokenService"
	"pood/v2/config"
)

type TypeInfoController struct{}

func NewTypeInfoController() *TypeInfoController {
	return &TypeInfoController{}
}

// GetTypeInfo
// @Summary Получить информацию по типам
// @Description Получить описание типов и вложеннымх в него сабтипов
// @Accept  json
// @Produce json
// @Tags    TypeInfo
// @Success 200 {object} models.TypeInfoResponse
// @Failure 401 {object} models.FailedResponse
// @Router  /typeInfo [get]
// @Security ApiKeyAuth
func (TypeInfoController) GetTypeInfo(c *fiber.Ctx) error {
	_, err := tokenService.CheckToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	var typeInfo []models.TypeInfo
	err = config.Db.
		Model(models.TypeInfo{}).
		Preload("SubType").
		Find(&typeInfo).
		Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	var response []models.TypeInfoResponse
	data, err := json.Marshal(typeInfo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": response,
	})
}
