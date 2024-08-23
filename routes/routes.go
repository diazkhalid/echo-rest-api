package routes

import (
	"net/http"
	"rest-api-echo/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", controllers.FetchAllUsers)
	e.POST("/users", controllers.CreateUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users", controllers.DeleteUser)

	e.POST("/roles", controllers.CreateRole)
	e.GET("/roles/:id", controllers.FetchRoleById)
	e.GET("/roles", controllers.FetchAllRoles)
	e.DELETE("/roles/:id", controllers.DeleteRoleById)

	e.POST("/login", controllers.Login)
	e.POST("/logout", controllers.Logout)

	return e
}
