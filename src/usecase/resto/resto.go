package resto

import (
	"app/internal/model"
	"app/internal/model/constant"
	"app/internal/repository/menu"
	"app/internal/repository/order"
	"app/internal/repository/user"
	"errors"

	"github.com/google/uuid"
)

type restoUseCase struct {
	menuRepo  menu.Repository
	orderRepo order.Repository
	userRepo  user.Repository
}

func GetuseCase(menuRepo menu.Repository, orderRepo order.Repository, userRepo user.Repository) Usecase {
	return &restoUseCase{
		menuRepo:  menuRepo,
		orderRepo: orderRepo,
		userRepo:  userRepo,
	}
}

func (r *restoUseCase) GetMenuList(menuType string) ([]model.MenuItem, error) {
	return r.menuRepo.GetMenuList(menuType)
}

func (r *restoUseCase) Order(request model.OrderMenuRequest) (model.Order, error) {
	productOrderData := make([]model.ProductOrder, len(request.OrderProducts))

	for i, orderProduct := range request.OrderProducts {
		menuData, err := r.menuRepo.GetMenu(orderProduct.OrderCode)
		if err != nil {
			return model.Order{}, err
		}
		productOrderData[i] = model.ProductOrder{
			ID:         uuid.New().String(),
			OrderCode:  orderProduct.OrderCode,
			Quantity:   orderProduct.Quantity,
			TotalPrice: int64(orderProduct.Quantity) * int64(menuData.Price),
			Status:     constant.ProductOrderStatusPreparing,
		}

	}

	order := model.Order{
		ID:            uuid.New().String(),
		Status:        constant.OrderStatusProcessed,
		ProductOrders: productOrderData,
		ReferenceID:   request.ReferenceID,
	}

	createdOrder, err := r.orderRepo.CreateOrder(order)

	if err != nil {
		return model.Order{}, err
	}

	return createdOrder, nil

}

func (r *restoUseCase) GetOrderInfo(request model.GetOrderInfoRequest) (model.Order, error) {
	orderData, err := r.orderRepo.GetOrderInfo(request.OrderID)

	if err != nil {
		return orderData, err
	}

	return orderData, nil
}

func (r *restoUseCase) RegisterUser(request model.RegisterRequest) (model.User, error) {
	userRegistered, err := r.userRepo.CheckRegistered(request.Username)
	if err != nil {
		return model.User{}, err
	}

	if userRegistered {
		return model.User{}, errors.New("user already registered")
	}

	userHash, err := r.userRepo.GenerateUserHash(request.Password)

	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		ID:       uuid.New().String(),
		Username: request.Username,
		Hash:     userHash,
	}

	userData, err := r.userRepo.RegisterUser(user)

	if err != nil {
		return model.User{}, err
	}

	return userData, nil

}

func (r *restoUseCase) Login(request model.LoginRequest) (model.UserSession, error) {
	userData, err := r.userRepo.GetUserData(request.Username)

	if err != nil {
		return model.UserSession{}, err
	}

	verified, err := r.userRepo.VerifyLogin(request.Username, request.Password, userData)

	if err != nil {
		return model.UserSession{}, err
	}

	if !verified {
		return model.UserSession{}, errors.New("can't verify login")
	}

	userSession, err := r.userRepo.CreateUserSession(userData.ID)

	if err != nil {
		return model.UserSession{}, err
	}

	return userSession, nil
}
