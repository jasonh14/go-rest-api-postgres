package menu

import (
	"app/internal/model"

	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &menuRepo{db: db}
}

func (m *menuRepo) GetMenu(menuType string) ([]model.MenuItem, error) {
	var menuData []model.MenuItem
	result := m.db.Where(model.MenuItem{Type: model.MenuType(menuType)}).Find(&menuData)
	if result.Error != nil {
		// Handle the GORM error and return it as an error type
		return nil, result.Error
	}

	return menuData, nil
}
