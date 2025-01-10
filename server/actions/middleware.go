package actions

import (
	"fmt"
	"net/http"

	"server/models"
	v "server/variables"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
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

var corsMiddleware = middleware.CORSWithConfig(middleware.CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
})

func TransactionMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tx := db.Begin()
			if tx.Error != nil {
				fmt.Println(tx.Error.Error())
				return tx.Error
			}
			c.Set("tx", tx)
			err := next(c)
			if err != nil {
				fmt.Println(err.Error())
				tx.Rollback()
				return err
			}
			if commitErr := tx.Commit().Error; commitErr != nil {
				fmt.Println(commitErr.Error())
				return commitErr
			}
			return nil
		}
	}
}

func ErrorMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				if httpError, ok := err.(*echo.HTTPError); ok {
					return c.JSON(httpError.Code, models.HttpResponse{
						Status: models.HttpStatus{
							Code:    httpError.Code,
							Message: http.StatusText(httpError.Code),
						},
					})
				}
				return c.JSON(http.StatusInternalServerError, models.HttpResponseInternalServerError(err.Error()))
			}
			return nil
		}
	}
}
