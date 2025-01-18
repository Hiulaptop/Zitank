package models

import (
	"github.com/jackc/pgtype"
)

type Rooms struct {
	ID          uint             `db:"ID" json:"ID"`
	Name        pgtype.Varchar   `db:"Name" json:"Name"`
	Address     pgtype.Varchar   `db:"Address" json:"Address"`
	Description pgtype.Text      `db:"Description" json:"Description"`
	Price       float64          `db:"Price" json:"Price"`
	CreateDate  pgtype.Timestamp `db:"CreateDate" json:"CreateDate"`
	EditDate    pgtype.Timestamp `db:"EditDate" json:"EditDate"`
	UserID      uint             `db:"UserID" json:"UserID"`
}
