package models

import (
	"github.com/jackc/pgtype"
)

type Users struct {
	ID          uint             `db:"ID"`
	Username    pgtype.Varchar   `db:"Username"`
	Password    pgtype.Varchar   `db:"Password"`
	Fullname    pgtype.Varchar   `db:"Fullname"`
	Email       pgtype.Varchar   `db:"Email"`
	PhoneNumber pgtype.Varchar   `db:"PhoneNumber"`
	Gender      pgtype.Varchar   `db:"Gender"`
	Role        pgtype.Varchar   `db:"Role"`
	CreateDate  pgtype.Timestamp `db:"CreateDate"`
}
