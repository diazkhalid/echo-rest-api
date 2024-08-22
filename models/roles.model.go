package models

import (
	"net/http"
	"rest-api-echo/db"
	"time"
)

type Role struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

func CreateRole(name string) (Response, error) {
	var role Role
	var res Response
	con := db.CreateCon()

	date := time.Now()
	createdAt := date.Format("2006-01-02 15:04:05")

	sqlStatement := "INSERT INTO roles(name, created_at) VALUES(?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(name, createdAt)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	err = con.QueryRow("SELECT * FROM roles WHERE id = ?", lastInsertedId).Scan(&role.Id, &role.Name, &role.CreatedAt)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusCreated
	res.Message = "Success"
	res.Data = role

	return res, err
}

func FetchRoleById(id string) (Response, error) {
	var role Role
	var res Response
	con := db.CreateCon()

	sqlStatement := "SELECT * FROM roles WHERE id = ?"

	err := con.QueryRow(sqlStatement, id).Scan(&role.Id, &role.Name, &role.CreatedAt)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = role

	return res, err
}

func DeleteRoleById(id string) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "DELETE FROM roles WHERE id = ?"

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

func FetchAllRoles() (Response, error) {
	var role Role
	var roles []Role
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM roles"

	rows, err := con.Query(sqlStatement)

	if err != nil {
		return res, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&role.Id, &role.Name, &role.CreatedAt)
		if err != nil {
			return res, err
		}
		roles = append(roles, role)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = roles

	return res, err
}
