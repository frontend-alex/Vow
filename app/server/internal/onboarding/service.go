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

type Service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return Service{repository: repository}
}

func (s Service) Start(ctx context.Context, userID int64) (database.UserOnboarding, error) {
	completed, err := s.repository.GetCompletedOnboarding(ctx, userID)

	if err == nil {
		return completed, sharederrors.OnboardingErrors.AlreadyCompleted
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return database.UserOnboarding{}, sharederrors.OnboardingErrors.StartFailed
	}

	active, err := s.repository.GetActiveOnboarding(ctx, userID)
	if err == nil {
		return active, sharederrors.OnboardingErrors.AlreadyStarted
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return database.UserOnboarding{}, sharederrors.OnboardingErrors.StartFailed
	}

	onboarding := database.UserOnboarding{
		UserID:    userID,
		Status:    database.OnboardingStatusInProgress,
		Version:   1,
		StartedAt: time.Now(),
	}

	err = s.repository.CreateOnboarding(ctx, &onboarding)
	if err != nil {
		return database.UserOnboarding{}, sharederrors.OnboardingErrors.StartFailed
	}

	return onboarding, nil
}

func (s Service) Complete(ctx context.Context, userID int64, input CompleteOnboardingRequest) error {
	if err := s.repository.CompleteOnboarding(ctx, userID, input); err != nil {
		if apiError, ok := sharederrors.FromError(err); ok {
			return apiError
		}

		return sharederrors.OnboardingErrors.CompleteFailed
	}

	return nil
}

func buildOnboardingAnswers(userID int64, onboardingID int64, input CompleteOnboardingRequest) ([]database.OnboardingAnswer, error) {
	improvementGoalsJSON, err := json.Marshal(input.ImprovementGoals)
	if err != nil {
		return nil, err
	}

	blockedAppsJSON, err := json.Marshal(input.BlockedApps)
	if err != nil {
		return nil, err
	}

	unlockTasksJSON, err := json.Marshal(input.UnlockTasks)
	if err != nil {
		return nil, err
	}

	repeatDaysJSON, err := json.Marshal(input.RepeatDays)
	if err != nil {
		return nil, err
	}

	autoUnlockEnabled := input.AutoUnlockEnabled
	autoUnlockMinutes := input.AutoUnlockAfterMinutes

	return []database.OnboardingAnswer{
		{
			UserID:       userID,
			OnboardingID: onboardingID,
			QuestionKey:  "improvement_goals",
			AnswerType:   database.AnswerTypeMultiChoice,
			AnswerJSON:   datatypes.JSON(improvementGoalsJSON),
		},
		{
			UserID:       userID,
			OnboardingID: onboardingID,
			QuestionKey:  "risk_moment",
			AnswerType:   database.AnswerTypeSingleChoice,
			AnswerText:   optionalString(input.RiskMoment),
		},
		{
			UserID:       userID,
			OnboardingID: onboardingID,
			QuestionKey:  "blocked_apps",
			AnswerType:   database.AnswerTypeJSON,
			AnswerJSON:   datatypes.JSON(blockedAppsJSON),
		},
		{
			UserID:       userID,
			OnboardingID: onboardingID,
			QuestionKey:  "wake_time",
			AnswerType:   database.AnswerTypeText,
			AnswerText:   optionalString(input.WakeTime),
		},
		{
			UserID:       userID,
			OnboardingID: onboardingID,
			QuestionKey:  "repeat_days",
			AnswerType:   database.AnswerTypeJSON,
			AnswerJSON:   datatypes.JSON(repeatDaysJSON),
		},
		{
			UserID:       userID,
			OnboardingID: onboardingID,
			QuestionKey:  "unlock_tasks",
			AnswerType:   database.AnswerTypeJSON,
			AnswerJSON:   datatypes.JSON(unlockTasksJSON),
		},
		{
			UserID:       userID,
			OnboardingID: onboardingID,
			QuestionKey:  "difficulty",
			AnswerType:   database.AnswerTypeSingleChoice,
			AnswerText:   optionalString(input.Difficulty),
		},
		{
			UserID:        userID,
			OnboardingID:  onboardingID,
			QuestionKey:   "auto_unlock_enabled",
			AnswerType:    database.AnswerTypeBoolean,
			AnswerBoolean: &autoUnlockEnabled,
		},
		{
			UserID:       userID,
			OnboardingID: onboardingID,
			QuestionKey:  "auto_unlock_after_minutes",
			AnswerType:   database.AnswerTypeNumber,
			AnswerNumber: &autoUnlockMinutes,
		},
	}, nil
}

func buildBlockedApps(userID int64, input CompleteOnboardingRequest) []database.UserBlockedApp {
	apps := make([]database.UserBlockedApp, 0, len(input.BlockedApps))

	for _, app := range input.BlockedApps {
		apps = append(apps, database.UserBlockedApp{
			UserID:        userID,
			AppIdentifier: app.AppIdentifier,
			AppName:       app.AppName,
			AppCategory:   optionalString(app.AppCategory),
			Platform:      database.Platform(app.Platform),
			IsActive:      true,
		})
	}

	return apps
}

func buildUnlockTasks(userID int64, input CompleteOnboardingRequest) ([]database.UserUnlockTask, error) {
	tasks := make([]database.UserUnlockTask, 0, len(input.UnlockTasks))

	for index, task := range input.UnlockTasks {
		metadataJSON, err := json.Marshal(task.Metadata)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, database.UserUnlockTask{
			UserID:          userID,
			TaskType:        database.UnlockTaskType(task.TaskType),
			Title:           task.Title,
			Description:     optionalString(task.Description),
			DifficultyLevel: 1,
			SortOrder:       index,
			RequiredValue:   task.RequiredValue,
			Metadata:        datatypes.JSON(metadataJSON),
			IsActive:        true,
		})
	}

	return tasks, nil
}

func optionalString(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func beginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}
