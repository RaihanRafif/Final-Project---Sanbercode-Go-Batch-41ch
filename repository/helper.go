package repository

import (
	"database/sql"
	"finaltask/structs"
	"fmt"
)

func FindAccount(db *sql.DB, id int) (result []structs.Teacher, err error) {
	sqlStatement := `
	SELECT *
	FROM teacher
	WHERE id = $1;
	`
	rows, err := db.Query(sqlStatement, id)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var account = structs.Teacher{}

		err = rows.Scan(&account.ID, &account.Phone, &account.Username, &account.Password, &account.Email, &account.Role, &account.CreatedAt, &account.UpdatedAt)
		if err != nil {
			panic(err)
		}

		result = append(result, account)
	}
	return
}

func TeacherAuthorization(db *sql.DB, emailID string) (err error, result []structs.Teacher) {
	fmt.Println("EMAIL", emailID)

	sqlStatement := `
	SELECT *
	FROM teacher
	WHERE email = $1;
	`
	rows, err := db.Query(sqlStatement, emailID)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var account = structs.Teacher{}

		err = rows.Scan(&account.ID, &account.Phone, &account.Username, &account.Password, &account.Email, &account.Role, &account.CreatedAt, &account.UpdatedAt)
		if err != nil {
			panic(err)
		}

		result = append(result, account)
	}
	return
}

func TeacherAccessAuthorization(db *sql.DB, emailID string, classID int) (err error) {
	sqlStatement := `
	SELECT *
	FROM class
	LEFT JOIN teacher ON teacher.id = class.teacher_id
	WHERE teacher.email = $1 AND class.id = $2;
	`
	_, err = db.Query(sqlStatement, emailID, classID)

	if err != nil {
		panic(err)
	}
	return
}

func MemberAccessAuthorization(db *sql.DB, emailID string, classID int) (err error) {
	sqlStatement := `
	SELECT *
	FROM member
	LEFT JOIN student ON student.id = member.user_id
	WHERE student.email = $1 AND member.class_id = $2;
	`
	_, err = db.Query(sqlStatement, emailID, classID)

	if err != nil {
		panic(err)
	}
	return
}
