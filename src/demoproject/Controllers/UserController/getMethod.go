package usercontroller

import (
	m "demoproject/Models"
	"net/http"

	controller "demoproject/Controllers"

	"github.com/labstack/echo"
)

func Index(ctx echo.Context) (err error) {
	users := &[]m.User{}

	results := controller.Connection.Find(users)

	if results.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Empty set!")
	}

	return ctx.JSON(http.StatusOK, users)
}

func Show(ctx echo.Context) (err error) {
	user := &m.User{}

	result := controller.Connection.Find(user, ctx.Param("id"))

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, echo.Map{"error": "User Not Found!"})
	}

	return ctx.JSON(http.StatusOK, user)
}
