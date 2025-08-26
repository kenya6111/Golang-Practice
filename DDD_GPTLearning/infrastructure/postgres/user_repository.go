package postgres

import (
	"database/sql"
	domain "ddd_gpt_learning/domain/user"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Save(u *domain.User) error {
	query := `INSERT INTO users (id, username, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)`

	_, err := r.db.Exec(query, u.ID, u.Username, u.Email, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}

func (r *UserRepository) FindAll() ([]*domain.User, error) {
	rows, err := r.db.Query(`SELECT id, username, email, created_at, updated_at FROM users`)
	if err != nil {
		return nil, fmt.Errorf("failed to select users: %w", err)
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		u := &domain.User{}
		err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
