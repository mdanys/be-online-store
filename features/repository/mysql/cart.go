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

func (db *mysqlCartRepository) SelectCartByUserID(ctx context.Context, offset, limit, userId int64) (cart []domain.CartSQL, err error) {
	query := `SELECT c.id, c.user_id, u.name, ca.name, c.product_id, p.name, p.price, p.product_picture, p.qty, c.qty FROM cart c
	INNER JOIN user u ON u.id = c.user_id INNER JOIN product p ON p.id = c.product_id
	INNER JOIN category ca ON ca.id = p.category_id WHERE c.user_id = ? LIMIT ? OFFSET ?`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	rows, err := stmt.QueryContext(ctx, userId, limit, offset)
	if err != nil {
		log.Error(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var i domain.CartSQL
		err = rows.Scan(&i.CartID, &i.UserID, &i.UserName, &i.CategoryName, &i.ProductID, &i.ProductName, &i.ProductPrice, &i.ProductPicture,
			&i.ProductQty, &i.CartQty)

		cart = append(cart, i)
	}

	return
}

func (db *mysqlCartRepository) CountCartByUserID(ctx context.Context, userId int64) (count int64, err error) {
	query := `SELECT COUNT(id) FROM cart WHERE user_id = ?`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	row := stmt.QueryRowContext(ctx, userId)
	err = row.Scan(&count)
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (db *mysqlCartRepository) RemoveCart(ctx context.Context, cartId, userId int64) (err error) {
	query := `DELETE FROM cart WHERE id = ? AND user_id = ?`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	res, err := stmt.ExecContext(ctx, cartId, userId)
	if err != nil {
		log.Error(err)
		return
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
		return
	}

	if affected == 0 {
		err = errors.New("no data has been deleted")
		return
	}

	return
}

func (db *mysqlCartRepository) SelectCartByID(ctx context.Context, id int64) (cart domain.CartSQL, err error) {
	query := `SELECT c.id, c.user_id, u.name, ca.name, c.product_id, p.name, p.price, p.product_picture, p.qty, c.qty FROM cart c
	INNER JOIN user u ON u.id = c.user_id INNER JOIN product p ON p.id = c.product_id
	INNER JOIN category ca ON ca.id = p.category_id WHERE c.id = ?`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	row := stmt.QueryRowContext(ctx, id)
	err = row.Scan(&cart.CartID, &cart.UserID, &cart.UserName, &cart.CategoryName, &cart.ProductID, &cart.ProductName,
		&cart.ProductPrice, &cart.ProductPicture, &cart.ProductQty, &cart.CartQty)
	if err != nil {
		log.Error(err)
		return
	}

	return
}
