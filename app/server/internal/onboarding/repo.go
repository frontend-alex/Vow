package onboarding

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/vow/app/server/internal/platform/database"
	sharederrors "github.com/vow/app/server/internal/shared/errors"
	"gorm.io/datatypes"
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

func (r Repository) CompleteOnboarding(ctx context.Context, userID int64, input CompleteOnboardingRequest) error {
	now := time.Now()

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var user database.User
		if err := tx.Where("id = ?", userID).First(&user).Error; err != nil {
			return err
		}

		if user.OnboardingCompleted {
			return sharederrors.OnboardingErrors.AlreadyCompleted
		}

		var onboarding database.UserOnboarding
		if err := tx.
			Where("user_id = ? AND status = ?", userID, database.OnboardingStatusInProgress).
			Order("created_at DESC").
			First(&onboarding).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return sharederrors.OnboardingErrors.NotStarted
			}
			return err
		}

		onboarding.Status = database.OnboardingStatusCompleted
		onboarding.CompletedAt = &now
		if err := tx.Save(&onboarding).Error; err != nil {
			return err
		}

		answers, err := buildOnboardingAnswers(userID, onboarding.ID, input)
		if err != nil {
			return err
		}

		if err := tx.Create(&answers).Error; err != nil {
			return err
		}

		blockedApps := buildBlockedApps(userID, input)
		if len(blockedApps) > 0 {
			if err := tx.Create(&blockedApps).Error; err != nil {
				return err
			}
		}

		unlockTasks, err := buildUnlockTasks(userID, input)
		if err != nil {
			return err
		}

		if len(unlockTasks) > 0 {
			if err := tx.Create(&unlockTasks).Error; err != nil {
				return err
			}
		}

		repeatDaysJSON, err := json.Marshal(input.RepeatDays)
		if err != nil {
			return err
		}

		settings := database.MorningResetSettings{
			UserID:                 userID,
			WakeTime:               input.WakeTime,
			RepeatDays:             datatypes.JSON(repeatDaysJSON),
			LockStartsAfterAlarm:   true,
			AutoUnlockEnabled:      input.AutoUnlockEnabled,
			AutoUnlockAfterMinutes: input.AutoUnlockAfterMinutes,
			Difficulty:             database.Difficulty(input.Difficulty),
			IsActive:               true,
		}

		if settings.AutoUnlockAfterMinutes == 0 {
			settings.AutoUnlockAfterMinutes = 60
		}

		if settings.Difficulty == "" {
			settings.Difficulty = database.DifficultyBalanced
		}

		if err := tx.Create(&settings).Error; err != nil {
			return err
		}

		plan := database.DailyPlan{
			UserID:              userID,
			PlanDate:            beginningOfDay(now),
			Priority1:           optionalString(input.Priority1),
			Priority2:           optionalString(input.Priority2),
			Priority3:           optionalString(input.Priority3),
			MorningIntention:    optionalString(input.MorningIntention),
			DesiredMorningState: optionalString(input.DesiredMorningState),
		}

		if err := tx.Create(&plan).Error; err != nil {
			return err
		}

		return tx.Model(&database.User{}).
			Where("id = ?", userID).
			Updates(map[string]any{
				"onboarding_completed":    true,
				"onboarding_completed_at": now,
			}).Error
	})
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
