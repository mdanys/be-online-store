package domain

import (
	"context"
	"time"
)

type Order struct {
	ID         int64     `json:"id"`
	CartID     *int64    `json:"cart_id"`
	TotalPrice *float64  `json:"total_price"`
	Status     *string   `json:"status"`
	DtmCrt     time.Time `json:"dtm_crt"`
	DtmUpd     time.Time `json:"dtm_upd"`
}

type OrderRequest struct {
	OrderID    *string  `json:"order_id"`
	CartID     *int64   `json:"cart_id"`
	TotalPrice *float64 `json:"total_price"`
	Status     *string  `json:"status"`
}

// OrderMySQLRepository is Order repository in MySQL
type OrderMySQLRepository interface {
	InsertOrder(ctx context.Context, req OrderRequest) (err error)
}

// OrderUsecase is Order usecase
type OrderUsecase interface {
	CreateOrder(ctx context.Context, orderId ...int64) (link string, err error)
}
