package onboarding

type CompleteOnboardingRequest struct {
	ImprovementGoals []string `json:"improvementGoals"`
	RiskMoment       string   `json:"riskMoment"`

	BlockedApps []BlockedAppInput `json:"blockedApps"`

	WakeTime   string   `json:"wakeTime"`
	RepeatDays []string `json:"repeatDays"`

	UnlockTasks []UnlockTaskInput `json:"unlockTasks"`

	Difficulty string `json:"difficulty"`

	Priority1           string `json:"priority1"`
	Priority2           string `json:"priority2"`
	Priority3           string `json:"priority3"`
	MorningIntention    string `json:"morningIntention"`
	DesiredMorningState string `json:"desiredMorningState"`

	AutoUnlockEnabled      bool `json:"autoUnlockEnabled"`
	AutoUnlockAfterMinutes int  `json:"autoUnlockAfterMinutes"`
}

type BlockedAppInput struct {
	AppIdentifier string `json:"appIdentifier"`
	AppName       string `json:"appName"`
	AppCategory   string `json:"appCategory"`
	Platform      string `json:"platform"`
}

type UnlockTaskInput struct {
	TaskType      string         `json:"taskType"`
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	RequiredValue *int           `json:"requiredValue"`
	Metadata      map[string]any `json:"metadata"`
}
