package auth

import (
	"database/sql"
	"errors"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

// Store ...
type Store interface {
	Get(*Query) (*Auth, error)
	Create(*Auth) error
	List() ([]*Auth, error)
	Update(*Auth) error
	Delete(a *Auth) error
}

// NewPostgres ...
func NewPostgres(dsn string) (Store, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &AuthStore{
		Store: db,
	}, nil
}

// AuthStore ...
type AuthStore struct {
	Store *sqlx.DB
}

// Get ...
func (as *AuthStore) Get(q *Query) (*Auth, error) {
	query := squirrel.Select("*").From("auth").Where("deleted_at is null")

	if q.Email == "" && q.Token == "" && q.UserID == "" {
		return nil, errors.New("must proovide a query")
	}

	if q.Email != "" {
		query = query.Where("email = ?", q.Token)
	}

	if q.Token != "" {
		query = query.Where("token = ?", q.Token)
	}

	if q.UserID != "" {
		query = query.Where("user_id = ?", q.UserID)
	}

	sql, args, err := query.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	row := as.Store.QueryRowx(sql, args...)

	c := &Auth{}
	if err := row.StructScan(c); err != nil {
		return nil, err
	}

	return c, nil
}

// Create ...
func (as *AuthStore) Create(a *Auth) error {
	sql, args, err := squirrel.
		Insert("auth").
		Columns("user_id", "token", "blacklist", "kind").
		Values(a.UserID, a.Token, a.Blacklist, a.Kind).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return err
	}

	row := as.Store.QueryRowx(sql, args...)
	if err := row.StructScan(a); err != nil {
		return err
	}

	return nil
}

// List ...
func (as *AuthStore) List() ([]*Auth, error) {
	query := squirrel.Select("*").From("auth").Where("deleted_at is null")

	sql, args, err := query.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := as.Store.Queryx(sql, args...)
	if err != nil {
		return nil, err
	}

	aa := make([]*Auth, 0)

	for rows.Next() {
		a := &Auth{}
		if err := rows.StructScan(a); err != nil {
			return nil, err
		}

		aa = append(aa, a)
	}

	return aa, nil
}

// Update ...
func (as *AuthStore) Update(a *Auth) error {
	sql, args, err := squirrel.Update("auth").Set("blacklist", a.Blacklist).Where("id = ?", a.ID).Suffix("returning *").PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return err
	}

	row := as.Store.QueryRowx(sql, args...)
	return row.StructScan(a)
}

// Delete ...
func (as *AuthStore) Delete(a *Auth) error {
	row := as.Store.QueryRowx("update auth set deleted_at = $1 where id = $2 returning *", time.Now(), a.ID)

	if err := row.StructScan(a); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	return nil
}
