package main

import (
	"app/internal/database"
	"app/internal/delivery/rest"
	mRepo "app/internal/repository/menu"
	oRepo "app/internal/repository/order"
	uRepo "app/internal/repository/user"
	rUseCase "app/internal/usecase/resto"
	"crypto/rand"
	"crypto/rsa"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=postgres dbname=go_resto_app sslmode=disable"
)

func main() {

	e := echo.New()

	db := database.GetDB(dbAddress)

	// For Production, put this in env or secure place
	secret := "00112233445566778899AABBCCDDEEFF"
	signKey, err := rsa.GenerateKey(rand.Reader, 4096)

	if err != nil {
		panic(err)
	}

	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetRepository(db)
	userRepo, err := uRepo.GetRepository(db, secret, 64*1024, 4, 32, 1, signKey, 60*time.Second)

	if err != nil {
		panic(err)
	}

	restoUsecase := rUseCase.GetuseCase(menuRepo, orderRepo, userRepo)

	h := rest.NewHandler(restoUsecase)

	rest.LoadMiddlewares(e)

	rest.LoadRouters(e, h)

	e.Logger.Fatal(e.Start(":14045"))
}
