package controller

import (
	"errors"
	"github.com/geraldojrcg/go-sample-api/src/dto"
	"github.com/geraldojrcg/go-sample-api/src/service"
	"github.com/geraldojrcg/go-sample-api/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserController struct {
	Service service.UserService
}

func (u UserController) GetAll(c *fiber.Ctx) error {
	users, err := u.Service.GetAll()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (u UserController) GetById(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user, err := u.Service.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (u UserController) Create(c *fiber.Ctx) error {
	userDto := dto.CreateUserDto{}

	if err := c.BodyParser(&userDto); err != nil {
		return err
	}

	msg := utils.Validate(userDto)
	if len(msg) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(msg)
	}

	err := u.Service.Create(userDto)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (u UserController) Update(c *fiber.Ctx) error {
	updateDto := dto.UpdateUserDto{}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := c.BodyParser(&updateDto); err != nil {
		return err
	}

	msg := utils.Validate(updateDto)
	if len(msg) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(msg)
	}

	user, err := u.Service.Update(id, updateDto)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (u UserController) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = u.Service.Delete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
