package fileController

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"pood/v2/app/controllers/fileController/createFileService"
	"pood/v2/app/models"
	"pood/v2/app/services/tokenService"
	"pood/v2/config"
	"strconv"
	t "time"
)

type FileController struct{}

func NewActionController() *FileController {
	return &FileController{}
}

// CreateNewFile
// @Summary Загрузить файл
// @Accept  mpfd
// @Produce json
// @Tags    Files
// @Param   file  formData     file   false  "binary"
// @Success 201 {object} models.CreateFileResponse
// @Failure 401 {object} models.FailedResponse
// @Router  /file [post]
// @Security ApiKeyAuth
func (FileController) CreateNewFile(c *fiber.Ctx) error {
	envPath := os.Getenv("FILE_PATH")
	_, err := tokenService.CheckToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"detail": err.Error()})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	now := t.Now().Format("2006-01-02")
	_ = os.MkdirAll(fmt.Sprintf("%s/%s", envPath, now), 0770)

	path := fmt.Sprintf("%s/%d_%s", now, t.Now().UnixMilli(), file.Filename)
	err = c.SaveFile(file, fmt.Sprintf("%s/%s", os.Getenv("FILE_PATH"), path))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	respFile, err := createFileService.CreateFile(file, path)
	data, err := json.Marshal(*respFile)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	var response models.CreateFileResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"detail": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

// DeleteFile
// @Summary Удалить file
// @Description Удалить file по id
// @Accept  json
// @Produce json
// @Tags    Files
// @Success 204 {object} models.SuccessResponse
// @Failure 400 {object} models.FailedResponse
// @Failure 401 {object} models.FailedResponse
// @Param id path string true "id"
// @Router  /file/{id} [delete]
// @Security ApiKeyAuth
func (FileController) DeleteFile(c *fiber.Ctx) error {
	_, err := tokenService.CheckToken(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	fileId, err := strconv.Atoi(c.Params("id", "0"))
	if err != nil || fileId == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "action not found",
		})
	}

	err = config.Db.
		Model(models.File{}).
		Where(models.File{ID: uint(fileId)}).
		Update("delete_at", t.Now().Format("2006-01-02 15:04:05")).
		Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": "file not delete",
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"detail": "ok",
	})
}

func (FileController) GetFile(c *fiber.Ctx) error {
	envPath := os.Getenv("FILE_PATH")
	date := c.Params("date", "")
	name := c.Params("name", "")

	fullPath := fmt.Sprintf("%s/%s/%s", envPath, date, name)

	err := c.SendFile(fullPath, true)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"detail": "file not found"})
	}

	return nil
}
