package menu

import (
	"app/internal/model"
)

type Repository interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
	GetMenu(orderCode string) (model.MenuItem, error)
}
