package domain

import (
	"context"
	"time"
)

type Order struct {
	ID         int64     `json:"id"`
	OrderID    *string   `json:"order_id"`
	UserID     *int64    `json:"user_id"`
	CartID     *int64    `json:"cart_id"`
	TotalPrice *float64  `json:"total_price"`
	Status     *string   `json:"status"`
	DtmCrt     time.Time `json:"dtm_crt"`
	DtmUpd     time.Time `json:"dtm_upd"`
}

type OrderRequest struct {
	OrderID    *string  `json:"order_id"`
	UserID     *int64   `json:"user_id"`
	CartID     *int64   `json:"cart_id"`
	TotalPrice *float64 `json:"total_price"`
	Status     *string  `json:"status"`
}

type OrderReq struct {
	CartID []int64 `json:"cart_id" form:"cart_id"`
}

// OrderMySQLRepository is Order repository in MySQL
type OrderMySQLRepository interface {
	InsertOrder(ctx context.Context, req OrderRequest) (err error)
	EditOrderStatus(ctx context.Context, status string, id int64) (err error)
	SelectOrderByOrderID(ctx context.Context, orderId string, userId int64) (order []Order, err error)
	SelectOrderByUserID(ctx context.Context, userId int64) (order []Order, err error)
}

// OrderUsecase is Order usecase
type OrderUsecase interface {
	CreateOrder(ctx context.Context, userId int64, orderId ...int64) (link string, err error)
	UpdateOrderStatus(ctx context.Context, orderId string, userId int64) (err error)
	GetOrderByOrderID(ctx context.Context, orderId string, userId int64) (order []Order, err error)
	GetOrderByUserID(ctx context.Context, userId int64) (order []Order, err error)
}
