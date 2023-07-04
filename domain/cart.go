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

type CartSQL struct {
	CartID         int64   `json:"cart_id"`
	UserName       string  `json:"user_name"`
	CategoryName   string  `json:"category_name"`
	ProductName    string  `json:"product_name"`
	ProductPrice   float64 `json:"product_price"`
	ProductPicture string  `json:"product_picture"`
	ProductQty     int64   `json:"product_qty"`
	CartQty        int64   `json:"cart_qty"`
}

type CartResponse struct {
	Metadata Metadata  `json:"metadata"`
	Data     []CartSQL `json:"data"`
}

// CartMySQLRepository is Cart repository in MySQL
type CartMySQLRepository interface {
	InsertCart(ctx context.Context, req CartRequest) (err error)
	SelectCartByUserID(ctx context.Context, offset, limit, userId int64) (cart []CartSQL, err error)
	CountCartByUserID(ctx context.Context, userId int64) (count int64, err error)
	RemoveCart(ctx context.Context, cartId, userId int64) (err error)
}

// CartUsecase is Cart usecase
type CartUsecase interface {
	CreateCart(ctx context.Context, req CartRequest) (err error)
	GetCartByUserID(ctx context.Context, page, limit, userId int64) (cart CartResponse, err error)
	DeleteCart(ctx context.Context, cartId, userId int64) (err error)
}
