package database

import (
	"time"

	"gorm.io/datatypes"
)

type OnboardingStatus string

const (
	OnboardingStatusInProgress OnboardingStatus = "in_progress"
	OnboardingStatusCompleted  OnboardingStatus = "completed"
	OnboardingStatusAbandoned  OnboardingStatus = "abandoned"
)

type OnboardingAnswerType string

const (
	AnswerTypeSingleChoice OnboardingAnswerType = "single_choice"
	AnswerTypeMultiChoice  OnboardingAnswerType = "multi_choice"
	AnswerTypeText         OnboardingAnswerType = "text"
	AnswerTypeNumber       OnboardingAnswerType = "number"
	AnswerTypeBoolean      OnboardingAnswerType = "boolean"
	AnswerTypeJSON         OnboardingAnswerType = "json"
)

type UserOnboarding struct {
	ID int64 `gorm:"primaryKey"`

	UserID int64 `gorm:"not null;index"`
	User   User  `gorm:"constraint:OnDelete:CASCADE;"`

	Status      OnboardingStatus `gorm:"type:text;not null;default:'in_progress'"`
	Version     int              `gorm:"not null;default:1"`
	CurrentStep *string          `gorm:"type:text"`

	StartedAt   time.Time `gorm:"not null"`
	CompletedAt *time.Time

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`

	Answers []OnboardingAnswer `gorm:"foreignKey:OnboardingID"`
}

type OnboardingAnswer struct {
	ID int64 `gorm:"primaryKey"`

	OnboardingID int64          `gorm:"not null;index"`
	Onboarding   UserOnboarding `gorm:"constraint:OnDelete:CASCADE;"`

	UserID int64 `gorm:"not null;index"`
	User   User  `gorm:"constraint:OnDelete:CASCADE;"`

	QuestionKey string               `gorm:"type:text;not null;index"`
	AnswerType  OnboardingAnswerType `gorm:"type:text;not null"`

	AnswerText    *string `gorm:"type:text"`
	AnswerNumber  *int
	AnswerBoolean *bool
	AnswerJSON    datatypes.JSON `gorm:"type:jsonb"`

	CreatedAt time.Time `gorm:"not null"`
}
