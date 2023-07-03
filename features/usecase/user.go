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

func (uu *userUsecase) GetUserLogin(ctx context.Context, req domain.LoginRequest) (user domain.UserResponse, err error) {
	data, err := uu.userMySQLRepo.SelectUserLogin(ctx, req)
	if err != nil {
		log.Error(err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(*data.Password), []byte(req.Password))
	if err != nil {
		err = errors.New("password not match")
		log.Error(err)
		return
	}

	s, err := middleware.GenerateToken(int(data.ID), *data.Role)
	if err != nil {
		log.Error(err)
		return
	}

	user = domain.UserResponse{
		ID:          data.ID,
		Email:       data.Email,
		Name:        data.Name,
		Role:        data.Role,
		Dob:         data.Dob,
		Gender:      data.Gender,
		Address:     data.Address,
		UserPicture: data.UserPicture,
		DtmCrt:      data.DtmCrt,
		DtmUpd:      data.DtmUpd,
		Token:       &s,
	}

	return
}

func (uu *userUsecase) CreateUser(ctx context.Context, req domain.UserRequest) (user domain.User, err error) {
	generate, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
	if err != nil {
		err = errors.New("cannot encrypt password")
		log.Error(err)
		return
	}
	*req.Password = string(generate)

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
	user, err := uu.userMySQLRepo.SelectAllUser(ctx, offset, limit)
	if err != nil {
		log.Error(err)
		return
	}

	for _, v := range user {
		res.Data = append(res.Data, domain.UserResponse{
			ID:          v.ID,
			Email:       v.Email,
			Name:        v.Name,
			Dob:         v.Dob,
			Gender:      v.Gender,
			Address:     v.Address,
			UserPicture: v.UserPicture,
			DtmCrt:      v.DtmCrt,
			DtmUpd:      v.DtmUpd,
		})
	}

	res.Metadata = domain.Metadata{
		TotalData: int64(len(res.Data)),
		TotalPage: (int64(len(res.Data)) + limit - 1) / limit,
		Page:      page,
		Limit:     limit,
	}

	return
}

func (uu *userUsecase) UpdateUser(ctx context.Context, id int64, req domain.UserRequest) (err error) {
	user, err := uu.userMySQLRepo.SelectUserByID(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	if req.Password != nil {
		generate, er := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if er != nil {
			err = errors.New("cannot encrypt password")
			log.Error(err)
			return
		}
		*req.Password = string(generate)
	} else if req.Password == nil {
		req.Password = user.Password
	}

	if req.Name == nil {
		req.Name = user.Name
	}

	if req.Dob == nil {
		req.Dob = user.Dob
	}

	if req.Gender == nil {
		req.Gender = user.Gender
	}

	if req.Address == nil {
		req.Address = user.Address
	}

	if req.UserPicture == nil {
		req.UserPicture = user.UserPicture
	}

	err = uu.userMySQLRepo.EditUser(ctx, id, req)
	if err != nil {
		log.Error(err)
		return
	}

	return
}
