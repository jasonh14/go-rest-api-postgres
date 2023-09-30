package rest

import (
	"github.com/labstack/echo/v4"
)

func LoadRouters(e *echo.Echo, handler *handler) {
	authMiddleware := GetAuthMiddlewares(handler.restoUseCase)

	menuGroup := e.Group("/menu")
	menuGroup.GET("", handler.GetMenuList)

	orderGroup := e.Group("/order")
	orderGroup.POST("", handler.Order, authMiddleware.CheckAuth)
	orderGroup.GET("/:order_id", handler.GetOrderInfo, authMiddleware.CheckAuth)

	userGroup := e.Group("/user")
	userGroup.POST("/register", handler.RegisterUser)
	userGroup.POST("/login", handler.Login)
}
