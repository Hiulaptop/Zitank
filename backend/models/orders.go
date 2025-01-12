package models

import (
	"github.com/jackc/pgtype"
)

type Orders struct {
	ID           uint             `db:"ID"`
	FromTo       pgtype.Tsrange   `db:"FromTo"`
	State        pgtype.Varchar   `db:"State"`
	Note         pgtype.Text      `db:"Note"`
	CheckInDate  pgtype.Timestamp `db:"CheckInDate"`
	CheckOutDate pgtype.Timestamp `db:"CheckOutDate"`
	TotalPrice   float64          `db:"TotalPrice"`
	CreateDate   pgtype.Timestamp `db:"CreateDate"`
	EditDate     pgtype.Timestamp `db:"EditDate"`
	UserID       uint             `db:"UserID"`
	RoomID       uint             `db:"RoomID"`
}
