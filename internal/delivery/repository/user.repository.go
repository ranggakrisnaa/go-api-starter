package repository

import (
	"go-api-starter/internal/delivery/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB  *gorm.DB
	Log *logrus.Logger
}

func NewUserRepository(db *gorm.DB, log *logrus.Logger) *UserRepository {
	return &UserRepository{
		DB:  db,
		Log: log,
	}
}

func (r *UserRepository) FindByToken(db *gorm.DB, user *entity.User, token string) error {
	return r.DB.Where("token = ?", token).First(user).Error
}
