package rest

import "app/src/usecase/resto"

type handler struct {
	restoUseCase resto.Usecase
}

func NewHandler(restoUsecase resto.Usecase) *handler {
	return &handler{
		restoUseCase: restoUsecase,
	}
}
