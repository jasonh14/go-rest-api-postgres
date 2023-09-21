package main

import (
	"app/internal/database"
	"app/internal/delivery/rest"
	mRepo "app/internal/repository/menu"
	oRepo "app/internal/repository/order"
	rUseCase "app/internal/usecase/resto"

	"github.com/labstack/echo/v4"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=postgres dbname=go_resto_app sslmode=disable"
)

func main() {

	e := echo.New()

	db := database.GetDB(dbAddress)

	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetRepository(db)

	restoUsecase := rUseCase.GetuseCase(menuRepo, orderRepo)

	h := rest.NewHandler(restoUsecase)

	rest.LoadRouters(e, h)

	e.Logger.Fatal(e.Start(":14045"))
}
