package unitController

import (
	"github.com/gofiber/fiber/v2"
	"pood/v2/app/models/unitModel"
	"pood/v2/app/services/tokenService"
	"pood/v2/config"
)

type UnitController struct{}

func NewUnitController() *UnitController {
	return &UnitController{}
}

// GetUnits
// @Summary Единицы измерения
// @Accept  json
// @Produce json
// @Tags    TypeInfo
// @Success 200 {object} unitModel.Unit
// @Failure 401 {object} defaultModel.FailedResponse
// @Router  /unitInfo [get]
// @Security ApiKeyAuth
func (UnitController) GetUnits(c *fiber.Ctx) error {
	_, err := tokenService.CheckToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	var unit []unitModel.Unit
	err = config.Db.
		Model(unitModel.Unit{}).
		Find(&unit).
		Error

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": unit,
	})
}
