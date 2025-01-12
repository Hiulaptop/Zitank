package models

import (
	"github.com/jackc/pgtype"
)

type Posts struct {
	ID         uint             `db:"ID"`
	Title      pgtype.Varchar   `db:"Title"`
	Content    pgtype.Varchar   `db:"Content"`
	CreateDate pgtype.Timestamp `db:"CreateDate"`
	EditDate   pgtype.Timestamp `db:"EditDate"`
	UserID     uint             `db:"UserID"`
}
