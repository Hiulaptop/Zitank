package controller

import (
	"Zitank/models"

	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

func (UserController) GetUsers(db *models.PostgresStore) ([]models.Users, error) {
	var user []models.Users
	err := db.DB.Get(&user, "SELECT * FROM users")
	return user, err
}

func (UserController) GetUser(db *models.PostgresStore, id int) (models.Users, error) {
	var user models.Users
	err := db.DB.Get(&user, "SELECT * FROM users WHERE ID=?", id)
	return user, err
}

func (UserController) GetUserByUsername(db *models.PostgresStore, username string) (models.Users, error) {
	var user models.Users
	err := db.DB.Get(&user, "SELECT * FROM users WHERE Username=?", username)
	return user, err
}

func (UserController) GetUserByEmail(db *models.PostgresStore, email string) (models.Users, error) {
	var user models.Users
	err := db.DB.Get(&user, "SELECT * FROM users WHERE Email=?", email)
	return user, err
}

func (UserController) CreateUser(db *models.PostgresStore, user models.Users) error {
	_, err := db.DB.Exec("INSERT INTO users (Username, Password, Fullname, Email, PhoneNumber, Gender, Role) VALUES (?, ?, ?, ?, ?, ?, ?)", user.Username, user.Password, user.Fullname, user.Email, user.PhoneNumber, user.Gender, user.Role)
	return err
}

func (UserController) UpdateUser(db *models.PostgresStore, user models.Users) error {
	_, err := db.DB.Exec("UPDATE users SET Fullname=?, PhoneNumber=?, Gender=?, WHERE id=?", user.Fullname, user.PhoneNumber, user.Gender, user.ID)
	return err
}

func (UserController) ResetPassword(db *models.PostgresStore, newPassword string, id int) error {
	_, err := db.DB.Exec("UPDATE users SET Password=? WHERE id=?", newPassword, id)
	return err
}

func (UserController) DeleteUser(db *models.PostgresStore, id int) error {
	_, err := db.DB.Exec("DELETE FROM users WHERE ID=?", id)
	return err
}

func (UC UserController) RegisterUser(db *models.PostgresStore, user models.Users) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password.String), 11)
	if err != nil {
		return err
	}
	user.Password.Set(string(hashedPassword))
	err = UC.CreateUser(db, user)
	return err
}

func (UC UserController) LoginUserByUsername(db *models.PostgresStore, username string, password string) (models.Users, error) {
	user, err := UC.GetUserByUsername(db, username)
	if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(password))
	return user, err
}

func (UC UserController) RoleCheck(db *models.PostgresStore, id int) bool {
	user, err := UC.GetUser(db, id)
	if err != nil {
		return false
	}
	if user.Role.String == "admin" {
		return true
	}
	return false
}
