package mysql

import (
	"be-online-store/domain"
	"context"
	"errors"

	"database/sql"

	log "github.com/sirupsen/logrus"
)

type mysqlCartRepository struct {
	Conn *sql.DB
}

// NewMySQLCartRepository is constructor of MySQL repository
func NewMySQLCartRepository(Conn *sql.DB) domain.CartMySQLRepository {
	return &mysqlCartRepository{Conn}
}

func (db *mysqlCartRepository) InsertCart(ctx context.Context, req domain.CartRequest) (err error) {
	query := `INSERT INTO cart (user_id, product_id, qty, dtm_upd) VALUES (?, ?, ?, NOW())`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	_, err = stmt.ExecContext(ctx, req.UserID, req.ProductID, req.Qty)
	if err != nil {
		err = errors.New("failed to insert cart")
		log.Error(err)
		return
	}

	return
}
