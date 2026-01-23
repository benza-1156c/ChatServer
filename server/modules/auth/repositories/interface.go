package repositories

import (
	"chatserver/entities"

	"github.com/google/uuid"
)

type AuthRepo interface {
	FindOneUserByEmail(email string) (*entities.User, error)
	FindUserByEmail(email string) (bool, error)
	CreateUser(data *entities.User) error
	FindUserByID(id uuid.UUID) (*entities.User, error)
}
