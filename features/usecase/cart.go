package usecase

import (
	"be-online-store/domain"
	"context"
	"errors"

	log "github.com/sirupsen/logrus"
)

type cartUsecase struct {
	cartMySQLRepo    domain.CartMySQLRepository
	productMySQLRepo domain.ProductMySQLRepository
}

func NewCartUsecase(cartMySQLRepo domain.CartMySQLRepository, productMySQLRepo domain.ProductMySQLRepository) domain.CartUsecase {
	return &cartUsecase{
		cartMySQLRepo:    cartMySQLRepo,
		productMySQLRepo: productMySQLRepo,
	}
}

func (cu *cartUsecase) CreateCart(ctx context.Context, req domain.CartRequest) (err error) {
	product, err := cu.productMySQLRepo.SelectProductByID(ctx, *req.ProductID)
	if err != nil {
		log.Error(err)
		return
	}

	if *req.Qty > *product.Qty {
		err = errors.New("not enough stock")
		return
	}

	err = cu.cartMySQLRepo.InsertCart(ctx, req)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (cu *cartUsecase) GetCartByUserID(ctx context.Context, page, limit, userId int64) (cart domain.CartResponse, err error) {
	offset := limit * (page - 1)
	c, err := cu.cartMySQLRepo.SelectCartByUserID(ctx, offset, limit, userId)
	if err != nil {
		log.Error(err)
		return
	}

	total, err := cu.cartMySQLRepo.CountCartByUserID(ctx, userId)
	if err != nil {
		log.Error(err)
		return
	}

	for _, v := range c {
		cart.Data = append(cart.Data, domain.CartSQL{
			CartID:         v.CartID,
			UserName:       v.UserName,
			CategoryName:   v.CategoryName,
			ProductName:    v.ProductName,
			ProductPrice:   v.ProductPrice,
			ProductPicture: v.ProductPicture,
			ProductQty:     v.ProductQty,
			CartQty:        v.CartQty,
		})
	}

	cart.Metadata = domain.Metadata{
		TotalData: total,
		TotalPage: (total + limit - 1) / limit,
		Page:      page,
		Limit:     limit,
	}

	return
}

func (cu *cartUsecase) DeleteCart(ctx context.Context, cartId, userId int64) (err error) {
	err = cu.cartMySQLRepo.RemoveCart(ctx, cartId, userId)
	if err != nil {
		log.Error(err)
		return
	}

	return
}
