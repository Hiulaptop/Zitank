package models

import (
	"github.com/jackc/pgtype"
)

type Rooms struct {
	ID          uint             `db:"ID"`
	Name        pgtype.Varchar   `db:"Name"`
	Address     pgtype.Varchar   `db:"Address"`
	Description pgtype.Text      `db:"Description"`
	Price       float64          `db:"Price"`
	CreateDate  pgtype.Timestamp `db:"CreateDate"`
	EditDate    pgtype.Timestamp `db:"EditDate"`
	UserID      uint             `db:"UserID"`
}
