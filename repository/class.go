package repository

import (
	"database/sql"
	"finaltask/structs"
	"time"
)

func CreateClass(db *sql.DB, class structs.Class) (err error) {
	sqlStatement := `
	INSERT INTO class (topic,max_marks,teacher_id,description,filename,created_at,updated_at)
	VALUES ($1,$2,$3,$4,$5,$6,$7)
	Returning *
	`
	err = db.QueryRow(sqlStatement, class.Topic, class.MaxMarks, class.TeacherID, class.Description, class.Filename, time.Now(), time.Now()).
		Scan(&class.ID, &class.Topic, &class.MaxMarks, &class.TeacherID, &class.Description, &class.Filename, &class.CreatedAt, &class.UpdatedAt)

	if err != nil {
		panic(err)
	}
	return
}

func GetAllClass(db *sql.DB) (results []structs.Class, err error) {
	sql := `
	SELECT id,topic,max_marks,teacher_id,description,COALESCE(filename, ''),created_at,updated_at FROM
	class
	`
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var class = structs.Class{}

		err = rows.Scan(&class.ID, &class.Topic, &class.MaxMarks, &class.TeacherID, &class.Description, &class.Filename, &class.CreatedAt, &class.UpdatedAt)

		if err != nil {
			panic(err)
		}

		results = append(results, class)
	}
	return
}

func GetAllClassByTeacherID(db *sql.DB, id int) (results []structs.Class, err error) {
	sql := `
	SELECT class.*  
	FROM class
	WHERE class.teacher_id = $1
	`
	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var Class = structs.Class{}
		err = rows.Scan(&Class.ID, &Class.Topic, &Class.MaxMarks, &Class.TeacherID, &Class.Description, &Class.Filename, &Class.CreatedAt, &Class.UpdatedAt)

		if err != nil {
			panic(err)
		}

		results = append(results, Class)

	}

	return
}

// func GetAllClassByStudentID(db *sql.DB, id int) (results []structs.Class, err error) {
// 	sql := `
// 	SELECT class_id FROM
// 	member
// 	WHERE user_id = $1;
// 	`
// 	rows, err := db.Query(sql, id)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for rows.Next() {
// 		var member = structs.Member{}
// 		err = rows.Scan(&member.ClassID)

// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Println(member)
// 	}

// 	return
// }

func TeacherGetClassByClassID(db *sql.DB, classID int) (err error, results []structs.GetClassByClassID, members []structs.MemberDetail) {
	sql := `
	SELECT id,topic,max_marks,description,filename
	FROM class
	WHERE class.id = $1;
	`

	rows, err := db.Query(sql, classID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var class = structs.GetClassByClassID{}

		err = rows.Scan(&class.ID, &class.Topic, &class.MaxMarks, &class.Description, &class.Filename)
		if err != nil {
			return
		}

		results = append(results, class)
	}

	sql = `
	SELECT member.user_id,users.username,users.email,COALESCE(files.filename, '')
	FROM member
	LEFT JOIN files ON files.user_id = member.user_id AND files.class_id = $1
	LEFT JOIN users ON users.id = member.user_id
	WHERE member.class_id = $1
	`
	rows, err = db.Query(sql, classID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var member = structs.MemberDetail{}

		err = rows.Scan(&member.UserID, &member.Username, &member.Email, &member.Filename)

		// err = rows.Scan(&GetClassByClassID.Class.ID, &GetClassByClassID.Class.Topic, &GetClassByClassID.Class.MaxMarks, &GetClassByClassID.Class.Description, &GetClassByClassID.Students)

		if err != nil {
			return
		}

		members = append(members, member)
	}

	return
}

// func GetClassByClassID(db *sql.DB, classID int, userID int) (err error, results []structs.GetClassByClassID) {
// 	sql := `
// 	SELECT class.id,class.topic,class.max_marks,class.description,files.filename
// 	FROM class
// 	LEFT JOIN member ON member.class_id = $1
// 	LEFT JOIN student ON student.id = member.user_id
// 	LEFT JOIN files ON files.user_id = student.id AND files.class_id = $1
// 	WHERE class.id = $1;
// 	`
// 	rows, err := db.Query(sql, classID)
// 	if err != nil {
// 		return
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		var GetClassByClassID = structs.GetClassByClassID{}
// 		// var ArrayOfStudent = structs.NewStudent{}

// 		// *&GetClassByClassID.Student.Password = ""
// 		// err = rows.Scan(&GetClassByClassID.Class.ID, &GetClassByClassID.Class.Topic, &GetClassByClassID.Class.MaxMarks, &GetClassByClassID.Class.TeacherID, &GetClassByClassID.Class.Description, &GetClassByClassID.Class.CreatedAt, &GetClassByClassID.Class.UpdatedAt, &GetClassByClassID.Member.ID, &GetClassByClassID.Member.ClassID, &GetClassByClassID.Member.UserID, &GetClassByClassID.Member.CreatedAt, &GetClassByClassID.Member.UpdatedAt, &GetClassByClassID.Student.ID, &GetClassByClassID.Student.Phone, &GetClassByClassID.Student, &GetClassByClassID.Student.Email, &GetClassByClassID.Student.Username, &GetClassByClassID.Student.Role, &GetClassByClassID.Student.CreatedAt, &GetClassByClassID.Student.UpdatedAt, &GetClassByClassID.File.ID, &GetClassByClassID.File.Filename, &GetClassByClassID.File.UserID, &GetClassByClassID.File.ClassID, &GetClassByClassID.File.CreatedAt, &GetClassByClassID.File.UpdatedAt)

// 		err = rows.Scan(&GetClassByClassID.ID, &GetClassByClassID.Topic, &GetClassByClassID.MaxMarks, &GetClassByClassID.Description, &GetClassByClassID.Filename)

// 		// err = rows.Scan(&GetClassByClassID.Class.ID, &GetClassByClassID.Class.Topic, &GetClassByClassID.Class.MaxMarks, &GetClassByClassID.Class.Description, &GetClassByClassID.Students)

// 		if err != nil {
// 			return
// 		}

// 		results = append(results, GetClassByClassID)
// 	}
// 	return
// }

// func GetClassByID(db *sql.DB, id int) (err error, result []structs.CustomClass) {
// 	sqlStatement := `
// 	SELECT class.*,COALESCE(files.filename, '')  FROM
// 	class
// 	LEFT JOIN files
// 	ON files.class_id = class.id
// 	WHERE class.id = $1
// 	`

// 	rows, err := db.Query(sqlStatement, id)

// 	defer rows.Close()

// 	for rows.Next() {
// 		var class = structs.CustomClass{}
// 		err = rows.Scan(&class.Class.ID, &class.Class.Topic, &class.Class.MaxMarks, &class.Class.TeacherID, &class.Class.Description, &class.Class.CreatedAt, &class.Class.UpdatedAt, &class.Filename)

// 		if err != nil {
// 			panic(err)
// 		}

// 		result = append(result, class)
// 	}
// 	return
// }

func UpdateClass(db *sql.DB, class structs.Class) (err error) {
	sql := "UPDATE class SET topic = $1,max_marks=$2,description=$3,filename=$4,updated_at = $5 WHERE id = $6"

	errs := db.QueryRow(sql, class.Topic, class.MaxMarks, class.Description, class.Filename, time.Now(), class.ID)

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

// STUDENT SECTION----------------------------------------------------------------------------------------

func StudentGetAllClassByStudentID(db *sql.DB, id int) (results []structs.StudentClassDetail, err error) {
	sql := `
	SELECT class.id,class.topic,class.max_marks,class.description,class.filename,user.username,,COALESCE(studentFiles.filename, ''),COALESCE(mark.mark, 0)
	FROM member
	LEFT JOIN class ON class.id = member.class_id
	LEFT JOIN user ON user.id = class.teacher_id
	LEFT JOIN files ON files.class_id = member.class_id
	LEFT JOIN mark ON mark.class_id = member.class_id AND mark.student_id = member.user_id
	WHERE member.user_id = $1
	`
	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var data = structs.StudentClassDetail{}
		err = rows.Scan(&data.ID, &data.Topic, &data.MaxMarks, &data.Description, &data.Username, &data.ClassFile, &data.StudentFile, &data.Mark)
		if err != nil {
			panic(err)
		}
		results = append(results, data)
	}

	return
}
