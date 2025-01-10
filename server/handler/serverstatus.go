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

func GetMemStatusHandler(c echo.Context) error {
	v := models.HttpResponseStatusOk()
	d, err := controller.GetMemStatus()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponseInternalServerError(err.Error()))
	}
	v.Data = d
	return c.JSON(http.StatusOK, v)
}
