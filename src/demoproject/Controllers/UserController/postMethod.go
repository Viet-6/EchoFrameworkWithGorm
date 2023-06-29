package usercontroller

import (
	controller "demoproject/Controllers"
	m "demoproject/Models"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func Create(ctx echo.Context) (err error) {

	user := &m.User{}

	if err = ctx.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	validator := validator.New()

	if err = validator.Struct(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	result := controller.Connection.Create(user)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Faild to create user!")
	}

	return ctx.JSON(http.StatusOK, echo.Map{"id": user.ID})
}
