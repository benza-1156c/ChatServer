package repositories

import "chatserver/modules/auth/dto"

type AuthRepo interface {
	FindUserByEmail(email string) (bool, error)
	CreateUser(data *dto.RegisterReq, avatarURL *string) error
}
