package actions

import (
	"net/http"

	v "server/variables"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var loggerMiddleware = middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
	LogStatus: true, LogURI: true, LogMethod: true, LogRemoteIP: true, LogError: true,
	LogValuesFunc: func(c echo.Context, rv middleware.RequestLoggerValues) error {
		if rv.Error != nil {
			c.Logger().Infof("\n⇨ %v [%v] %v (%v %v) "+v.TCR+"Error: %v\n"+v.TCC, rv.RemoteIP, rv.Method, rv.URI, rv.Status, http.StatusText(rv.Status), rv.Error)
		} else {
			c.Logger().Infof("\n⇨ %v [%v] %v (%v %v)\n", rv.RemoteIP, rv.Method, rv.URI, rv.Status, http.StatusText(rv.Status))
		}
		return nil
	},
})
