package mysql

import (
	"be-online-store/domain"
	"context"
	"errors"

	"database/sql"

	log "github.com/sirupsen/logrus"
)

type mysqlOrderRepository struct {
	Conn *sql.DB
}

// NewMySQLOrderRepository is constructor of MySQL repository
func NewMySQLOrderRepository(Conn *sql.DB) domain.OrderMySQLRepository {
	return &mysqlOrderRepository{Conn}
}

func (db *mysqlOrderRepository) InsertOrder(ctx context.Context, req domain.OrderRequest) (err error) {
	query := `INSERT INTO order (order_id, cart_id, total_price, status, dtm_crt, dtm_upd) VALUES (?, ?, ?, NOW(), NOW())`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	_, err = stmt.ExecContext(ctx, req.OrderID, req.CartID, req.TotalPrice, req.Status)
	if err != nil {
		err = errors.New("failed to create order")
		log.Error(err)
		return
	}

	return
}
