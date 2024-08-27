package domain
import "time"

type Log struct {
	ID        string    `json:"id"`
	Action    string    `json:"action"`
	UserID    string    `json:"user_id"`
	Timestamp time.Time `json:"timestamp"`
	Details   string    `json:"details"`
}