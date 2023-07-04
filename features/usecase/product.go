package usecase

import (
	"be-online-store/domain"
	"context"

	log "github.com/sirupsen/logrus"
)

type productUsecase struct {
	productMySQLRepo domain.ProductMySQLRepository
}

func NewProductUsecase(productMySQLRepo domain.ProductMySQLRepository) domain.ProductUsecase {
	return &productUsecase{
		productMySQLRepo: productMySQLRepo,
	}
}

func (pu *productUsecase) CreateProduct(ctx context.Context, req domain.ProductRequest) (product domain.Product, err error) {
	id, err := pu.productMySQLRepo.InsertProduct(ctx, req)
	if err != nil {
		log.Error(err)
		return
	}

	product, err = pu.productMySQLRepo.SelectProductByID(ctx, id)
	if err != nil {
		log.Error(err)
		return
	}

	return
}
