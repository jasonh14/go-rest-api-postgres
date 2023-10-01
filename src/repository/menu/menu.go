package menu

import (
	"app/src/model"
	"app/src/tracing"
	"context"

	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &menuRepo{db: db}
}

func (m *menuRepo) GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenuList")
	defer span.End()

	var menuData []model.MenuItem
	result := m.db.WithContext(ctx).Where(model.MenuItem{Type: model.MenuType(menuType)}).Find(&menuData)
	if result.Error != nil {
		// Handle the GORM error and return it as an error type
		return nil, result.Error
	}

	return menuData, nil
}

func (m *menuRepo) GetMenu(ctx context.Context, orderCode string) (model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenu")
	defer span.End()

	var menuData model.MenuItem
	result := m.db.WithContext(ctx).Where(model.MenuItem{OrderCode: orderCode}).Find(&menuData)
	if result.Error != nil {
		// Handle the GORM error and return it as an error type
		return menuData, result.Error
	}

	return menuData, nil
}
