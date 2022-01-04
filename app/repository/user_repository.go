package repository

import (
	"go_scaffold/app/model"
	"go_scaffold/global"
	"gorm.io/gorm"
)

type userRepository struct{}

type UserRepository interface {
	repository
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (u *userRepository) modelFilter(param map[string]interface{}) *gorm.DB {
	panic("implement me")
}

func (u *userRepository) Add(model *model.Model) error {
	tx := global.Database.Model(model.UserModel).Begin()

	if err := tx.Create(model.UserModel).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (u *userRepository) Update(params map[string]interface{}, model *model.Model) error {
	panic("implement me")
}

func (u *userRepository) Delete(params map[string]interface{}, model *model.Model) error {
	panic("implement me")
}

func (u *userRepository) List(params map[string]interface{}) map[string]interface{} {
	panic("implement me")
}

func (u *userRepository) All(params map[string]interface{}) map[string]interface{} {
	panic("implement me")
}
