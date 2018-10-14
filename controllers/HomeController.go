package controllers

import (
	"net/http"
	"github.com/labstack/echo"
)

// Home renders tbe index page of the site
func Home(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}
