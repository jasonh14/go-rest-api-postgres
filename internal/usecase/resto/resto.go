package resto

import (
	"app/internal/model"
	"app/internal/model/constant"
	"app/internal/repository/menu"
	"app/internal/repository/order"

	"github.com/google/uuid"
)

type restoUseCase struct {
	menuRepo  menu.Repository
	orderRepo order.Repository
}

func GetuseCase(menuRepo menu.Repository, orderRepo order.Repository) Usecase {
	return &restoUseCase{
		menuRepo:  menuRepo,
		orderRepo: orderRepo,
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
