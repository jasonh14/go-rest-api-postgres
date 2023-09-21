package rest

import (
	"github.com/labstack/echo/v4"
)

func LoadRouters(e *echo.Echo, handler *handler) {
	e.GET("/menu", handler.GetMenuList)
	e.POST("/order", handler.Order)
	e.GET("/order/:order_id", handler.GetOrderInfo)
}
