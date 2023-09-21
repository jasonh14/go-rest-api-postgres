package resto

import "app/internal/model"

type Usecase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
