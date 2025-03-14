package models

import (
	"Zitank/utils"

	"github.com/jackc/pgtype"
)

type Rooms struct {
	ID          uint             `db:"id" json:"id"`
	Name        string           `db:"name" json:"name"`
	Address     string           `db:"address" json:"address"`
	Description pgtype.Text      `db:"description" json:"description"`
	Price       float64          `db:"price" json:"price"`
	CreateDate  pgtype.Timestamp `db:"createdate" json:"createdate"`
	EditDate    pgtype.Timestamp `db:"editdate" json:"editdate"`
	UserID      uint             `db:"userid" json:"userid"`
}

type RoomObject struct {
	Name        string              `json:"name"`
	Address     string              `json:"address"`
	Description pgtype.Text         `json:"description"`
	Price       utils.StringToFloat `json:"price"`
}

type RoomRepository interface {
	GetRooms() ([]*Rooms, error)
	GetRoom(int) (*Rooms, error)

	CreateRoom(*RoomObject, int) (int, error)

	UpdateRoom(*Rooms) error

	DeleteRoom(int) error
}
