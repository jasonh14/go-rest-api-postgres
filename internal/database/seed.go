package database

import (
	"app/internal/model"
	"app/internal/model/constant"
	"fmt"

	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
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

	if db.Migrator().HasTable(&model.MenuItem{}) {
		// The table exists, so don't auto-migrate
		fmt.Println("MenuItem table already exists.")
	} else {
		// The table doesn't exist, so auto-migrate it
		db.AutoMigrate(&model.MenuItem{})
		fmt.Println("MenuItem table created.")
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
}
