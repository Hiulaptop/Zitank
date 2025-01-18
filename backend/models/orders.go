package models

import (
	"github.com/jackc/pgtype"
)

type Orders struct {
	ID           uint             `db:"ID" json:"ID"`
	FromTo       pgtype.Tsrange   `db:"FromTo" json:"FromTo"`
	State        pgtype.Varchar   `db:"State" json:"State"`
	Note         pgtype.Text      `db:"Note" json:"Note"`
	CheckInDate  pgtype.Timestamp `db:"CheckInDate" json:"CheckInDate"`
	CheckOutDate pgtype.Timestamp `db:"CheckOutDate" json:"CheckOutDate"`
	TotalPrice   float64          `db:"TotalPrice" json:"TotalPrice"`
	CreateDate   pgtype.Timestamp `db:"CreateDate" json:"CreateDate"`
	EditDate     pgtype.Timestamp `db:"EditDate" json:"EditDate"`
	UserID       uint             `db:"UserID" json:"UserID"`
	RoomID       uint             `db:"RoomID" json:"RoomID"`
}
