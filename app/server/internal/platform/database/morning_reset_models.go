package database

import (
	"time"

	"gorm.io/datatypes"
)

type Platform string

const (
	PlatformIOS     Platform = "ios"
	PlatformAndroid Platform = "android"
	PlatformWeb     Platform = "web"
	PlatformDesktop Platform = "desktop"
)

type Difficulty string

const (
	DifficultyGentle   Difficulty = "gentle"
	DifficultyBalanced Difficulty = "balanced"
	DifficultyStrong   Difficulty = "strong"
)

type UnlockTaskType string

const (
	UnlockTaskWalkSteps      UnlockTaskType = "walk_steps"
	UnlockTaskDrinkWater     UnlockTaskType = "drink_water"
	UnlockTaskStretch        UnlockTaskType = "stretch"
	UnlockTaskScanQR         UnlockTaskType = "scan_qr"
	UnlockTaskWriteIntention UnlockTaskType = "write_intention"
	UnlockTaskBreathing      UnlockTaskType = "breathing"
	UnlockTaskCustom         UnlockTaskType = "custom"
)

type MorningResetStatus string

const (
	MorningResetStatusScheduled    MorningResetStatus = "scheduled"
	MorningResetStatusActive       MorningResetStatus = "active"
	MorningResetStatusCompleted    MorningResetStatus = "completed"
	MorningResetStatusAutoUnlocked MorningResetStatus = "auto_unlocked"
	MorningResetStatusAbandoned    MorningResetStatus = "abandoned"
	MorningResetStatusFailed       MorningResetStatus = "failed"
)

type UnlockMethod string

const (
	UnlockMethodTasksCompleted  UnlockMethod = "tasks_completed"
	UnlockMethodAutoUnlock      UnlockMethod = "auto_unlock"
	UnlockMethodManualOverride  UnlockMethod = "manual_override"
	UnlockMethodEmergencyUnlock UnlockMethod = "emergency_unlock"
)

type UserBlockedApp struct {
	ID int64 `gorm:"primaryKey"`

	UserID int64 `gorm:"not null;index;uniqueIndex:idx_user_app_platform"`
	User   User  `gorm:"constraint:OnDelete:CASCADE;"`

	AppIdentifier string   `gorm:"type:text;not null;uniqueIndex:idx_user_app_platform"`
	AppName       string   `gorm:"type:text;not null"`
	AppCategory   *string  `gorm:"type:text"`
	Platform      Platform `gorm:"type:text;not null;uniqueIndex:idx_user_app_platform"`

	IsActive bool `gorm:"not null;default:true"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type UserUnlockTask struct {
	ID int64 `gorm:"primaryKey"`

	UserID int64 `gorm:"not null;index"`
	User   User  `gorm:"constraint:OnDelete:CASCADE;"`

	TaskType UnlockTaskType `gorm:"type:text;not null"`

	Title       string  `gorm:"type:text;not null"`
	Description *string `gorm:"type:text"`

	DifficultyLevel int `gorm:"not null;default:1"`
	SortOrder       int `gorm:"not null;default:0"`

	RequiredValue *int
	Metadata      datatypes.JSON `gorm:"type:jsonb"`

	IsActive bool `gorm:"not null;default:true"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`

	Completions []MorningResetTaskCompletion `gorm:"foreignKey:UnlockTaskID"`
}

type MorningResetSettings struct {
	ID int64 `gorm:"primaryKey"`

	UserID int64 `gorm:"not null;uniqueIndex"`
	User   User  `gorm:"constraint:OnDelete:CASCADE;"`

	WakeTime   string         `gorm:"type:text;not null"`
	RepeatDays datatypes.JSON `gorm:"type:jsonb;not null"`

	LockStartsAfterAlarm   bool `gorm:"not null;default:true"`
	AutoUnlockEnabled      bool `gorm:"not null;default:true"`
	AutoUnlockAfterMinutes int  `gorm:"not null;default:60"`

	Difficulty Difficulty `gorm:"type:text;not null;default:'balanced'"`

	IsActive bool `gorm:"not null;default:true"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type DailyPlan struct {
	ID int64 `gorm:"primaryKey"`

	UserID int64 `gorm:"not null;index;uniqueIndex:idx_user_plan_date"`
	User   User  `gorm:"constraint:OnDelete:CASCADE;"`

	PlanDate time.Time `gorm:"not null;uniqueIndex:idx_user_plan_date"`

	Priority1 *string `gorm:"type:text"`
	Priority2 *string `gorm:"type:text"`
	Priority3 *string `gorm:"type:text"`

	MorningIntention    *string `gorm:"type:text"`
	DesiredMorningState *string `gorm:"type:text"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`

	Sessions []MorningResetSession `gorm:"foreignKey:DailyPlanID"`
}

type MorningResetSession struct {
	ID int64 `gorm:"primaryKey"`

	UserID int64 `gorm:"not null;index"`
	User   User  `gorm:"constraint:OnDelete:CASCADE;"`

	DailyPlanID *int64
	DailyPlan   *DailyPlan `gorm:"constraint:OnDelete:SET NULL;"`

	ScheduledFor     time.Time `gorm:"not null;index"`
	AlarmDismissedAt *time.Time
	LockStartedAt    *time.Time
	UnlockedAt       *time.Time

	UnlockMethod *UnlockMethod      `gorm:"type:text"`
	Status       MorningResetStatus `gorm:"type:text;not null;default:'scheduled';index"`

	CompletedTaskCount int `gorm:"not null;default:0"`
	RequiredTaskCount  int `gorm:"not null;default:0"`

	FirstBlockedAppAttemptAt *time.Time
	BlockedAppAttemptCount   int `gorm:"not null;default:0"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`

	TaskCompletions []MorningResetTaskCompletion `gorm:"foreignKey:SessionID"`
}

type MorningResetTaskCompletion struct {
	ID int64 `gorm:"primaryKey"`

	SessionID int64               `gorm:"not null;index"`
	Session   MorningResetSession `gorm:"constraint:OnDelete:CASCADE;"`

	UnlockTaskID int64          `gorm:"not null;index"`
	UnlockTask   UserUnlockTask `gorm:"constraint:OnDelete:CASCADE;"`

	CompletedAt *time.Time

	Status string `gorm:"type:text;not null;default:'pending'"`

	MeasuredValue *int
	Metadata      datatypes.JSON `gorm:"type:jsonb"`

	CreatedAt time.Time `gorm:"not null"`
}
