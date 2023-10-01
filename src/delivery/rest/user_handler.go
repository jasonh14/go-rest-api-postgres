package rest

import (
	"app/src/model"
	"app/src/tracing"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) RegisterUser(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "RegisterUser")
	defer span.End()

	var request model.RegisterRequest

	err := json.NewDecoder(c.Request().Body).Decode(&request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	result, err := h.restoUseCase.RegisterUser(ctx, request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"data": result,
	})

}

func (h *handler) Login(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "Login")
	defer span.End()

	var request model.LoginRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error(), "test": "test1"})
	}

	userSession, err := h.restoUseCase.Login(ctx, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error(), "test": "test2"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": userSession,
	})

}
