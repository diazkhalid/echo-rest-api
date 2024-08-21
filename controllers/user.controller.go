package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api-echo/models"

	"github.com/labstack/echo/v4"
)

func FetchAllUsers(c echo.Context) error {
	result, err := models.FetchAllUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func CreateUser(c echo.Context) error {
	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.CreateUser(jsonBody["email"].(string), jsonBody["username"].(string), jsonBody["password"].(string))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
