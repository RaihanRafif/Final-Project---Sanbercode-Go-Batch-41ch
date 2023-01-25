package repository

import (
	"database/sql"
	"finaltask/structs"
	"time"
)

func CreateUser(db *sql.DB, user structs.User) (err error) {
	sqlStatement := `
	INSERT INTO users (phone,username,password,email,role,created_at,updated_at)
	VALUES ($1,$2,$3,$4,$5,$6,$7)
	Returning *
	`

	err = db.QueryRow(sqlStatement, user.Phone, user.Username, user.Password, user.Email, user.Role, time.Now(), time.Now()).
		Scan(&user.ID, &user.Phone, &user.Username, &user.Password, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		panic(err)
	}
	return
}

func LoginUser(db *sql.DB, user structs.User) (err error, result []structs.User) {
	sqlStatement := `
	SELECT id,email,password,user
	FROM users
	WHERE email = $1;
	`
	rows, err := db.Query(sqlStatement, user.Email)

	defer rows.Close()

	for rows.Next() {
		var user = structs.User{}

		err = rows.Scan(&user.ID, &user.Email, &user.Password, &user.Role)
		if err != nil {
			panic("Username, Password , or Role is invalid")
		}

		result = append(result, user)
	}
	return
}

func UpdateUser(db *sql.DB, id int, user structs.User) (err error) {
	sql := "UPDATE users SET username= $1,phone=$2,password=$3,email=$4,updated_at=$5 WHERE id = $6"

	errs := db.QueryRow(sql, user.Username, user.Phone, user.Password, user.Email, time.Now(), id)

	return errs.Err()
}

func DeleteUser(db *sql.DB, id int) (err error) {
	sql := "DELETE FROM users WHERE id = $1"

	errs := db.QueryRow(sql, id)

	return errs.Err()
}
