package rest

import (
	"app/src/model"
	"app/src/model/constant"
	"app/src/tracing"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (h *handler) Order(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "Order")
	defer span.End()

	var request model.OrderMenuRequest

	err := json.NewDecoder(c.Request().Body).Decode(&request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	userID := c.Request().Context().Value(constant.AuthContextKey).(string)
	request.UserID = userID

	orderData, err := h.restoUseCase.Order(ctx, request)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[delivery][rest][order_handler][Order] unable to create order")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"data": orderData})
}

func (h *handler) GetOrderInfo(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "GetOrderInfo")
	defer span.End()

	orderID := c.Param("order_id")
	userID := c.Request().Context().Value(constant.AuthContextKey).(string)

	orderData, err := h.restoUseCase.GetOrderInfo(ctx, model.GetOrderInfoRequest{UserID: userID, OrderID: orderID})

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err,
		}).Error("[delivery][rest][order_handler][GetOrderInfo] unable to get order info")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": orderData})
}
