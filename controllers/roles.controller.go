package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api-echo/models"

	"github.com/labstack/echo/v4"
)

func CreateRole(c echo.Context) error {
	jsonReqBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonReqBody)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.CreateRole(jsonReqBody["name"].(string))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, result)
}

func FetchRoleById(c echo.Context) error {
	id := c.Param("id")

	result, err := models.FetchRoleById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteRoleById(c echo.Context) error {
	id := c.Param("id")

	result, err := models.DeleteRoleById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchAllRoles(c echo.Context) error {
	result, err := models.FetchAllRoles()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
