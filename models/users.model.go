package models

import (
	"net/http"
	"rest-api-echo/db"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func FetchAllUsers() (Response, error) {
	var user User
	var users []User
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT * FROM users"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Email, &user.Username, &user.Password)
		if err != nil {
			return res, err
		}
		users = append(users, user)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = users

	return res, err
}

func CreateUser(email string, username string, password string) (Response, error) {
	var user User
	var res Response
	con := db.CreateCon()

	sqlStatement := "INSERT INTO users(email, username, password) VALUES(?,?,?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(email, username, password)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	err = con.QueryRow("SELECT * FROM users WHERE id = ?", lastInsertedId).Scan(&user.Id, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return res, err
	}
	res.Status = http.StatusCreated
	res.Message = "Success"
	res.Data = user

	return res, err
}

func UpdateUser(id int, email string, username string, password string) (Response, error) {
	var user User
	var res Response
	con := db.CreateCon()

	sqlStatement := "UPDATE users SET email = ?, username = ?, password = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(email, username, password, id)
	if err != nil {
		return res, err
	}

	err = con.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return res, err
	}
	res.Status = http.StatusCreated
	res.Message = "Success"
	res.Data = user

	return res, err
}

func DeleteUser(id int) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "DELETE FROM users WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}
	return res, err
}
