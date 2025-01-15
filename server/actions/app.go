package actions

import (
	"net/http"
	"server/handler"
	"server/models"
	"sync"

	_ "server/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var (
	app     *echo.Echo
	appOnce sync.Once
)

// @title MonDash API
// @version 0.0.1
// @description MonDash API SERVER Swagger
// @host localhost:8080
// @BasePath /
func App() *echo.Echo {
	appOnce.Do(func() {
		app = echo.New()

		app.Use(corsMiddleware)
		app.Use(loggerMiddleware)
		app.Use(TransactionMiddleware(models.Db))
		app.Use(ErrorMiddleware())
		app.Use(middleware.Recover())

		app.GET("/", readyz)
		app.GET("/swagger/*", echoSwagger.WrapHandler)

		app.GET("/ws", handler.WS)

		app.GET("/info", handler.GetServerInfo)
		app.GET("/cpu", handler.GetCpuStatusHandler)
		app.GET("/mem", handler.GetMemStatusHandler)
		app.GET("/net", handler.GetNetStatusHandler)
	})
	return app
}

func readyz(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"status": "OK"})
}
