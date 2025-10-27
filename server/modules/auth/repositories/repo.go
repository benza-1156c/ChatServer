package repositories

import (
	"chatserver/entities"

	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) AuthRepo {
	return &authRepo{db: db}
}

func (r *authRepo) FindOneUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepo) FindUserByEmail(email string) (bool, error) {
	var exists bool
	if err := r.db.
		Raw("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", email).
		Scan(&exists).
		Error; err != nil {
		return false, err
	}

	return exists, nil
}

func (r *authRepo) CreateUser(data *entities.User) error {
	return r.db.Create(data).Error
}
