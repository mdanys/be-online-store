package domain

import (
	"context"
	"time"
)

type Cart struct {
	ID        int64     `json:"id"`
	UserID    *int64    `json:"user_id"`
	ProductID *int64    `json:"product_id"`
	Qty       *int64    `json:"qty"`
	DtmUpd    time.Time `json:"dtm_upd"`
}

type CartRequest struct {
	UserID    *int64 `json:"user_id" form:"user_id"`
	ProductID *int64 `json:"product_id" form:"product_id" validate:"required"`
	Qty       *int64 `json:"qty" form:"qty" validate:"required"`
}

// CartMySQLRepository is Cart repository in MySQL
type CartMySQLRepository interface {
	InsertCart(ctx context.Context, req CartRequest) (err error)
}

// CartUsecase is Cart usecase
type CartUsecase interface {
	CreateCart(ctx context.Context, req CartRequest) (err error)
}
