package models

import (
	"github.com/jackc/pgtype"
)

type Posts struct {
	ID         uint             `db:"ID" json:"ID"`
	Title      pgtype.Varchar   `db:"Title" json:"Title"`
	Content    pgtype.Varchar   `db:"Content" json:"Content"`
	CreateDate pgtype.Timestamp `db:"CreateDate" json:"CreateDate"`
	EditDate   pgtype.Timestamp `db:"EditDate" json:"EditDate"`
	UserID     uint             `db:"UserID" json:"UserID"`
}
