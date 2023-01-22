package structs

import "time"

type Member struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	ClassID   int64     `json:"class_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
