package rest

import (
	"app/internal/model"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) RegisterUser(c echo.Context) error {
	var request model.RegisterRequest

	err := json.NewDecoder(c.Request().Body).Decode(&request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	result, err := h.restoUseCase.RegisterUser(request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"data": result,
	})

}

func (h *handler) Login(c echo.Context) error {
	var request model.LoginRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	userSession, err := h.restoUseCase.Login(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error(), "test": "test"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": userSession,
	})

}
