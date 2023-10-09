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

type TodoController struct {
	Service service.TodoService
}

func (t TodoController) GetAll(c *fiber.Ctx) error {
	todos, err := t.Service.GetAll()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(todos)
}

func (t TodoController) GetById(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	todo, err := t.Service.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}

func (t TodoController) Create(c *fiber.Ctx) error {
	todoDto := dto.CreateTodoDto{}

	if err := c.BodyParser(&todoDto); err != nil {
		return err
	}

	msg := utils.Validate(todoDto)
	if len(msg) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(msg)
	}

	err := t.Service.Create(todoDto)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (t TodoController) Update(c *fiber.Ctx) error {
	updateDto := dto.UpdateTodoDto{}

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

	todo, err := t.Service.Update(id, updateDto)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(todo)
}

func (t TodoController) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = t.Service.Delete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
