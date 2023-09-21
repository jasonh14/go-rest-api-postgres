package menu

import (
	"app/internal/model"
)

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
