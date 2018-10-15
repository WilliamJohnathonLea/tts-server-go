package controllers

import (
	"net/http"
	"github.com/labstack/echo"
)

// PostMessagePage renders the page where the user can send a message
func PostMessagePage(c echo.Context) error {
	return c.Render(http.StatusOK, "post_message", nil)
}

// PostMessage validates a form submission and sends the message for voice synthesis
func PostMessage(c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/post-message")
}
