package controllers

import (
	"net/http"
	"github.com/labstack/echo"
)

// PostMessagePage renders the page where the user can send a message
func PostMessagePage(c echo.Context) error {
	return c.String(http.StatusOK, "Post a message!")
}
