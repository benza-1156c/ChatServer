package usecases

import "chatserver/modules/auth/dto"

type AuthUsecases interface {
	Register(req *dto.RegisterReq) error
}
