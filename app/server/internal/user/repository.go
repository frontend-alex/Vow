package user

import (
	"context"

	"github.com/vow/app/server/internal/platform/database"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (r Repository) GetUser(ctx context.Context, id int64) (database.User, error) {
	var user database.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	return user, err
}
