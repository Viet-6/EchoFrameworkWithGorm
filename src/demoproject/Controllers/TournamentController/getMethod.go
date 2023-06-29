package tournamentcontroller

import (
	m "demoproject/Models"
	"net/http"

	controller "demoproject/Controllers"

	"github.com/labstack/echo"
)

func Index(ctx echo.Context) (err error) {

	tournaments := &[]m.Tournament{}

	results := controller.Connection.Find(tournaments)

	if results.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Empty set!")
	}

	return ctx.JSON(http.StatusOK, tournaments)
}

func Show(ctx echo.Context) (err error) {
	tournament := &m.Tournament{}

	result := controller.Connection.Find(tournament, ctx.Param("id"))

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, echo.Map{"error": "Tournament Not Found!"})
	}

	return ctx.JSON(http.StatusOK, tournament)
}
