package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetMenuList(c echo.Context) error {
	menuType := c.FormValue("menu_type")

	menuData, err := h.restoUseCase.GetMenuList(menuType)

	if err != nil {
		fmt.Printf(err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": menuData})

}
