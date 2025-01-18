package models

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/jmoiron/sqlx"
)

type PostgresStore struct {
	DB *sqlx.DB
}

type AppResource struct {
	Store     *PostgresStore
	TokenAuth *jwtauth.JWTAuth
}

func (s *PostgresStore) Connect(DB_URL string) error {
	db, err := sqlx.Open("postgres", DB_URL)
	if err != nil {
		return err
	}
	s.DB = db
	return nil
}

func (s *PostgresStore) Close() error {
	return s.DB.Close()
}

func NewAppResource(store *PostgresStore, token *jwtauth.JWTAuth) *AppResource {
	return &AppResource{
		Store:     store,
		TokenAuth: token,
	}
}
