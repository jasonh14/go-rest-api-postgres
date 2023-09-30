package rest

import "app/internal/usecase/resto"

type handler struct {
	restoUseCase resto.Usecase
}

func NewHandler(restoUsecase resto.Usecase) *handler {
	return &handler{
		restoUseCase: restoUsecase,
	}
}
