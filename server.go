package main

import (
	"github.com/WilliamJohnathonLea/tts-server-go/controllers"
	"io"
	"html/template"
	"github.com/labstack/echo"
)

// Template is a custom template render to use with Echo
type Template struct {
	templates *template.Template
}

// Render generates the views from the templates
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	t := &Template{template.Must(template.ParseGlob("views/*.html"))}
	e.Renderer = t
	e.GET("/", controllers.Home)
	e.Logger.Fatal(e.Start(":9000"))
}
