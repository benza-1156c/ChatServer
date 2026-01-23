package usecases

import (
	"chatserver/entities"
	"chatserver/modules/auth/dto"
	"chatserver/modules/auth/repositories"
	"chatserver/pkg/utils"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authUsecases struct {
	repo repositories.AuthRepo
}

func NewAuthUsecases(repo repositories.AuthRepo) AuthUsecases {
	return &authUsecases{
		repo: repo,
	}
}

func (u *authUsecases) Register(req *dto.RegisterReq) error {
	exuser, _ := u.repo.FindUserByEmail(req.Email)
	if exuser {
		return errors.New("อีเมลนี้ถูกใช้แล้ว")
	}

	password, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	newUser := &entities.User{
		ID:       uuid.New(),
		Username: req.UserName,
		Email:    req.Email,
		Password: password,
	}

	if err := u.repo.CreateUser(newUser); err != nil {
		return err
	}

	return nil
}

func (u *authUsecases) Login(data *dto.LoginReq) (string, error) {
	user, err := u.repo.FindOneUserByEmail(data.Email)
	if err != nil {
		return "", errors.New("ไม่พบผู้ใช้")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return "", errors.New("รหัสผ่านไม่ถูกต้อง")
	}

	accesstoken, err := utils.GenerateJWT(user.ID.String(), 700000*time.Hour)
	if err != nil {
		return "", err
	}

	return accesstoken, nil
}

func (u *authUsecases) FindUserByID(id uuid.UUID) (*entities.User, error) {
	cacheKey := utils.GenerateCacheKey("user", id.String())

	cacheUser, err := utils.RedisGet[entities.User](cacheKey)
	if err == nil {
		return cacheUser, nil
	}

	user, err := u.repo.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	go utils.RedisSet(cacheKey, user, 5*time.Minute)

	return user, nil
}
