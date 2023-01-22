package structs

import "time"

type Mark struct {
	ID        int64     `json:"id"`
	ClassID   int64     `json:"class_id"`
	Mark      int64     `json:"mark"`
	StudentID int64     `json:"student_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
