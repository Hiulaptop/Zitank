package models

import "github.com/jmoiron/sqlx"

type PostgresStore struct {
	DB *sqlx.DB
}

type AppResource struct {
	Store *PostgresStore
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

func NewAppResource(store *PostgresStore) *AppResource {
	return &AppResource{
		Store: store,
	}
}
