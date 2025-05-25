package api

import (
	"io"
	"net/http"

	"github.com/aquaticcalf/gocomp-play-repo/pages"
	"github.com/labstack/echo/v4"
	. "maragu.dev/gomponents"
)

var app *echo.Echo

func init() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		node := pages.Home()
		return c.HTML(200, RenderString(node))
	})

	e.GET("/*", func(c echo.Context) error {
		node := pages.NotFound()
		return c.HTML(404, RenderString(node))
	})

	app = e
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}

// RenderString renders a gomponents Node to a string
func RenderString(n Node) string {
	s, err := RenderStringE(n)
	if err != nil {
		return ""
	}
	return s
}

// RenderStringE renders a gomponents Node to a string, returning any error
func RenderStringE(n Node) (string, error) {
	r, err := Render(n)
	if err != nil {
		return "", err
	}
	defer r.Close()

	b, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// Render renders a gomponents Node to an io.Reader
func Render(n Node) (io.ReadCloser, error) {
	pr, pw := io.Pipe()
	go func() {
		err := n.Render(pw)
		pw.CloseWithError(err)
	}()
	return pr, nil
}
