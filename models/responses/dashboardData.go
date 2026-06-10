package responses

type DashboardData struct {
	ActiveCases   int `json:"active_cases"`
	HearingsToday int `json:"hearings_today"`
	TotalUsers    int `json:"total_users"`
	UrgentTasks   int `json:"urgent_tasks"`
}
