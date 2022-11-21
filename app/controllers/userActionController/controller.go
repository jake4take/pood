package userActionController

import (
	"Pood/app/controllers/queryController"
	"Pood/app/controllers/tokenController"
	"Pood/app/models/actionModel"
	"Pood/app/models/userActionModel"
	"Pood/app/models/userModel"
	"Pood/config"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
)

type UserActionController struct{}

func NewUserActionController() *UserActionController {
	return &UserActionController{}
}

// GetMyActions
// @Summary Получить мои actions
// @Description Получить мои actions по токену
// @Accept  json
// @Produce json
// @Tags    Actions
// @Success 200 {object} userActionModel.MyActionsResponse
// @Param deleted query boolean false "deleted"
// @Router  /actions/my [get]
// @Security ApiKeyAuth
func (UserActionController) GetMyActions(c *fiber.Ctx) error {
	user, err := tokenController.CheckToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}
	queries := queryController.GetQueries(c)
	db := queryController.ConfigurationDbQuery(config.Db, queries)
	fmt.Println()

	userActions := GetUserActions(db, *user)

	data, err := json.Marshal(userActions)
	if err != nil {
		panic(err)
	}

	var response = userActionModel.MyActionsResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"actions": response,
	})
}

// CreateMyAction
// @Summary Создать новый action
// @Description Создать новый action и привязаться к нему по токену
// @Accept  json
// @Produce json
// @Tags    Actions
// @Param body body actionModel.ActionCreateRequest true "body"
// @Success 200 {object} swagModel.AccessCreateResponse
// @Router  /action [post]
// @Security ApiKeyAuth
func (UserActionController) CreateMyAction(c *fiber.Ctx) error {
	user, err := tokenController.CheckToken(c)
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

	action, err := FindActionByName(*reqAction)
	if err != nil {
		action = CreateAction(*reqAction)
	}

	_, err = CreateUserAction(*user, *action)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"detail": "action created",
	})
}

// DeleteAction
// @Summary Удалить userAction
// @Description Удалить userAction по id по токену
// @Accept  json
// @Produce json
// @Tags    Actions
// @Param id path string true "id"
// @Router  /action/{id} [delete]
// @Security ApiKeyAuth
func (UserActionController) DeleteAction(c *fiber.Ctx) error {
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

	err = DeleteUserAction(userActionModel.UserAction{UserId: user.ID, ActionId: uint(actionId)})
	if err != nil {
		fmt.Println(actionId, err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "action not delete",
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"detail": "ok",
	})
}

func GetUserActions(db *gorm.DB, user userModel.User) (actions []userActionModel.UserAction) {
	db.Preload("Action").
		Where(userActionModel.UserAction{UserId: user.ID}).
		Find(&actions)

	return actions
}

func FindActionByName(action actionModel.Action) (*actionModel.Action, error) {
	var resp actionModel.Action

	err := config.Db.
		Where(actionModel.Action{Name: action.Name}).
		First(&resp).
		Error

	if err != nil {
		return nil, errors.New("action not found")
	}

	return &resp, nil
}

func CreateAction(action actionModel.Action) (resp *actionModel.Action) {
	err := config.Db.FirstOrCreate(&resp, &action).Error
	if err != nil {
		panic(err)
	}

	return resp
}

func CreateUserAction(user userModel.User, action actionModel.Action) (*userActionModel.UserAction, error) {
	var userAction userActionModel.UserAction
	err := config.Db.
		FirstOrCreate(&userAction, &userActionModel.UserAction{
			UserId:   user.ID,
			ActionId: action.ID,
		}).
		Error

	if err != nil {
		return nil, errors.New("userAction create error")
	}

	return &userAction, nil
}

func DeleteUserAction(userAction userActionModel.UserAction) error {
	err := config.Db.
		Model(userActionModel.UserAction{}).
		Where(userAction).
		Update("deleted", true).
		Error

	if err != nil {
		return err
	}

	return nil
}

func HaveUserAction(userAction userActionModel.UserAction) error {
	var resp userActionModel.UserAction
	userAction.Deleted = false

	if err := config.Db.Where(&userAction).First(&resp).Error; err != nil {
		return err
	}

	return nil
}
