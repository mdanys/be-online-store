package usecase

import (
	"be-online-store/domain"
	"be-online-store/utils/middleware"
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userMySQLRepo domain.UserMySQLRepository
}

func NewUserUsecase(userMySQLRepo domain.UserMySQLRepository) domain.UserUsecase {
	return &userUsecase{
		userMySQLRepo: userMySQLRepo,
	}
}

func (uu *userUsecase) GetUserLogin(ctx context.Context, req domain.LoginRequest) (user domain.UserLogin, err error) {
	data, err := uu.userMySQLRepo.SelectUserLogin(ctx, req)
	if err != nil {
		log.Error(err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password))
	if err != nil {
		err = errors.New("password not match")
		log.Error(err)
		return
	}

	s, err := middleware.GenerateToken(int(data.ID), data.Role)
	if err != nil {
		log.Error(err)
		return
	}

	user.Token = s
	user.ID = data.ID
	user.Email = data.Email
	user.Name = data.Name
	user.Role = data.Role
	user.Dob = data.Dob
	user.Gender = data.Gender
	user.Address = data.Address
	user.UserPicture = data.UserPicture
	user.DtmCrt = data.DtmCrt
	user.DtmUpd = data.DtmUpd

	return
}

func (uu *userUsecase) CreateUser(ctx context.Context, req domain.UserRequest) (user domain.User, err error) {
	generate, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		err = errors.New("cannot encrypt password")
		log.Error(err)
		return
	}
	req.Password = string(generate)

	id, err := uu.userMySQLRepo.InsertUser(ctx, req)
	if err != nil {
		log.Error(err)
		return
	}

	user, err = uu.userMySQLRepo.SelectUserByID(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (uu *userUsecase) GetUserByID(ctx context.Context, id int64) (user domain.User, err error) {
	user, err = uu.userMySQLRepo.SelectUserByID(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (uu *userUsecase) GetAllUser(ctx context.Context, page, limit int64) (res domain.GetAllResponse, err error) {
	offset := limit * (page - 1)
	res.Data, err = uu.userMySQLRepo.SelectAllUser(ctx, offset, limit)
	if err != nil {
		log.Error(err)
		return
	}

	res.Metadata = domain.Metadata{
		TotalData: int64(len(res.Data)),
		TotalPage: (int64(len(res.Data)) + limit - 1) / limit,
		Page:      page,
		Limit:     limit,
	}

	return
}
