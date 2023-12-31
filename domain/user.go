package domain

import (
	"context"
	"time"
)

type User struct {
	ID          int64     `json:"id"`
	Email       *string   `json:"email"`
	Password    *string   `json:"password,omitempty"`
	Name        *string   `json:"name"`
	Role        *string   `json:"role,omitempty"`
	Dob         *string   `json:"dob"`
	Gender      *string   `json:"gender"`
	Address     *string   `json:"address"`
	UserPicture *string   `json:"user_picture"`
	DtmCrt      time.Time `json:"dtm_crt"`
	DtmUpd      time.Time `json:"dtm_upd"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UserResponse struct {
	ID          int64     `json:"id"`
	Email       *string   `json:"email,omitempty"`
	Name        *string   `json:"name,omitempty"`
	Role        *string   `json:"role,omitempty"`
	Dob         *string   `json:"dob,omitempty"`
	Gender      *string   `json:"gender,omitempty"`
	Address     *string   `json:"address,omitempty"`
	UserPicture *string   `json:"user_picture,omitempty"`
	DtmCrt      time.Time `json:"dtm_crt"`
	DtmUpd      time.Time `json:"dtm_upd"`
	Token       *string   `json:"token,omitempty"`
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

type GetAllResponse struct {
	Metadata Metadata       `json:"metadata"`
	Data     []UserResponse `json:"data"`
}

type Metadata struct {
	TotalData int64 `json:"total_data"`
	TotalPage int64 `json:"total_page"`
	Page      int64 `json:"page"`
	Limit     int64 `json:"limit"`
}

// UserMySQLRepository is User repository in MySQL
type UserMySQLRepository interface {
	SelectUserLogin(ctx context.Context, req LoginRequest) (user User, err error)
	InsertUser(ctx context.Context, req UserRequest) (id int64, err error)
	SelectUserByID(ctx context.Context, id int64) (user User, err error)
	SelectAllUser(ctx context.Context, offset, limit int64) (user []User, err error)
	EditUser(ctx context.Context, id int64, req UserRequest) (err error)
}

// UserUsecase is User usecase
type UserUsecase interface {
	GetUserLogin(ctx context.Context, req LoginRequest) (user UserResponse, err error)
	CreateUser(ctx context.Context, req UserRequest) (user User, err error)
	GetUserByID(ctx context.Context, id int64) (user User, err error)
	GetAllUser(ctx context.Context, page, limit int64) (res GetAllResponse, err error)
	UpdateUser(ctx context.Context, id int64, req UserRequest) (err error)
}
