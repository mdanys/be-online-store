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
