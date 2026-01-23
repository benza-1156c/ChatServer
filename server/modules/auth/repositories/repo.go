package repositories

import (
	"chatserver/entities"

	"github.com/google/uuid"
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

func (r *authRepo) FindUserByID(id uuid.UUID) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
