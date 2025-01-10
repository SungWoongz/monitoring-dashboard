package handler

import (
	"fmt"
	"net/http"
	"server/controller"
	"server/models"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
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
