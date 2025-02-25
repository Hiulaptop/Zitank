package models

import (
	"github.com/jackc/pgtype"
)

type Users struct {
	ID          uint             `db:"id" json:"id"`
	Username    string           `db:"username" json:"username"`
	Password    string           `db:"password" json:"password"`
	Fullname    string           `db:"fullname" json:"fullname"`
	Email       string           `db:"email" json:"email"`
	PhoneNumber string           `db:"phonenumber" json:"phonenumber"`
	Gender      string           `db:"gender" json:"gender"`
	Role        string           `db:"role" json:"role"`
	CreateDate  pgtype.Timestamp `db:"createdate" json:"createdate"`
}

type UserRepository interface {
	GetUsers() ([]*Users, error)
	GetUser(int) (*Users, error)
	GetUserByUsername(string) (*Users, error)
	GetUserByEmail(string) (*Users, error)

	CreateUser(*Users) (uint, error)
	RegisterUser(*Users) (uint, error)

	LoginUserByUsername(string, string) (*Users, error)

	RoleCheck(int) string

	UpdateUser(*Users) error

	ResetPassword(string, int) error

	DeleteUser(int) error
}
