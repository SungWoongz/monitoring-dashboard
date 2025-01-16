package handler

import (
	"context"
	"net/http"
	"server/controller"
	"server/models"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// @Summary WebSocket Endpoint
// @Description Handles WebSocket connections and streams system status updates (CPU, Memory, Network) every second.
// @Tags WebSocket
// @Success 200 {object} models.ALL_Status "System status updates are sent via WebSocket."
// @Router /ws [get]
func WS(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		c.Logger().Errorf("WebSocket Upgrade Error: %s", err.Error())
		return err
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	handleErrors := func(err error) {
		if err.Error() == "websocket: close sent" {
			c.Logger().Debugf("WebSocket closed: %v", err)
		} else if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
			c.Logger().Infof("WebSocket closed: %v", err)
		} else {
			c.Logger().Errorf("WebSocket error: %s", err.Error())
		}
	}

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				cpu, _ := controller.GetCpuCurrentStatus()
				mem, _ := controller.GetMemCurrentStatus()
				net, _ := controller.GetNetCurrentStatus()
				if err := conn.WriteJSON(models.ALL_Status{
					CPU: cpu,
					MEM: mem,
					NET: net,
				}); err != nil {
					handleErrors(err)
					cancel()
					return
				}
			}
		}
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			handleErrors(err)
			cancel()
			break
		}
	}

	return nil
}

// GetServerInfo godoc
// @Summary Get Server Information
// @Description Retrieves information about the server, including hostname, OS, uptime, platform, and virtualization details.
// @Tags Server
// @Accept json
// @Produce json
// @Success 200 {object} models.HttpResponse
// @Failure 500 {object} models.HttpResponse
// @Router /info [get]
func GetServerInfo(c echo.Context) error {
	v := models.HttpResponseStatusOk()
	i, err := controller.GetServerInfo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.HttpResponseInternalServerError(err.Error()))
	}
	v.Data = i
	return c.JSON(http.StatusOK, v)
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
