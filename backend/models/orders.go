package models

import (
	"github.com/jackc/pgx/pgtype"
)

type Orders struct {
	ID           uint             `db:"id" json:"id"`
	FromTo       pgtype.Tsrange   `db:"fromto" json:"fromto"`
	State        string           `db:"state" json:"state"`
	Note         string           `db:"note" json:"note"`
	CheckInDate  pgtype.Timestamp `db:"checkindate" json:"checkindate"`
	CheckOutDate pgtype.Timestamp `db:"checkoutdate" json:"checkoutdate"`
	TotalPrice   float64          `db:"totalprice" json:"totalprice"`
	CreateDate   pgtype.Timestamp `db:"createdate" json:"createdate"`
	EditDate     pgtype.Timestamp `db:"editdate" json:"editdate"`
	UserID       uint             `db:"userid" json:"userid"`
	RoomID       uint             `db:"roomid" json:"roomid"`
}

type OrderRepository interface {
	GetOrders() ([]*Orders, error)
	GetOrder(int) (*Orders, error)
	GetAllOrderFromUser(int) ([]*Orders, error)
	GetAllOrderFromRoom(int) ([]*Orders, error)
	GetOrdersInRange(int, pgtype.Tsrange) ([]*Orders, error)

	CreateOrder(*Orders) (int, error)

	UpdateOrder(*Orders) error

	DeleteOrder(int) error

	CheckOrder(*Orders) bool
}
