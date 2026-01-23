package usecases

import (
	"chatserver/entities"
	"chatserver/modules/auth/dto"

	"github.com/google/uuid"
)

type AuthUsecases interface {
	Register(req *dto.RegisterReq) error
	Login(data *dto.LoginReq) (string, error)
	FindUserByID(id uuid.UUID) (*entities.User, error)
}
