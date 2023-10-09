package service

import (
	"github.com/geraldojrcg/go-sample-api/src/dto"
	"github.com/geraldojrcg/go-sample-api/src/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoService struct {
	Db          *gorm.DB
	UserService UserService
}

func (t TodoService) GetAll() ([]model.Todo, error) {
	var todos []model.Todo

	if err := t.Db.Model(model.Todo{}).Preload("User").Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (t TodoService) GetById(id uuid.UUID) (model.Todo, error) {
	var todo model.Todo

	if err := t.Db.Model(model.Todo{}).Preload("User").Where("id = ?", id).First(&todo).Error; err != nil {
		return todo, err
	}

	return todo, nil
}

func (t TodoService) Create(dto dto.CreateTodoDto) error {
	userId, err := uuid.Parse(dto.UserID)
	if err != nil {
		return err
	}

	_, err = t.UserService.GetById(userId)
	if err != nil {
		return err
	}

	todo := model.Todo{
		BaseModel: model.BaseModel{
			ID: uuid.New(),
		},
		Description: dto.Description,
		UserID:      userId,
	}

	result := t.Db.Create(&todo)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (t TodoService) Update(id uuid.UUID, dto dto.UpdateTodoDto) (model.Todo, error) {
	todo, err := t.GetById(id)
	if err != nil {
		return todo, err
	}

	if result := t.Db.Model(&todo).Updates(dto); result.Error != nil {
		return todo, result.Error
	}

	return todo, nil
}

func (t TodoService) Delete(id uuid.UUID) error {
	if result := t.Db.Delete(&model.Todo{}, id); result.Error != nil {
		return result.Error
	}

	return nil
}
