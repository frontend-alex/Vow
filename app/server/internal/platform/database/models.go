package database

import "time"

type User struct {
	ID           int64     `gorm:"primaryKey"`
	Email        string    `gorm:"not null;uniqueIndex"`
	Name         string    `gorm:"not null"`
	PasswordHash string    `gorm:"not null"`

	OnboardingCompleted   bool       `gorm:"not null;default:false"`
	OnboardingCompletedAt *time.Time

	Timezone string `gorm:"not null;default:'UTC'"`
	Locale   string `gorm:"not null;default:'en'"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`

	Onboardings          []UserOnboarding
	OnboardingAnswers    []OnboardingAnswer
	BlockedApps          []UserBlockedApp
	UnlockTasks          []UserUnlockTask
	MorningResetSettings *MorningResetSettings
	DailyPlans           []DailyPlan
	MorningResetSessions []MorningResetSession
}