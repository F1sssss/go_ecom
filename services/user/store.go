package user

import (
	"database/sql"

	"github.com/F1sssss/go_ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (s *Store) GetUserByUsernameOrEmail(usernameOrEmail string) (*types.User, error) {
	return nil, nil
}

func (s *Store) CreateUser(u types.User) error {
	return nil
}

func (s *Store) UpdateUser(u types.User) error {
	return nil
}
