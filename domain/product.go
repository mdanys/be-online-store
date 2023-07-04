package domain

import (
	"context"
	"time"
)

type Product struct {
	ID             *int64    `json:"id,omitempty"`
	CategoryID     *int64    `json:"category_id,omitempty"`
	Name           *string   `json:"name,omitempty"`
	Price          *float64  `json:"price,omitempty"`
	Qty            *int64    `json:"qty,omitempty"`
	Rating         *float32  `json:"rating,omitempty"`
	Detail         *string   `json:"detail,omitempty"`
	ProductPicture *string   `json:"product_picture,omitempty"`
	DtmCrt         time.Time `json:"dtm_crt"`
	DtmUpd         time.Time `json:"dtm_upd"`
}

type ProductSQL struct {
	ProductID      *int64    `json:"product_id,omitempty"`
	CategoryName   *string   `json:"category_name,omitempty"`
	ProductName    *string   `json:"product_name,omitempty"`
	ProductPrice   *float64  `json:"product_price,omitempty"`
	ProductQty     *int64    `json:"product_qty,omitempty"`
	ProductRating  *float32  `json:"product_rating,omitempty"`
	ProductDetail  *string   `json:"product_detail,omitempty"`
	ProductPicture *string   `json:"product_picture,omitempty"`
	DtmCrt         time.Time `json:"dtm_crt"`
	DtmUpd         time.Time `json:"dtm_upd"`
}

type ProductRequest struct {
	CategoryID     *int64   `json:"category_id" form:"category_id" validate:"required"`
	Name           *string  `json:"name" form:"name" validate:"required"`
	Price          *float64 `json:"price" form:"price" validate:"required"`
	Qty            *int64   `json:"qty" form:"qty" validate:"required"`
	Rating         *float32 `json:"rating" form:"rating"`
	Detail         *string  `json:"detail" form:"detail"`
	ProductPicture *string  `json:"product_picture" form:"product_picture"`
}

// ProductMySQLRepository is Product repository in MySQL
type ProductMySQLRepository interface {
	InsertProduct(ctx context.Context, req ProductRequest) (id int64, err error)
	SelectProductByID(ctx context.Context, id int64) (product Product, err error)
	SelectListProduct(ctx context.Context, offset, limit, categoryId int64) (product []ProductSQL, err error)
	EditQty(ctx context.Context, id, qty int64) (err error)
}

// ProductUsecase is Product usecase
type ProductUsecase interface {
	CreateProduct(ctx context.Context, req ProductRequest) (product Product, err error)
	GetListProduct(ctx context.Context, page, limit, categoryId int64) (product []ProductSQL, err error)
}
