package repository

import (
	"database/sql"
	"finaltask/structs"
	"time"
)

func CreateMember(db *sql.DB, member structs.Member) (err error) {
	sqlStatement := `
	INSERT INTO member (class_id,user_id,created_at,updated_at)
	VALUES ($1,$2,$3,$4)
	Returning *
	`
	err = db.QueryRow(sqlStatement, member.ClassID, member.UserID, time.Now(), time.Now()).
		Scan(&member.ID, &member.ClassID, &member.UserID, &member.CreatedAt, &member.UpdatedAt)

	if err != nil {
		panic(err)
	}
	return
}

func DeleteMember(db *sql.DB, id int) (err error) {
	sql := `
	DELETE FROM member WHERE member.id = $1;
	`
	errs := db.QueryRow(sql, id)
	return errs.Err()
}
