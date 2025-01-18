package models

import (
	"github.com/jackc/pgtype"
)

type Albums struct {
	ID          uint             `db:"ID" json:"ID"`
	Name        pgtype.Varchar   `db:"Name" json:"Name"`
	AuthorName  pgtype.Varchar   `db:"AuthorName" json:"AuthorName"`
	ReleaseDate pgtype.Timestamp `db:"ReleaseDate" json:"ReleaseDate"`
	Type        pgtype.Varchar   `db:"Type" json:"Type"`
	Description pgtype.Text      `db:"Description" json:"Description"`
	Link        pgtype.Varchar   `db:"Link" json:"Link"`
	CreateDate  pgtype.Timestamp `db:"CreateDate" json:"CreateDate"`
	EditDate    pgtype.Timestamp `db:"EditDate" json:"EditDate"`
	UserID      uint             `db:"UserID" json:"UserID"`
}
