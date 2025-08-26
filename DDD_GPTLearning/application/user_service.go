package application

import (
	"ddd_gpt_learning/domain/user"
)

type UserService struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) ResisterUser(username, email string) (*user.User, error) {
	u := user.NewUser(username, email)
	if err := s.repo.Save(u); err != nil {
		return nil, err
	}
	return u, nil
}

// GetUsers ユーザー一覧を取得するユースケース
func (s *UserService) GetUsers() ([]*user.User, error) {
	return s.repo.FindAll()
}
