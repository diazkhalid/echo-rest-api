package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api-echo/models"
	"rest-api-echo/utils"
	"time"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	var user models.User
	jsonReqBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonReqBody)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	result, err := models.GetUserByUsername(jsonReqBody["username"].(string))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	user = result
	match := utils.CheckPasswordHash(jsonReqBody["password"].(string), user.Password)
	if !match {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Wrong password"})
	}

	accessToken, err := utils.GenerateAccessToken(user.Id, user.Email, user.Username, user.RoleName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:     "session",
		Value:    accessToken,
		Expires:  expiration,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(c.Response(), &cookie)
	return c.JSON(http.StatusOK, map[string]string{"message": "Login success", "token": accessToken})
}

func Logout(c echo.Context) error {
	cookie := http.Cookie{
		Name:     "session",
		Value:    "",
		Expires:  time.Unix(0, 0),
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(c.Response(), &cookie)
	return c.JSON(http.StatusOK, map[string]string{"message": "Logout success"})
}
