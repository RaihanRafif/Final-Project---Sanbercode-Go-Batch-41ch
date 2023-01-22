package repository

import (
	"database/sql"
	"finaltask/structs"
	"time"
)

func CreateTeacher(db *sql.DB, teacher structs.Teacher) (err error) {
	sqlStatement := `
	INSERT INTO teacher (phone,username,password,email,role,created_at,updated_at)
	VALUES ($1,$2,$3,$4,$5,$6,$7)
	Returning *
	`

	err = db.QueryRow(sqlStatement, teacher.Phone, teacher.Username, teacher.Password, teacher.Email, teacher.Role, time.Now(), time.Now()).
		Scan(&teacher.ID, &teacher.Phone, &teacher.Username, &teacher.Password, &teacher.Email, &teacher.Role, &teacher.CreatedAt, &teacher.UpdatedAt)

	if err != nil {
		panic(err)
	}
	return
}

func LoginTeacher(db *sql.DB, teacher structs.Teacher) (err error, result []structs.Teacher) {
	sqlStatement := `
	SELECT id,email,password
	FROM teacher
	WHERE email = $1;
	`
	rows, err := db.Query(sqlStatement, teacher.Email)

	defer rows.Close()

	for rows.Next() {
		var teacher = structs.Teacher{}

		err = rows.Scan(&teacher.ID, &teacher.Email, &teacher.Password)
		if err != nil {
			panic(err)
		}

		result = append(result, teacher)
	}
	return
}

func UpdateTeacher(db *sql.DB, id int, teacher structs.Teacher) (err error) {
	sql := "UPDATE teacher SET username= $1,phone=$2,password=$3,email=$4,updated_at=$5 WHERE id = $6"

	errs := db.QueryRow(sql, teacher.Username, teacher.Phone, teacher.Password, teacher.Email, time.Now(), id)

	return errs.Err()
}

func DeleteTeacher(db *sql.DB, id int) (err error) {
	sql := "DELETE FROM teacher WHERE id = $1"

	errs := db.QueryRow(sql, id)

	return errs.Err()
}
