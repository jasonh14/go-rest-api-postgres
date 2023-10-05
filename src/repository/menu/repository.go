package menu

import (
	"app/src/model"
	"context"
)

//go:generate mockgen -package=mocks -mock_names=Repository=MockMenuRepository -destination=../../mocks/menu_respository_mock.go -source=repository.go

type Repository interface {
	GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error)
	GetMenu(ctx context.Context, orderCode string) (model.MenuItem, error)
}
