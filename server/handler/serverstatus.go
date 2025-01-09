package handler

import (
	"net/http"
	"server/controller"
	"server/models"

	"github.com/labstack/echo/v4"
)

func GetCpuStatusHandler(c echo.Context) error {
	v := models.HttpResponseStatusOk()
	d, err := controller.GetCpuStatus()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponseInternalServerError())
	}
	v.Data = d
	return c.JSON(http.StatusOK, v)
}

func GetMemStatusHandler(c echo.Context) error {
	v := models.HttpResponseStatusOk()
	d, err := controller.GetMemStatus()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponseInternalServerError())
	}
	v.Data = d
	return c.JSON(http.StatusOK, v)
}
