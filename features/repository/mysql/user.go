package mysql

import (
	"be-online-store/domain"
	"context"
	"database/sql"
	"errors"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type mysqlUserRepository struct {
	Conn *sql.DB
}

// NewMySQLUserRepository is constructor of MySQL repository
func NewMySQLUserRepository(Conn *sql.DB) domain.UserMySQLRepository {
	return &mysqlUserRepository{Conn}
}

func (db *mysqlUserRepository) SelectUserLogin(ctx context.Context, req domain.LoginRequest) (user domain.User, err error) {
	query := `SELECT id, email, name, role, dob, gender, address, user_picture, dtm_crt, dtm_upd FROM user WHERE email = ?`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	row := stmt.QueryRowContext(ctx, req.Email)
	err = row.Scan(&user.ID, &user.Email, &user.Name, &user.Role, &user.Dob, &user.Gender, &user.Address, &user.UserPicture, &user.DtmCrt, &user.DtmUpd)
	if err != nil {
		err = errors.New("user not found")
		log.Error(err)
		return
	}

	return
}

func (db *mysqlUserRepository) InsertUser(ctx context.Context, req domain.UserRequest) (id int64, err error) {
	query := `INSERT INTO user (email, password, name, role, dob, gender, address, user_picture, dtm_crt, dtm_upd) VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	res, err := stmt.ExecContext(ctx, req.Email, req.Password, req.Name, "customer", req.Dob, req.Gender, req.Address, req.UserPicture)
	if err != nil {
		err = errors.New("failed to create user")
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

func (db *mysqlUserRepository) SelectUserByID(ctx context.Context, id int64) (user domain.User, err error) {
	query := `SELECT id, email, name, role, dob, gender, address, user_picture, dtm_crt, dtm_upd FROM user WHERE id = ?`
	log.Debug(query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}

	row := stmt.QueryRowContext(ctx, id)
	err = row.Scan(&user.ID, &user.Email, &user.Name, &user.Role, &user.Dob, &user.Gender, &user.Address, &user.UserPicture, &user.DtmCrt, &user.DtmUpd)
	if err != nil {
		err = errors.New("user not found")
		log.Error(err)
		return
	}

	return
}

func (db *mysqlUserRepository) SelectAllUser(ctx context.Context, offset, limit int64) (user []domain.User, err error) {
	query := `SELECT id, email, name, dob, gender, address, user_picture, dtm_crt, dtm_upd FROM user WHERE role = 'customer'`

	if limit > 0 {
		query += " LIMIT " + strconv.Itoa(int(limit))
	}

	if offset > 0 {
		query += " OFFSET " + strconv.Itoa(int(offset))
	}

	log.Debug(query)

	rows, err := db.Conn.QueryContext(ctx, query)
	if err != nil {
		log.Error(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var i domain.User
		err = rows.Scan(&i.ID, &i.Email, &i.Name, &i.Dob, &i.Gender, &i.Address, &i.UserPicture, &i.DtmCrt, &i.DtmUpd)
		if err != nil {
			log.Error(err)
			return
		}

		user = append(user, i)
	}

	return
}
