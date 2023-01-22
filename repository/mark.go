package repository

import (
	"database/sql"
	"finaltask/structs"
	"time"
)

func CreateMark(db *sql.DB, mark structs.Mark) (err error) {
	sqlStatement := `
	INSERT INTO class (class_id,,ark,student_id,created_at,updated_at)
	VALUES ($1,$2,$3,$4,$5)
	Returning *
	`
	err = db.QueryRow(sqlStatement, mark.ClassID, mark.Mark, mark.StudentID, time.Now(), time.Now()).
		Scan(&mark.ID, &mark.ClassID, &mark.Mark, &mark.StudentID, &mark.CreatedAt, &mark.UpdatedAt)

	if err != nil {
		panic(err)
	}
	return
}

func GetMarksByID(db *sql.DB, id int) (err error, results []structs.Mark) {
	sqlStatement := `
	SELECT *
	FROM marks
	WHERE class_id = $1;
	`

	rows, err := db.Query(sqlStatement, id)

	defer rows.Close()

	for rows.Next() {
		var mark = structs.Mark{}

		err = rows.Scan(&mark.ID, &mark.ClassID, &mark.Mark, &mark.StudentID, &mark.CreatedAt, &mark.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, mark)
	}
	return
}

func UpdateMark(db *sql.DB, mark structs.Mark) (err error) {
	sql := "UPDATE mark SET mark=$1,updated-at =$2 WHERE id = $3"

	errs := db.QueryRow(sql, mark, time.Now(), mark.ID)

	return errs.Err()
}

func DeleteMark(db *sql.DB, id int) (err error) {
	sql := "DELETE FROM mark WHERE id = $1"

	errs := db.QueryRow(sql, id)

	return errs.Err()
}
