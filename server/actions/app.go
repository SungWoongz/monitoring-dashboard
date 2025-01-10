package actions

import (
	"net/http"
	"server/handler"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	app     *echo.Echo
	appOnce sync.Once
)

func App() *echo.Echo {
	appOnce.Do(func() {
		app = echo.New()

		app.Use(loggerMiddleware)
		app.Use(middleware.Recover())

		app.GET("/", readyz)

		app.GET("/ws", handler.WS)

		app.GET("/cpu", handler.GetCpuStatusHandler)
		app.GET("/mem", handler.GetMemStatusHandler)
	})
	return app
}

func readyz(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"status": "OK"})
}
