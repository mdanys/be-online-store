package domain

import (
	"context"
	"time"
)

type Category struct {
	ID     int64     `json:"id"`
	Name   string    `json:"name"`
	DtmCrt time.Time `json:"dtm_crt"`
	DtmUpd time.Time `json:"dtm_upd"`
}

type CategoryRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

// CategoryMySQLRepository is Category repository in MySQL
type CategoryMySQLRepository interface {
	InsertCategory(ctx context.Context, name string) (err error)
	SelectAllCategory(ctx context.Context) (category []Category, err error)
}

// CategoryUsecase is Category usecase
type CategoryUsecase interface {
	CreateCategory(ctx context.Context, name string) (err error)
	GetAllCategory(ctx context.Context) (category []Category, err error)
}
