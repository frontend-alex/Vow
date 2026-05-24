package onboarding

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

func (r Repository) WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return r.db.WithContext(ctx).Transaction(fn)
}

func (r Repository) CreateOnboarding(ctx context.Context, onboarding *database.UserOnboarding) error {
	return r.db.WithContext(ctx).Create(onboarding).Error
}

func (r Repository) GetActiveOnboarding(ctx context.Context, userID int64) (database.UserOnboarding, error) {
	var onboarding database.UserOnboarding

	err := r.db.WithContext(ctx).
		Where("user_id = ? AND status = ?", userID, database.OnboardingStatusInProgress).
		Order("created_at DESC").
		First(&onboarding).Error

	return onboarding, err
}

func (r Repository) GetCompletedOnboarding(ctx context.Context, userID int64) (database.UserOnboarding, error) {
	var onboarding database.UserOnboarding

	err := r.db.WithContext(ctx).
		Where("user_id = ? AND status = ?", userID, database.OnboardingStatusCompleted).
		Order("created_at DESC").
		First(&onboarding).Error

	return onboarding, err
}
