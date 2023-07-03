package mysql

import (
	"be-online-store/domain"
	"context"
	"errors"

	"database/sql"

	log "github.com/sirupsen/logrus"
)

type mysqlCategoryRepository struct {
	Conn *sql.DB
}

// NewMySQLCategoryRepository is constructor of MySQL repository
func NewMySQLCategoryRepository(Conn *sql.DB) domain.CategoryMySQLRepository {
	return &mysqlCategoryRepository{Conn}
}

func (db *mysqlCategoryRepository) InsertCategory(ctx context.Context, name string) (err error) {
	query := `INSERT INTO category (name, dtm_crt, dtm_upd) VALUES (?, NOW(), NOW())`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	_, err = stmt.ExecContext(ctx, name)
	if err != nil {
		err = errors.New("failed to create category")
		log.Error(err)
		return
	}

	return
}
