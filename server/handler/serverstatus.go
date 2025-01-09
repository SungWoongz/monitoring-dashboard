package handler

import (
	"net/http"
	"server/controller"
	"server/models"

	"github.com/labstack/echo/v4"
)

func GetCpuStatusHandler(c echo.Context) error {
	v := models.HttpResponseStatusOk()
	v.Data = controller.GetCpuStatus()
	return c.JSON(http.StatusOK, v)
}
