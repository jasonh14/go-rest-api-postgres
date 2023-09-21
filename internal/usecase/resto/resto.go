package resto

import (
	"app/internal/model"
	"app/internal/repository/menu"
)

type restoUseCase struct {
	menuRepo menu.Repository
}

func GetuseCase(menuRepo menu.Repository) Usecase {
	return &restoUseCase{
		menuRepo: menuRepo,
	}
}

func (r *restoUseCase) GetMenu(menuType string) ([]model.MenuItem, error) {
	return r.menuRepo.GetMenu(menuType)
}
