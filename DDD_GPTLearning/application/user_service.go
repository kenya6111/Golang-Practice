package application

import (
	domain "ddd_gpt_learning/domain/user"
	"ddd_gpt_learning/infrastructure/postgres"
)

type UserService struct {
	repo *postgres.UserRepository
}

func NewUserService(repo *postgres.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) ResisterUser(username, email string) (*domain.User, error) {
	u := domain.NewUser(username, email)
	if err := s.repo.Save(u); err != nil {
		return nil, err
	}
	return u, nil
}

// GetUsers ユーザー一覧を取得するユースケース
func (s *UserService) GetUsers() ([]*domain.User, error) {
	return s.repo.FindAll()
}
