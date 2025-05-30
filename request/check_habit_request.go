package request

type CheckHabitRequest struct {
	ID        int    `json:"id"`
	CheckDate string `json:"check_date"`
}
