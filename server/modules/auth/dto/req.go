package dto

import "mime/multipart"

type RegisterReq struct {
	UserName string                `json:"username"`
	Email    string                `json:"email"`
	Password string                `json:"password"`
	Avatar   *multipart.FileHeader `json:"avatar"`
}
