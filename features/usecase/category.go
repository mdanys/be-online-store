package usecase

import (
	"be-online-store/domain"
	"context"

	log "github.com/sirupsen/logrus"
)

type categoryUsecase struct {
	categoryMySQLRepo domain.CategoryMySQLRepository
}

func NewCategoryUsecase(categoryMySQLRepo domain.CategoryMySQLRepository) domain.CategoryUsecase {
	return &categoryUsecase{
		categoryMySQLRepo: categoryMySQLRepo,
	}
}

func (cu *categoryUsecase) CreateCategory(ctx context.Context, name string) (err error) {
	err = cu.categoryMySQLRepo.InsertCategory(ctx, name)
	if err != nil {
		log.Error(err)
		return
	}

	return
}
