package database

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&UserOnboarding{},
		&OnboardingAnswer{},
		&UserBlockedApp{},
		&UserUnlockTask{},
		&MorningResetSettings{},
		&DailyPlan{},
		&MorningResetSession{},
		&MorningResetTaskCompletion{},
	)
}
