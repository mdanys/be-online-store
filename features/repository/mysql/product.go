package mysql

import (
	"be-online-store/domain"
	"context"
	"errors"

	"database/sql"

	log "github.com/sirupsen/logrus"
)

type mysqlProductRepository struct {
	Conn *sql.DB
}

// NewMySQLProductRepository is constructor of MySQL repository
func NewMySQLProductRepository(Conn *sql.DB) domain.ProductMySQLRepository {
	return &mysqlProductRepository{Conn}
}

func (db *mysqlProductRepository) InsertProduct(ctx context.Context, req domain.ProductRequest) (id int64, err error) {
	query := `INSERT INTO product (category_id, name, price, qty, rating, detail, product_picture, dtm_crt, dtm_upd) VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	res, err := stmt.ExecContext(ctx, req.CategoryID, req.Name, req.Price, req.Qty, req.Rating, req.Detail, req.ProductPicture)
	if err != nil {
		err = errors.New("failed to create product")
		log.Error(err)
		return
	}

	id, err = res.LastInsertId()
	if err != nil {
		log.Error(err)
		return
	}

	return
}

func (db *mysqlProductRepository) SelectProductByID(ctx context.Context, id int64) (product domain.Product, err error) {
	query := `SELECT id, category_id, name, price, qty, rating, detail, product_picture, dtm_crt, dtm_upd FROM product WHERE id = ?`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	row := stmt.QueryRowContext(ctx, id)
	err = row.Scan(&product.ID, &product.CategoryID, &product.Name, &product.Price, &product.Qty, &product.Rating,
		&product.Detail, &product.ProductPicture, &product.DtmCrt, &product.DtmUpd)
	if err != nil {
		log.Error(err)
		return
	}

	return
}
