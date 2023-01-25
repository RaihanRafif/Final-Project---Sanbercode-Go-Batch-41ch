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

func UpdateFiles(db *sql.DB, file structs.File, id int, fileId int) (err error) {
	sqlStatement := "UPDATE files SET filename=$1, updated_at=$2 WHERE class_id=$3 AND id = $4"

	errs := db.QueryRow(sqlStatement, file.Filename, time.Now(), id, fileId)

	return errs.Err()
}

func DeleteFile(db *sql.DB, id int, fileId int) (err error) {
	sql := `
	DELETE FROM files WHERE id=$1 AND class_id=$2;
	`
	errs := db.QueryRow(sql, fileId, id)

	return errs.Err()
}
