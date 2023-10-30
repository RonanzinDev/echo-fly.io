package routes

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ronanzindev/echo-fly.io/pkg/adm"
	"github.com/ronanzindev/echo-fly.io/pkg/user"
)

type TemplateRegistry struct {
	templates *template.Template
}

func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func NewRouter() (e *echo.Echo) {
	e = echo.New()
	e.Debug = true
	e.Renderer = &TemplateRegistry{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})
	e.GET("/html", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{"name": "Gabriel"})
	})
	e.GET("/login", adm.AdmPage)
	e.POST("/login", adm.Login)
	e.GET("/users", user.GetUsers)
	return
}
