package usecase

import (
	"be-online-store/domain"
	"context"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type orderUsecase struct {
	orderMySQLRepo domain.OrderMySQLRepository
	cartMySQLRepo  domain.CartMySQLRepository
}

func NewOrderUsecase(orderMySQLRepo domain.OrderMySQLRepository, cartMySQLRepo domain.CartMySQLRepository) domain.OrderUsecase {
	return &orderUsecase{
		orderMySQLRepo: orderMySQLRepo,
		cartMySQLRepo:  cartMySQLRepo,
	}
}

func (cu *orderUsecase) CreateOrder(ctx context.Context, cartId ...int64) (link string, err error) {
	temp := uuid.New()
	orderId := "Order-" + temp.String()
	var grandTotal float64
	for _, v := range cartId {
		cart, er := cu.cartMySQLRepo.SelectCartByID(ctx, v)
		if er != nil {
			log.Error(er)
			return
		}

		totalPrice := cart.ProductPrice * float64(cart.CartQty)
		grandTotal += totalPrice
		status := "waiting"

		err = cu.orderMySQLRepo.InsertOrder(ctx, domain.OrderRequest{
			OrderID:    &orderId,
			CartID:     &v,
			TotalPrice: &totalPrice,
			Status:     &status,
		})
		if err != nil {
			log.Error(err)
			return
		}
	}

	return
}
