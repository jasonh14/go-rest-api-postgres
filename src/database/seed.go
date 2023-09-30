package database

import (
	"app/src/model"
	"app/src/model/constant"

	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	db.AutoMigrate(&model.MenuItem{}, &model.Order{}, &model.ProductOrder{}, &model.User{})

	foodMenu := []model.MenuItem{
		{
			Name:      "Pizza",
			OrderCode: "P001",
			Price:     100,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Burger",
			OrderCode: "B001",
			Price:     100,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Fries",
			OrderCode: "F001",
			Price:     100,
			Type:      constant.MenuTypeFood,
		},
	}

	drinkMenu := []model.MenuItem{
		{
			Name:      "Tea",
			OrderCode: "T001",
			Price:     100,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Coffee",
			OrderCode: "C001",
			Price:     100,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Milk",
			OrderCode: "M001",
			Price:     100,
			Type:      constant.MenuTypeDrink,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
}
