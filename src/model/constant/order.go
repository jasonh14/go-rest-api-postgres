package constant

import "app/src/model"

const (
	OrderStatusProcessed model.OrderStatus = "processed"
	OrderStatusFinished  model.OrderStatus = "finished"
	OrderStatusFailed    model.OrderStatus = "failed"
)

const (
	ProductOrderStatusPreparing = "preparing"
	ProductOrderStatusFinished  = "finished"
)
