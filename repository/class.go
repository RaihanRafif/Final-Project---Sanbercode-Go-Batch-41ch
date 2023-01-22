package repository

import (
	"database/sql"
	"finaltask/structs"
	"fmt"
	"time"
)

func CreateClass(db *sql.DB, class structs.Class, teacher_id int) (err error) {
	sqlStatement := `
	INSERT INTO class (topic,max_marks,teacher_id,description,created_at,updated_at)
	VALUES ($1,$2,$3,$4,$5,$6)
	Returning *
	`
	err = db.QueryRow(sqlStatement, class.Topic, class.MaxMarks, teacher_id, class.Description, time.Now(), time.Now()).
		Scan(&class.ID, &class.Topic, &class.MaxMarks, &class.TeacherID, &class.Description, &class.CreatedAt, &class.UpdatedAt)

	fmt.Println("PPP", err)

	if err != nil {
		panic(err)
	}
	return
}

func GetAllClass(db *sql.DB) (results []structs.Class, err error) {
	sql := `
	SELECT * FROM 
	class
	`
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var class = structs.Class{}

		err = rows.Scan(&class.ID, &class.Topic, &class.MaxMarks, &class.TeacherID, &class.Description, &class.CreatedAt, &class.UpdatedAt)

		if err != nil {
			panic(err)
		}

		results = append(results, class)
	}
	return
}

func GetAllClassByTeacherID(db *sql.DB, id int) (results []structs.Class, err error) {
	sql := `
	SELECT * FROM 
	class 
	WHERE teacher_id = $1
	`
	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var class = structs.Class{}
		err = rows.Scan(&class.ID, &class.Topic, &class.MaxMarks, &class.TeacherID, &class.Description, &class.CreatedAt, &class.UpdatedAt)

		if err != nil {
			panic(err)
		}

		results = append(results, class)
	}
	return
}

func GetClassByClassID(db *sql.DB, classID int) (err error, results []structs.GetClassByClassID) {
	sql := `
	SELECT class.id,class.topic,class.max_marks,class.description,student.phone,student.username,files.filename
	FROM class
	LEFT JOIN member ON member.class_id = $1
	LEFT JOIN student ON student.id = member.user_id
	LEFT JOIN files ON files.user_id = student.id AND files.class_id = $1
	WHERE class.id = $1;
	`
	rows, err := db.Query(sql, classID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var GetClassByClassID = structs.GetClassByClassID{}
		// var ArrayOfStudent = structs.NewStudent{}

		// *&GetClassByClassID.Student.Password = ""
		// err = rows.Scan(&GetClassByClassID.Class.ID, &GetClassByClassID.Class.Topic, &GetClassByClassID.Class.MaxMarks, &GetClassByClassID.Class.TeacherID, &GetClassByClassID.Class.Description, &GetClassByClassID.Class.CreatedAt, &GetClassByClassID.Class.UpdatedAt, &GetClassByClassID.Member.ID, &GetClassByClassID.Member.ClassID, &GetClassByClassID.Member.UserID, &GetClassByClassID.Member.CreatedAt, &GetClassByClassID.Member.UpdatedAt, &GetClassByClassID.Student.ID, &GetClassByClassID.Student.Phone, &GetClassByClassID.Student, &GetClassByClassID.Student.Email, &GetClassByClassID.Student.Username, &GetClassByClassID.Student.Role, &GetClassByClassID.Student.CreatedAt, &GetClassByClassID.Student.UpdatedAt, &GetClassByClassID.File.ID, &GetClassByClassID.File.Filename, &GetClassByClassID.File.UserID, &GetClassByClassID.File.ClassID, &GetClassByClassID.File.CreatedAt, &GetClassByClassID.File.UpdatedAt)

		err = rows.Scan(&GetClassByClassID.ID, &GetClassByClassID.Topic, &GetClassByClassID.MaxMarks, &GetClassByClassID.Description, &GetClassByClassID.Phone, &GetClassByClassID.Username, &GetClassByClassID.Filename)

		// err = rows.Scan(&GetClassByClassID.Class.ID, &GetClassByClassID.Class.Topic, &GetClassByClassID.Class.MaxMarks, &GetClassByClassID.Class.Description, &GetClassByClassID.Students)

		if err != nil {
			panic(err)
		}

		results = append(results, GetClassByClassID)
	}
	return
}

// func GetClassByID(db *sql.DB, id int) (err error, result []structs.ClassDataCustom) {
// 	sqlStatement := `
// 	SELECT *
// 	FROM class
// 	INNER JOIN Files ON Files.class_id = Class.id
// 	WHERE class.id = $1;
// 	`

// 	rows, err := db.Query(sqlStatement, id)

// 	defer rows.Close()

// 	for rows.Next() {
// 		var ClassDataCustom = structs.ClassDataCustom{}

// 		err = rows.Scan(&ClassDataCustom.Class.ID, &ClassDataCustom.Class.Topic, &ClassDataCustom.Class.MaxMarks, &ClassDataCustom.Class.TeacherID, &ClassDataCustom.Class.Description, &ClassDataCustom.Class.CreatedAt, &ClassDataCustom.Class.UpdatedAt, &ClassDataCustom.File.ID, &ClassDataCustom.File.Filename, &ClassDataCustom.File.UserID, &ClassDataCustom.File.ClassID, &ClassDataCustom.File.CreatedAt, &ClassDataCustom.File.UpdatedAt)
// 		if err != nil {
// 			panic(err)
// 		}

// 		result = append(result, ClassDataCustom)
// 	}
// 	return
// }

// func GetAllClassCustom(db *sql.DB) (err error, results []structs.ClassDataCustom) {
// 	sql := "SELECT * FROM class INNER JOIN Files ON Files.class_id = Class.id "
// 	rows, err := db.Query(sql)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var ClassDataCustom = structs.ClassDataCustom{}

// 		err = rows.Scan(&ClassDataCustom.Class.ID, &ClassDataCustom.Class.Topic, &ClassDataCustom.Class.MaxMarks, &ClassDataCustom.Class.TeacherID, &ClassDataCustom.Class.Description, &ClassDataCustom.Class.CreatedAt, &ClassDataCustom.Class.UpdatedAt, &ClassDataCustom.File.ID, &ClassDataCustom.File.Filename, &ClassDataCustom.File.UserID, &ClassDataCustom.File.ClassID, &ClassDataCustom.File.CreatedAt, &ClassDataCustom.File.UpdatedAt)

// 		if err != nil {
// 			panic(err)
// 		}

// 		results = append(results, ClassDataCustom)
// 	}
// 	return
// }

func UpdateClass(db *sql.DB, class structs.Class) (err error) {
	sql := "UPDATE class SET topic = $1,max_marks=$2,description=$3,updated_at = $4 WHERE id = $5"

	errs := db.QueryRow(sql, class.Topic, class.MaxMarks, class.Description, time.Now(), class.ID)

	return errs.Err()
}

func DeleteClass(db *sql.DB, id int) (err error) {
	sql := `
	DELETE FROM class c WHERE c.id=$1;
	`
	errs := db.QueryRow(sql, id)

	sqlMember := `
	DELETE 
	FROM member m
	WHERE m.class_id = $1;
	`
	errs = db.QueryRow(sqlMember, id)

	sql_ := `DELETE 
	FROM files f
	WHERE f.class_id = $1;`

	errs = db.QueryRow(sql_, id)

	return errs.Err()
}
