package auth

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/vow/app/server/internal/platform/database"
	"github.com/vow/app/server/internal/shared/apperror"
	"gorm.io/gorm"
)

const postgresUniqueViolationCode = "23505"

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

	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		if isUniqueConstraintError(err) {
			return database.User{}, ErrEmailAlreadyExists
		}
		return database.User{}, apperror.Internal()
	}

	return user, nil
}

func (r Repository) GetUserByEmail(ctx context.Context, email string) (database.User, error) {
	var user database.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return user, err
}

func isUniqueConstraintError(err error) bool {
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return true
	}

	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == postgresUniqueViolationCode
}
