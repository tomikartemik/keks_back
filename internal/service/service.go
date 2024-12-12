package service

import (
	"tg_shop/internal/model"
	"tg_shop/internal/repository"
)

type Service struct {
	User
	Ad
	Category
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:     NewUserService(repos.User, repos.Ad, repos.Category),
		Ad:       NewAdService(repos.Ad, repos.User, repos.Category),
		Category: NewCategoryService(repos.Category),
	}
}

type User interface {
	CreateUser(id int, user model.User) (model.User, error)
	GetUserById(id int) (model.User, error)
	CreateOrUpdateUser(user model.User) (model.User, error)
	GetUserAsSellerByID(telegramIDStr string) (model.UserAsSeller, error)
}

type Ad interface {
	GetAdList(categoryIDStr string) ([]model.AdShortInfo, error)
	GetAdByID(idStr string) (model.AdInfo, error)
}

type Category interface {
	GetCategoryList() ([]model.Category, error)
}
