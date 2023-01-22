package repository

import (
	"database/sql"
	"finaltask/structs"
	"time"
)

func CreateStudent(db *sql.DB, student structs.Student) (err error) {
	sqlStatement := `
	INSERT INTO student (phone,username,password,email,role,created_at,updated_at)
	VALUES ($1,$2,$3,$4,$5,$6,$7)
	Returning *
	`

	err = db.QueryRow(sqlStatement, student.Phone, student.Username, student.Password, student.Email, student.Role, time.Now(), time.Now()).
		Scan(&student.ID, &student.Phone, &student.Username, &student.Password, &student.Email, &student.Role, &student.CreatedAt, &student.UpdatedAt)

	if err != nil {
		panic(err)
	}
	return
}

func LoginStudent(db *sql.DB, student structs.Student) (err error, result []structs.Student) {
	sqlStatement := `
	SELECT id,email,password
	FROM student
	WHERE email = $1;
	`
	rows, err := db.Query(sqlStatement, student.Email)

	defer rows.Close()
	for rows.Next() {
		var student = structs.Student{}
		err = rows.Scan(&student.ID, &student.Email, &student.Password)
		if err != nil {
			panic(err)
		}
		result = append(result, student)
	}
	return
}

func UpdateStudent(db *sql.DB, id int, student structs.Student) (err error) {
	sql := "UPDATE student SET username= $1,phone=$2,password=$3,email=$4,updated_at=$5 WHERE id = $6"

	errs := db.QueryRow(sql, student.Username, student.Phone, student.Password, student.Email, time.Now(), id)

	return errs.Err()
}

func DeleteStudent(db *sql.DB, id int) (err error) {
	sql := "DELETE FROM student WHERE id = $1"

	errs := db.QueryRow(sql, id)

	return errs.Err()
}
