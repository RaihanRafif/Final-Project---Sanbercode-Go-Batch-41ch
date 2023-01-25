package structs

import "time"

type Class struct {
	ID          int64     `json:"id"`
	Topic       string    `json:"topic"`
	MaxMarks    int64     `json:"max_marks"`
	TeacherID   int64     `json:"teacher_id"`
	Description string    `json:"description"`
	Filename    string    `json:"filename"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CustomClass struct {
	Class    Class
	Filename string `json:"filename"`
}

// type GetClassByClassID struct {
// 	Class    newClass
// 	Students []NewStudent
// }

// type newClass struct {
// 	ID          int64  `json:"id"`
// 	Topic       string `json:"topic"`
// 	MaxMarks    int64  `json:"max_marks"`
// 	Description string `json:"description"`
// }

// type NewStudent struct {
// 	Phone    int64  `json:"phone"`
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// 	File     newFile
// }

// type newFile struct {
// 	Filename string `json:"filename"`
// }

type GetClassByClassID struct {
	ID          int64  `json:"id"`
	Topic       string `json:"topic"`
	MaxMarks    int64  `json:"max_marks"`
	Description string `json:"description"`
	Filename    string `json:"filename"`
}
