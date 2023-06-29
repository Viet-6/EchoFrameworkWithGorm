package usercontroller

import (
	controller "demoproject/Controllers"
	m "demoproject/Models"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func Update(ctx echo.Context) (err error) {

	user := &m.UpdateUser{}

	result := controller.Connection.First(user, ctx.Param("id"))

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, echo.Map{"error": "User Not Found!"})
	}

	if err = ctx.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	validator := validator.New()

	if err = validator.Struct(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	result = controller.Connection.Select("Username", "UpdatedAt").Updates(user)

	if result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Faild to update user!")
	}

	return ctx.JSON(http.StatusOK, echo.Map{"id": user.ID})
}
