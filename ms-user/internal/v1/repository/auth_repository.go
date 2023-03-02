package repository

import (
	"gorm.io/gorm"
	"ms-workspace/ms-user/internal/v1/repository/model"
)

type UserFilter struct {
	Email       string
	PhoneNumber string
}

type AuthRepository interface {
	CreateUser(user *model.User) error
	FindFirstUser(filter UserFilter) (*model.User, error)
	CreateActiveCode(code *model.ActiveCode) error
}

type AuthRepositoryImpl struct {
	database *gorm.DB
}

func NewAuthRepository(database *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{database: database}
}

func (u *AuthRepositoryImpl) CreateUser(user *model.User) error {
	return u.database.Create(user).Error
}

func (u *AuthRepositoryImpl) FindFirstUser(filter UserFilter) (*model.User, error) {
	var user model.User

	query := u.database.Model(&model.User{})

	if filter.Email != "" {
		query = query.Or("email = ?", filter.Email)
	}

	if filter.PhoneNumber != "" {
		query = query.Or("phone_number = ?", filter.PhoneNumber)
	}

	if err := query.First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *AuthRepositoryImpl) CreateActiveCode(code *model.ActiveCode) error {
	return u.database.Create(code).Error
}
