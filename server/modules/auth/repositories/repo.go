package repositories

import (
	"chatserver/modules/auth/dto"

	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return &authRepo{db: db}
}

func (r *authRepo) FindUserByEmail(email string) (bool, error) {
	var count int64
	if err := r.db.Raw("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *authRepo) CreateUser(data *dto.RegisterReq, avatarURL *string) error {
	return r.db.Exec(
		"INSERT INTO users (user_name, email, password, avatar) VALUES (?, ?, ?, ?)",
		data.UserName, data.Email, data.Password, avatarURL,
	).Error
}
