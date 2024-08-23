package models

import (
	"database/sql"
	"errors"
	"net/http"
	"rest-api-echo/db"
	"rest-api-echo/utils"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleId   int    `json:"role_id"`
	RoleName string `json:"role_name"`
}

func CheckRoleAvailable(roleId int) bool {
	con := db.CreateCon()
	sqlStatement := "SELECT * FROM roles WHERE id = ?"

	var role Role
	err := con.QueryRow(sqlStatement, roleId).Scan(&role.Id, &role.Name, &role.CreatedAt)

	return err == nil
}

func FetchAllUsers() (Response, error) {
	var user User
	var users []User
	var res Response

	con := db.CreateCon()
	sqlStatement := "SELECT us.*, r.name as role_name FROM users as us LEFT JOIN roles as r ON us.role_id = r.id"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.RoleId, &user.RoleName)
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

func CreateUser(email string, username string, password string, roleId int) (Response, error) {
	var user User
	var res Response
	con := db.CreateCon()

	if !(CheckRoleAvailable(roleId)) {
		return res, errors.New("Role not available")
	}

	sqlStatement := "INSERT INTO users(email, username, password, role_id) VALUES(?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(email, username, hashedPassword, roleId)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	err = con.QueryRow("SELECT us.*, r.name as role_name FROM users as us LEFT JOIN roles as r ON us.role_id = r.id WHERE us.id = ?", lastInsertedId).Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.RoleId, &user.RoleName)
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

func GetUserByUsername(username string) (User, error) {
	var user User
	con := db.CreateCon()
	sqlStatement := "SELECT us.*, r.name as role_name FROM users as us LEFT JOIN roles as r ON us.role_id = r.id WHERE username = ?"
	err := con.QueryRow(sqlStatement, username).Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.RoleId, &user.RoleName)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("User Not Found")
		}
		return user, err
	}

	return user, nil
}
