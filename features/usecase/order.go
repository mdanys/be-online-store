package usecase

import (
	"be-online-store/domain"
	"be-online-store/utils/midtrans"
	"context"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type orderUsecase struct {
	orderMySQLRepo   domain.OrderMySQLRepository
	cartMySQLRepo    domain.CartMySQLRepository
	productMySQLRepo domain.ProductMySQLRepository
}

func NewOrderUsecase(orderMySQLRepo domain.OrderMySQLRepository, cartMySQLRepo domain.CartMySQLRepository, productMySQLRepo domain.ProductMySQLRepository) domain.OrderUsecase {
	return &orderUsecase{
		orderMySQLRepo:   orderMySQLRepo,
		cartMySQLRepo:    cartMySQLRepo,
		productMySQLRepo: productMySQLRepo,
	}
}

func (ou *orderUsecase) CreateOrder(ctx context.Context, userId int64, cartId ...int64) (link string, err error) {
	temp := uuid.New()
	orderId := "Order-" + temp.String()
	var grandTotal float64
	for _, v := range cartId {
		cart, er := ou.cartMySQLRepo.SelectCartByID(ctx, v)
		if er != nil {
			log.Error(er)
			return
		}

		totalPrice := *cart.ProductPrice * float64(*cart.CartQty)
		grandTotal += totalPrice
		status := "waiting"

		err = ou.orderMySQLRepo.InsertOrder(ctx, domain.OrderRequest{
			OrderID:    &orderId,
			UserID:     &userId,
			CartID:     &v,
			TotalPrice: &totalPrice,
			Status:     &status,
		})
		if err != nil {
			log.Error(err)
			return
		}
	}

	pay := midtrans.OrderMidtrans(orderId, int64(grandTotal))
	link = pay.RedirectURL

	return
}

func (ou *orderUsecase) UpdateOrderStatus(ctx context.Context, orderId string, userId int64) (err error) {
	check := midtrans.CheckMidtrans(orderId)
	status := check.TransactionStatus

	order, err := ou.orderMySQLRepo.SelectOrderByOrderID(ctx, orderId)
	if err != nil {
		log.Error(err)
		return
	}

	for _, v := range order {
		err = ou.orderMySQLRepo.EditOrderStatus(ctx, status, v.ID)
		if err != nil {
			log.Error(err)
			return
		}

		if status == "settlement" {
			cart, er := ou.cartMySQLRepo.SelectCartByID(ctx, *v.CartID)
			if er != nil {
				log.Error(er)
				return
			}

			product, er := ou.productMySQLRepo.SelectProductByID(ctx, *cart.ProductID)
			if er != nil {
				log.Error(er)
				return
			}

			qty := *product.Qty - *cart.CartQty

			err = ou.productMySQLRepo.EditQty(ctx, *cart.ProductID, qty)
			if err != nil {
				log.Error(err)
				return
			}

			err = ou.cartMySQLRepo.RemoveCart(ctx, *cart.CartID, userId)
			if err != nil {
				log.Error(err)
				return
			}
		}
	}

	return
}
