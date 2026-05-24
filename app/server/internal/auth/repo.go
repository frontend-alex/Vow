package auth

import (
	"context"

	"github.com/vow/app/server/internal/platform/database"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type CreateUserParams struct {
	Email        string
	Name         string
	PasswordHash string
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r Repository) CreateUser(ctx context.Context, arg CreateUserParams) (database.User, error) {
	user := database.User{
		Email:        arg.Email,
		Name:         arg.Name,
		PasswordHash: arg.PasswordHash,
	}

	err := r.db.WithContext(ctx).Create(&user).Error
	return user, err
}

func (r Repository) GetUserByEmail(ctx context.Context, email string) (database.User, error) {
	var user database.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return user, err
}
