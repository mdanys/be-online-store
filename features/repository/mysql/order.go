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
	query := `INSERT INTO order (order_id, cart_id, total_price, status, dtm_crt, dtm_upd) VALUES (?, ?, ?, ?, NOW(), NOW())`
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

func (db *mysqlOrderRepository) EditOrderStatus(ctx context.Context, status string, id int64) (err error) {
	query := `UPDATE order SET status = ? WHERE id = ?`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	res, err := stmt.ExecContext(ctx, status, id)
	if err != nil {
		log.Error(err)
		return
	}

	affect, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
		return
	}

	if affect == 0 {
		err = errors.New("no data updated")
		return
	}

	return
}

func (db *mysqlOrderRepository) SelectOrderByOrderID(ctx context.Context, orderId string) (order []domain.Order, err error) {
	query := `SELECT id, order_id, cart_id, total_price, status, dtm_crt, dtm_upd FROM order WHERE order_id = ?`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	rows, err := stmt.QueryContext(ctx, orderId)
	if err != nil {
		log.Error(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var i domain.Order
		err = rows.Scan(&i.ID, &i.OrderID, &i.CartID, &i.TotalPrice, &i.Status, &i.DtmCrt, &i.DtmUpd)
		if err != nil {
			log.Error(err)
			return
		}

		order = append(order, i)
	}

	return
}
