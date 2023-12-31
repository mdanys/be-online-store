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

func (db *mysqlCategoryRepository) SelectAllCategory(ctx context.Context) (category []domain.Category, err error) {
	query := `SELECT id, name, dtm_crt, dtm_upd FROM category`
	log.Debug(query)

	rows, err := db.Conn.QueryContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var i domain.Category
		err = rows.Scan(&i.ID, &i.Name, &i.DtmCrt, &i.DtmUpd)
		if err != nil {
			log.Error(err)
			return
		}

		category = append(category, i)
	}

	return
}
