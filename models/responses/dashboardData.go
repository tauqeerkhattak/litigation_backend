package responses

type DashboardData struct {
	ActiveCases   int64 `json:"active_cases"`
	HearingsToday int64 `json:"hearings_today"`
	TotalUsers    int64 `json:"total_users"`
	UrgentTasks   int64 `json:"urgent_tasks"`
}
