package usecases

import (
	"chatserver/modules/auth/dto"
	"chatserver/modules/auth/repositories"
	"context"
	"errors"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type authUsecases struct {
	repo repositories.AuthRepo
	cld  *cloudinary.Cloudinary
}

func NewAuthUsecases(repo repositories.AuthRepo, cld *cloudinary.Cloudinary) AuthUsecases {
	return &authUsecases{
		repo: repo,
		cld:  cld,
	}
}

func (u *authUsecases) Register(req *dto.RegisterReq) error {
	exuser, _ := u.repo.FindUserByEmail(req.Email)
	if exuser {
		return errors.New("อีเมลนี้ถูกใช้แล้ว")
	}

	var avatarURL *string
	if req.Avatar != nil {
		resp, err := u.cld.Upload.Upload(context.Background(), req.Avatar,
			uploader.UploadParams{
				Folder:   "chatserver/avatars",
				PublicID: fmt.Sprintf("user_%s", req.Email),
			})
		if err != nil {
			return errors.New("อัพโหลดรูปไม่สำเร็จ")
		}

		avatarURL = &resp.SecureURL
	}

	if err := u.repo.CreateUser(req, avatarURL); err != nil {
		return err
	}

	return nil
}
