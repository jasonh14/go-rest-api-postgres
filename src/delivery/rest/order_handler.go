package rest

import (
	"app/internal/model"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) Order(c echo.Context) error {
	var request model.OrderMenuRequest

	err := json.NewDecoder(c.Request().Body).Decode(&request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	orderData, err := h.restoUseCase.Order(request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"data": orderData})
}

func (h *handler) GetOrderInfo(c echo.Context) error {

	orderID := c.Param("order_id")
	orderData, err := h.restoUseCase.GetOrderInfo(model.GetOrderInfoRequest{OrderID: orderID})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": orderData})
}
