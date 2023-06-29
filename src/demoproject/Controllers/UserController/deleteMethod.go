package usercontroller

import (
	controller "demoproject/Controllers"
	m "demoproject/Models"
	"net/http"

	"github.com/labstack/echo"
)

func Delete(ctx echo.Context) (err error) {

	user := &m.UpdateUser{}

	result := controller.Connection.First(user, ctx.Param("id"))

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": "User Not Found!"})
	}

	result = controller.Connection.Delete(user)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Faild to delete user!")
	}

	return ctx.JSON(http.StatusOK, echo.Map{"count": result.RowsAffected})
}
