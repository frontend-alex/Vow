package onboarding

type CompleteOnboardingRequest struct {
	ImprovementGoals []string `json:"improvementGoals" validate:"required,min=1,dive,required,max=100" sanitize:"trim"`
	RiskMoment       string   `json:"riskMoment" validate:"required,max=100" sanitize:"trim"`

	BlockedApps []BlockedAppInput `json:"blockedApps" validate:"required,dive"`

	WakeTime   string   `json:"wakeTime" validate:"required,max=20" sanitize:"trim"`
	RepeatDays []string `json:"repeatDays" validate:"required,min=1,dive,required,oneof=monday tuesday wednesday thursday friday saturday sunday" sanitize:"trim,lower"`

	UnlockTasks []UnlockTaskInput `json:"unlockTasks" validate:"required,dive"`

	Difficulty string `json:"difficulty" validate:"required,oneof=gentle balanced strong" sanitize:"trim,lower"`

	Priority1           string `json:"priority1" validate:"max=200" sanitize:"trim"`
	Priority2           string `json:"priority2" validate:"max=200" sanitize:"trim"`
	Priority3           string `json:"priority3" validate:"max=200" sanitize:"trim"`
	MorningIntention    string `json:"morningIntention" validate:"max=1000" sanitize:"trim"`
	DesiredMorningState string `json:"desiredMorningState" validate:"max=200" sanitize:"trim"`

	AutoUnlockEnabled      bool `json:"autoUnlockEnabled"`
	AutoUnlockAfterMinutes int  `json:"autoUnlockAfterMinutes" validate:"min=0,max=1440"`
}

type BlockedAppInput struct {
	AppIdentifier string `json:"appIdentifier" validate:"required,max=255" sanitize:"trim"`
	AppName       string `json:"appName" validate:"required,max=100" sanitize:"trim"`
	AppCategory   string `json:"appCategory" validate:"max=100" sanitize:"trim"`
	Platform      string `json:"platform" validate:"required,oneof=ios android web desktop" sanitize:"trim,lower"`
}

type UnlockTaskInput struct {
	TaskType      string         `json:"taskType" validate:"required,oneof=walk_steps drink_water stretch scan_qr write_intention breathing custom" sanitize:"trim,lower"`
	Title         string         `json:"title" validate:"required,max=100" sanitize:"trim"`
	Description   string         `json:"description" validate:"max=500" sanitize:"trim"`
	RequiredValue *int           `json:"requiredValue" validate:"omitempty,min=0"`
	Metadata      map[string]any `json:"metadata"`
}
