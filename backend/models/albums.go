package models

import (
	"github.com/jackc/pgtype"
)

type Albums struct {
	ID          uint             `db:"ID"`
	Name        pgtype.Varchar   `db:"Name"`
	AuthorName  pgtype.Varchar   `db:"AuthorName"`
	ReleaseDate pgtype.Timestamp `db:"ReleaseDate"`
	Type        pgtype.Varchar   `db:"Type"`
	Description pgtype.Text      `db:"Description"`
	Link        pgtype.Varchar   `db:"Link"`
	CreateDate  pgtype.Timestamp `db:"CreateDate"`
	EditDate    pgtype.Timestamp `db:"EditDate"`
	UserID      uint             `db:"UserID"`
}
