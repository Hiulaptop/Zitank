package models

import (
	"github.com/jackc/pgtype"
)

type Users struct {
	ID          uint             `db:"ID" json:"ID"`
	Username    pgtype.Varchar   `db:"Username" json:"Username"`
	Password    pgtype.Varchar   `db:"Password" json:"Password"`
	Fullname    pgtype.Varchar   `db:"Fullname" json:"Fullname"`
	Email       pgtype.Varchar   `db:"Email" json:"Email"`
	PhoneNumber pgtype.Varchar   `db:"PhoneNumber" json:"PhoneNumber"`
	Gender      pgtype.Varchar   `db:"Gender" json:"Gender"`
	Role        pgtype.Varchar   `db:"Role" json:"Role"`
	CreateDate  pgtype.Timestamp `db:"CreateDate" json:"CreateDate"`
}
