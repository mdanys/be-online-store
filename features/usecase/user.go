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
