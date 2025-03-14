package repositories

import (
	"Zitank/models"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	DB *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (UR UserRepo) GetUsers() ([]*models.Users, error) {
	var user []*models.Users
	err := UR.DB.Select(&user, `SELECT * FROM users`)
	return user, err
}

func (UR UserRepo) GetUser(id int) (*models.Users, error) {
	var user models.Users
	err := UR.DB.Get(&user, `SELECT * FROM users WHERE ID=$1`, id)
	return &user, err
}

func (UR UserRepo) GetUserByUsername(username string) (*models.Users, error) {
	var user models.Users
	err := UR.DB.Get(&user, `SELECT * FROM users WHERE Username=$1 Limit 1`, username)
	return &user, err
}

func (UR UserRepo) GetUserByEmail(email string) (*models.Users, error) {
	var user models.Users
	err := UR.DB.Get(&user, `SELECT * FROM users WHERE Email=$1`, email)
	return &user, err
}

func (UR UserRepo) CreateUser(user *models.Users) (int, error) {
	var id int
	err := UR.DB.QueryRow(`INSERT INTO users (Username, Password, Fullname, Email, PhoneNumber, Gender, Role) VALUES ($1, $2, $3, $4, $5, $6, $7)`, user.Username, user.Password, user.Fullname, user.Email, user.PhoneNumber, user.Gender, user.Role).Scan(&id)
	return id, err
}

func (UR UserRepo) UpdateUser(user *models.Users) error {
	_, err := UR.DB.Exec(`UPDATE users SET Fullname=$1, PhoneNumber=$2, Gender=$3, WHERE id=$4`, user.Fullname, user.PhoneNumber, user.Gender, user.ID)
	return err
}

func (UR UserRepo) ResetPassword(newPassword string, id int) error {
	_, err := UR.DB.Exec(`UPDATE users SET Password=$1 WHERE id=$2`, newPassword, id)
	return err
}

func (UR UserRepo) DeleteUser(id int) error {
	_, err := UR.DB.Exec(`DELETE FROM users WHERE ID=$1`, id)
	return err
}

func (UR UserRepo) RegisterUser(user *models.Users) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 11)
	if err != nil {
		return 0, err
	}
	user.Password = (string(hashedPassword))
	id, err := UR.CreateUser(user)
	return id, err
}

func (UR UserRepo) LoginUserByUsername(username string, password string) (*models.Users, error) {
	user, err := UR.GetUserByUsername(username)
	if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return user, err
}

func (UR UserRepo) RoleCheck(id int) string {
	user, err := UR.GetUser(id)
	if err != nil {
		return "err"
	}
	return user.Role
}
