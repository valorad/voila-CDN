package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func FileIndex(c echo.Context) error {
	return c.String(http.StatusOK, "file works!")
}
