package repositories

import "chatserver/entities"

type AuthRepo interface {
	FindOneUserByEmail(email string) (*entities.User, error)
	FindUserByEmail(email string) (bool, error)
	CreateUser(data *entities.User) error
}
