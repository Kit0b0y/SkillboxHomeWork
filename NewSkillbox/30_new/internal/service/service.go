package service

import (
	user_app "github.com/kit0b0y/skillboxHomeWork/NewSkillbox/30_new"
	"github.com/kit0b0y/skillboxHomeWork/NewSkillbox/30_new/internal/repository"
)

type User interface {
	CreateUser(user user_app.RequestCreate) (string, error)
	MakeFriends(sourceId, targetId string) (string, error)
	DeleteUser(id string) (string, error)
	GetFriends(id string) ([]string, error)
	UpdateAge(id, age string) (string, error)
}

type Service struct {
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
	}
}
