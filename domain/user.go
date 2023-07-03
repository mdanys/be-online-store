package domain

import (
	"context"
	"time"
)

type User struct {
	ID          int64     `json:"id"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	Role        string    `json:"role"`
	Dob         string    `json:"dob"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	UserPicture string    `json:"user_picture"`
	DtmCrt      time.Time `json:"dtm_crt"`
	DtmUpd      time.Time `json:"dtm_upd"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UserLogin struct {
	ID          int64     `json:"id"`
	Email       string    `json:"email"`
	Name        string    `json:"name"`
	Role        string    `json:"role"`
	Dob         string    `json:"dob"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	UserPicture string    `json:"user_picture"`
	DtmCrt      time.Time `json:"dtm_crt"`
	DtmUpd      time.Time `json:"dtm_upd"`
	Token       string    `json:"token"`
}

type UserRequest struct {
	Email       string `json:"email" form:"email" validate:"required,email"`
	Password    string `json:"password" form:"password" validate:"required"`
	Name        string `json:"name" form:"name" validate:"required"`
	Dob         string `json:"dob" form:"dob"`
	Gender      string `json:"gender" form:"gender"`
	Address     string `json:"address" form:"address"`
	UserPicture string `json:"user_picture" form:"user_picture"`
}

// UserMySQLRepository is User repository in MySQL
type UserMySQLRepository interface {
	SelectUserLogin(ctx context.Context, req LoginRequest) (user User, err error)
	InsertUser(ctx context.Context, req UserRequest) (id int64, err error)
	SelectUserByID(ctx context.Context, id int64) (user User, err error)
}

// UserUsecase is User usecase
type UserUsecase interface {
	GetUserLogin(ctx context.Context, req LoginRequest) (user UserLogin, err error)
	CreateUser(ctx context.Context, req UserRequest) (user User, err error)
}
