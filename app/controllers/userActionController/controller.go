package userActionController

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"pood/v2/app/controllers/userActionController/activeActionsService"
	"pood/v2/app/controllers/userActionController/deletedActionService"
	"pood/v2/app/controllers/userActionController/doneActionService"
	"pood/v2/app/controllers/userActionController/getMyActionService"
	"pood/v2/app/controllers/userActionController/getStatsService"
	"pood/v2/app/controllers/userActionController/updateUserActionService"
	"pood/v2/app/models"
	"pood/v2/app/services/queryService"
	"pood/v2/app/services/tokenService"
	"pood/v2/config"
	"strconv"
)

type UserActionController struct{}

func NewUserActionController() *UserActionController {
	return &UserActionController{}
}

// GetMyUserActions
// @Summary Получить мои actions
// @Description Получить мои actions по токену
// @Accept  json
// @Produce json
// @Tags    UserActions
// @Success 200 {array} models.MyActionsResponse
// @Failure 400 {object} models.FailedResponse
// @Failure 401 {object} models.FailedResponse
// @Param deleted query boolean false "deleted"
// @Param order query string false "field[eq]"
// @Router  /userActions/my [get]
// @Security ApiKeyAuth
func (UserActionController) GetMyUserActions(c *fiber.Ctx) error {
	user, err := tokenService.CheckToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"detail": err.Error()})
	}
	queries := queryService.GetQueries(c)
	db := queryService.ConfigurationDbQuery(config.Db, queries)

	userActions, err := getMyActionService.GetUserActions(db, *user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	data, err := json.Marshal(userActions)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	var response []models.MyActionsResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": response})
}

// DeleteAction
// @Summary Удалить userAction
// @Description Удалить userAction по id по токену
// @Accept  json
// @Produce json
// @Tags    UserActions
// @Success 204 {object} models.SuccessResponse
// @Failure 400 {object} models.FailedResponse
// @Failure 401 {object} models.FailedResponse
// @Param id path string true "id"
// @Router  /userAction/{id} [delete]
// @Security ApiKeyAuth
func (UserActionController) DeleteAction(c *fiber.Ctx) error {
	_, err := tokenService.CheckToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	userActionId, err := strconv.Atoi(c.Params("id", "0"))
	if err != nil || userActionId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "action not found",
		})
	}

	err = deletedActionService.DeleteUserAction(models.UserAction{ID: uint(userActionId)})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "action not delete",
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"detail": "ok",
	})
}

// Done
// @Summary Сделал action
// Description Not required field - description *string*, files *[]int*
// @Description **If action.type = 3**, field count *float* is required;
// @Accept  json
// @Produce json
// @Tags    UserActions
// @Param id path string true "id"
// @Param body body models.CreateLogRequest true "body"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.FailedResponse
// @Failure 401 {object} models.FailedResponse
// @Router  /userAction/{id}/done [post]
// @Security ApiKeyAuth
func (UserActionController) Done(c *fiber.Ctx) error {
	user, err := tokenService.CheckToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	id, err := strconv.Atoi(c.Params("id", "0"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	if id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": "id is required"})
	}

	var request models.CreateLogRequest
	var reqFiles *[]int64
	err = json.Unmarshal(c.Body(), &request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	if request.FileIds != nil && len(*request.FileIds) > 0 {
		reqFiles = request.FileIds
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	var actionLog models.Log
	err = json.Unmarshal(requestData, &actionLog)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	actionLog.UserId = user.ID
	actionLog.UserActionId = uint(id)

	userAction, err := doneActionService.HaveUserAction(models.UserAction{ID: uint(id)})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "user not have action",
		})
	}

	response, err := doneActionService.CheckLogType(*userAction, actionLog, reqFiles)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"detail": response,
	})
}

// GetStats
// @Summary Получить статистику
// @Description Получить статистику по id action по токену
// @Accept  json
// @Produce json
// @Tags    UserActions
// @Param id path string true "id"
// @Param filter[log_date][gte] query string false "date"
// @Param filter[log_date][lte] query string false "date"
// @Param order query string false "field[eq]"
// @Success 200 {array} models.GetStatsResponse
// @Failure 400 {object} models.FailedResponse
// @Failure 401 {object} models.FailedResponse
// @Router  /userAction/{id}/stats [get]
// @Security ApiKeyAuth
func (UserActionController) GetStats(c *fiber.Ctx) error {
	user, err := tokenService.CheckToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	userActionId, err := strconv.Atoi(c.Params("id", "0"))
	if err != nil || userActionId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "user_action not found",
		})
	}

	queries := queryService.GetQueries(c)
	db := queryService.ConfigurationDbQuery(config.Db, queries)

	var response models.GetStatsResponse
	db.Preload("Files", "delete_at IS NULL").
		Where(models.Log{UserActionId: uint(userActionId), UserId: user.ID}).
		Find(&response.Stats)

	userAction, err := getStatsService.GetUserAction(models.UserAction{ID: uint(userActionId), UserId: user.ID})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "user not have this user_action",
		})
	}
	response.UserActionId = userAction.ID
	response.Action = *userAction.Action
	response.Count = len(response.Stats)

	return c.Status(fiber.StatusOK).JSON(response)
}

// ActiveUserActions
// @Summary Получить мои активные actions с типом интервал
// @Description Получить мои активные actions с типом интервал по токену
// @Accept  json
// @Produce json
// @Tags    UserActions
// @Success 200 {object} models.MyActiveActions
// @Failure 400 {object} models.FailedResponse
// @Failure 401 {object} models.FailedResponse
// @Param deleted query boolean false "deleted"
// @Param order query string false "field[eq]"
// @Router  /userActions/my/active [get]
// @Security ApiKeyAuth
func (UserActionController) ActiveUserActions(c *fiber.Ctx) error {
	user, err := tokenService.CheckToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"detail": err.Error()})
	}
	queries := queryService.GetQueries(c)
	db := queryService.ConfigurationDbQuery(config.Db, queries)

	userActions, err := activeActionsService.GetActiveUserActions(db, *user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	var activeAction []models.MyActiveActions
	for _, item := range userActions {
		if len(item.Logs) != 0 {
			newItem := models.MyActiveActions{
				ID:          item.ID,
				Action:      *item.Action,
				StartTime:   *item.Logs[0].StartTime,
				Description: item.Logs[0].Description,
			}
			activeAction = append(activeAction, newItem)
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": activeAction})
}

// UpdatePrivateUserAction
// @Summary Редактировать userAction
// @Description Редактировать userAction по id по токену
// @Accept  json
// @Produce json
// @Tags    UserActions
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.FailedResponse
// @Failure 401 {object} models.FailedResponse
// @Param id path string true "id"
// @Param body body models.UpdateRequest true "body"
// @Router  /userAction/{id}/private [put]
// @Security ApiKeyAuth
func (UserActionController) UpdatePrivateUserAction(c *fiber.Ctx) error {
	user, err := tokenService.CheckToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"detail": err.Error()})
	}

	userActionId, err := strconv.Atoi(c.Params("id", "0"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}
	if userActionId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": "id is required"})
	}

	userAction, err := updateUserActionService.GetUserActionById(models.UserAction{ID: uint(userActionId), UserId: user.ID})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": "user has not this user_action"})
	}

	var request models.UpdateRequest
	err = json.Unmarshal(c.Body(), &request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	if request.Private == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": "private is required body parameter"})
	}

	userAction.Private = *request.Private

	err = updateUserActionService.UpdateUserActionPrivateStatus(userAction)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "private status updated"})
}
