package structs

import "time"

type Teacher struct {
	ID        int64     `json:"id"`
	Phone     int64     `json:"phone"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
