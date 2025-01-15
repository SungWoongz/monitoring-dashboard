package handler

import (
	"fmt"
	"net/http"
	"server/controller"
	"server/models"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	upgrader = websocket.Upgrader{}
)

func WS(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}

// GetCpuStatusHandler godoc
// @Summary Get CPU Status
// @Description Retrieves CPU status data based on the specified interval and limit.
// @Tags CPU
// @Accept json
// @Produce json
// @Param interval query string false "Interval in seconds (default: 1)"
// @Param limit query string false "Number of records to fetch (default: 30)"
// @Success 200 {object} models.HttpResponse
// @Failure 400 {object} models.HttpResponse
// @Failure 500 {object} models.HttpResponse
// @Router /cpu [get]
func GetCpuStatusHandler(c echo.Context) error {
	tx := c.Get("tx").(*gorm.DB)
	v := models.HttpResponseStatusOk()

	interval := c.QueryParam("interval")
	limit := c.QueryParam("limit")
	if interval == "" {
		interval = "1"
	}
	if limit == "" {
		limit = "30"
	}
	intervalInt, err := strconv.Atoi(interval)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponseBadRequest("Invalid interval parameter"))
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponseBadRequest("Invalid limit parameter"))
	}

	d, err := controller.GetCpuStatus(tx, intervalInt, limitInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponseInternalServerError(err.Error()))
	}
	v.Data = d
	return c.JSON(http.StatusOK, v)
}

// GetMemStatusHandler godoc
// @Summary Get Memory Status
// @Description Retrieves memory status data based on the specified interval and limit.
// @Tags Memory
// @Accept json
// @Produce json
// @Param interval query string false "Interval in seconds (default: 1)"
// @Param limit query string false "Number of records to fetch (default: 30)"
// @Success 200 {object} models.HttpResponse
// @Failure 400 {object} models.HttpResponse
// @Failure 500 {object} models.HttpResponse
// @Router /mem [get]
func GetMemStatusHandler(c echo.Context) error {
	tx := c.Get("tx").(*gorm.DB)
	v := models.HttpResponseStatusOk()

	interval := c.QueryParam("interval")
	limit := c.QueryParam("limit")
	if interval == "" {
		interval = "1"
	}
	if limit == "" {
		limit = "30"
	}
	intervalInt, err := strconv.Atoi(interval)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponseBadRequest("Invalid interval parameter"))
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponseBadRequest("Invalid limit parameter"))
	}

	d, err := controller.GetMemStatus(tx, intervalInt, limitInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponseInternalServerError(err.Error()))
	}
	v.Data = d
	return c.JSON(http.StatusOK, v)
}

// GetNetStatusHandler godoc
// @Summary Get Network Status
// @Description Retrieves network status data based on the specified interval and limit.
// @Tags Network
// @Accept json
// @Produce json
// @Param interval query string false "Interval in seconds (default: 1)"
// @Param limit query string false "Number of records to fetch (default: 30)"
// @Success 200 {object} models.HttpResponse
// @Failure 400 {object} models.HttpResponse
// @Failure 500 {object} models.HttpResponse
// @Router /net [get]
func GetNetStatusHandler(c echo.Context) error {
	tx := c.Get("tx").(*gorm.DB)
	v := models.HttpResponseStatusOk()

	interval := c.QueryParam("interval")
	limit := c.QueryParam("limit")
	if interval == "" {
		interval = "1"
	}
	if limit == "" {
		limit = "30"
	}
	intervalInt, err := strconv.Atoi(interval)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponseBadRequest("Invalid interval parameter"))
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.HttpResponseBadRequest("Invalid limit parameter"))
	}

	d, err := controller.GetNetStatus(tx, intervalInt, limitInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponseInternalServerError(err.Error()))
	}
	v.Data = d
	return c.JSON(http.StatusOK, v)
}
