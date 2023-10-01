package resto

import (
	"app/src/model"
	"app/src/model/constant"
	"app/src/repository/menu"
	"app/src/repository/order"
	"app/src/repository/user"
	"app/src/tracing"
	"context"
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

func (r *restoUseCase) GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetMenuList")
	defer span.End()
	return r.menuRepo.GetMenuList(ctx, menuType)
}

func (r *restoUseCase) Order(ctx context.Context, request model.OrderMenuRequest) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "Order")
	defer span.End()
	productOrderData := make([]model.ProductOrder, len(request.OrderProducts))

	for i, orderProduct := range request.OrderProducts {
		menuData, err := r.menuRepo.GetMenu(ctx, orderProduct.OrderCode)
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
		UserID:        request.UserID,
		Status:        constant.OrderStatusProcessed,
		ProductOrders: productOrderData,
		ReferenceID:   request.ReferenceID,
	}

	createdOrder, err := r.orderRepo.CreateOrder(ctx, order)

	if err != nil {
		return model.Order{}, err
	}

	return createdOrder, nil

}

func (r *restoUseCase) GetOrderInfo(ctx context.Context, request model.GetOrderInfoRequest) (model.Order, error) {
	ctx, span := tracing.CreateSpan(ctx, "GetOrderInfo")
	defer span.End()
	orderData, err := r.orderRepo.GetOrderInfo(ctx, request.OrderID)

	if request.UserID != orderData.UserID {
		return model.Order{}, errors.New("unauthorized different user ID")
	}

	if err != nil {
		return orderData, err
	}

	return orderData, nil
}

func (r *restoUseCase) RegisterUser(ctx context.Context, request model.RegisterRequest) (model.User, error) {
	ctx, span := tracing.CreateSpan(ctx, "RegisterUser")
	defer span.End()

	userRegistered, err := r.userRepo.CheckRegistered(ctx, request.Username)
	if err != nil {
		return model.User{}, err
	}

	if userRegistered {
		return model.User{}, errors.New("user already registered")
	}

	userHash, err := r.userRepo.GenerateUserHash(ctx, request.Password)

	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		ID:       uuid.New().String(),
		Username: request.Username,
		Hash:     userHash,
	}

	userData, err := r.userRepo.RegisterUser(ctx, user)

	if err != nil {
		return model.User{}, err
	}

	return userData, nil

}

func (r *restoUseCase) Login(ctx context.Context, request model.LoginRequest) (model.UserSession, error) {
	ctx, span := tracing.CreateSpan(ctx, "Login")
	defer span.End()

	userData, err := r.userRepo.GetUserData(ctx, request.Username)

	if err != nil {
		return model.UserSession{}, err
	}

	verified, err := r.userRepo.VerifyLogin(ctx, request.Username, request.Password, userData)

	if err != nil {
		return model.UserSession{}, err
	}

	if !verified {
		return model.UserSession{}, errors.New("can't verify login")
	}

	userSession, err := r.userRepo.CreateUserSession(ctx, userData.ID)

	if err != nil {
		return model.UserSession{}, err
	}

	return userSession, nil
}

func (r *restoUseCase) CheckSession(ctx context.Context, data model.UserSession) (userID string, err error) {
	ctx, span := tracing.CreateSpan(ctx, "Check Session")
	defer span.End()

	userID, err = r.userRepo.CheckSession(ctx, data)
	if err != nil {
		return "", err
	}

	return userID, nil

}
