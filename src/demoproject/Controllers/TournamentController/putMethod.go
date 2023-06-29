package tournamentcontroller

import (
	m "demoproject/Models"
	"net/http"

	controller "demoproject/Controllers"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func Update(ctx echo.Context) (err error) {

	tournament := &m.Tournament{}

	result := controller.Connection.First(tournament, ctx.Param("id"))

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": "Tournament Not Found!"})
	}

	if err = ctx.Bind(tournament); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	validator := validator.New()

	if err = validator.Struct(tournament); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	result = controller.Connection.Select("tournamentname", "UpdatedAt").Updates(tournament)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Faild to update Tournament!")
	}

	return ctx.JSON(http.StatusOK, echo.Map{"id": tournament.ID})
}
