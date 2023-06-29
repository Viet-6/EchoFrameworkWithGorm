package tournamentcontroller

import (
	m "demoproject/Models"
	"net/http"

	controller "demoproject/Controllers"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func Create(ctx echo.Context) (err error) {

	tournament := &m.Tournament{}

	user := &m.User{}

	result := controller.Connection.First(user, 2)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": "User Not Found!"})
	}

	if err = ctx.Bind(tournament); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	validator := validator.New()

	if err = validator.Struct(tournament); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	result = controller.Connection.Create(tournament)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Faild to create Tournament!")
	}

	controller.Connection.Model(user).Association("Tournaments").Append(tournament)

	return ctx.JSON(http.StatusOK, echo.Map{"id": tournament.ID})
}
