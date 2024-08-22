package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api-echo/models"
	"strconv"

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
	id := c.QueryParam("roleId")
	roleId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	jsonBody := make(map[string]interface{})
	err = json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.CreateUser(jsonBody["email"].(string), jsonBody["username"].(string), jsonBody["password"].(string), roleId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func UpdateUser(c echo.Context) error {
	jsonBody := make(map[string]interface{})
	id := c.Param("id")
	userID, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	err = json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.UpdateUser(userID, jsonBody["email"].(string), jsonBody["username"].(string), jsonBody["password"].(string))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func DeleteUser(c echo.Context) error {
	id := c.QueryParam("id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.DeleteUser(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
