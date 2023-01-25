package structs

import "time"

type Student struct {
	ID        int64     `json:"id"`
	Phone     int64     `json:"phone"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StudentClassDetail struct {
	ID          int64  `json:"id"`
	Topic       string `json:"topic"`
	MaxMarks    int64  `json:"max_marks"`
	Description string `json:"description"`
	Username    string `json:"username"`
	ClassFile   string `json:"class_file"`
	Mark        int64  `json:"mark"`
	StudentFile string `json:"student_file"`
}
