package repository

import (
	"database/sql"
	"finaltask/structs"
	"time"
)

func CreateFiles(db *sql.DB, file structs.File) (err error) {
	sqlStatement := `
	INSERT INTO files (filename,user_id,class_id,created_at,updated_at)
	VALUES ($1,$2,$3,$4,$5)
	Returning *
	`
	err = db.QueryRow(sqlStatement, file.Filename, file.UserID, file.ClassID, time.Now(), time.Now()).
		Scan(&file.ID, &file.Filename, &file.UserID, &file.ClassID, &file.CreatedAt, &file.UpdatedAt)

	if err != nil {
		panic(err)
	}
	return
}

func UpdateFiles(db *sql.DB, file structs.File, id int) (err error) {
	sqlStatement := "UPDATE files SET filename=$1, updated_at=$2 WHERE class_id=$3"

	errs := db.QueryRow(sqlStatement, file.Filename, time.Now(), id)

	return errs.Err()
}

// func GetFileByID(db *sql.DB, id int) {}
