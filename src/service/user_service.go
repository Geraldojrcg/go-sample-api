package service

import (
	"github.com/geraldojrcg/go-sample-api/src/dto"
	"github.com/geraldojrcg/go-sample-api/src/model"
	"github.com/geraldojrcg/go-sample-api/src/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	Db *gorm.DB
}

func (u UserService) GetAll() ([]model.User, error) {
	var users []model.User

	if err := u.Db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u UserService) GetById(id uuid.UUID) (model.User, error) {
	var user model.User

	if err := u.Db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u UserService) Create(dto dto.CreateUserDto) error {
	hash, err := utils.HashPassword(dto.Password)
	if err != nil {
		return err
	}

	user := model.User{
		BaseModel: model.BaseModel{
			ID: uuid.New(),
		},
		Name:     dto.Name,
		Email:    dto.Email,
		Password: hash,
	}

	result := u.Db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u UserService) Update(id uuid.UUID, dto dto.UpdateUserDto) (model.User, error) {
	user, err := u.GetById(id)
	if err != nil {
		return user, err
	}

	if result := u.Db.Model(&user).Updates(dto); result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (u UserService) Delete(id uuid.UUID) error {
	if result := u.Db.Delete(&model.User{}, id); result.Error != nil {
		return result.Error
	}

	return nil
}
