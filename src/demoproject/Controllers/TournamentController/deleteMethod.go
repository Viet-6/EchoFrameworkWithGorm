package tournamentcontroller

import (
	controller "demoproject/Controllers"
	m "demoproject/Models"
	"net/http"

	"github.com/labstack/echo"
)

func Delete(ctx echo.Context) (err error) {

	tournament := &m.Tournament{}

	result := controller.Connection.First(tournament, ctx.Param("id"))

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": "Tournament Not Found!"})
	}

	result = controller.Connection.Delete(tournament)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Faild to delete Tournament!")
	}

	return ctx.JSON(http.StatusOK, echo.Map{"count": result.RowsAffected})
}
