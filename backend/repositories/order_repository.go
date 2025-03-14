package repositories

import (
	"Zitank/models"

	"github.com/jackc/pgx/pgtype"
	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	DB *sqlx.DB
}

func NewOrderRepo(OR *sqlx.DB) *OrderRepo {
	return &OrderRepo{
		DB: OR,
	}
}

func (OR OrderRepo) GetOrders() ([]*models.Orders, error) {
	var orders []*models.Orders
	err := OR.DB.Select(&orders, `SELECT * FROM orders`)
	return orders, err
}

func (OR OrderRepo) GetOrdersInRange(roomID int, fromTo pgtype.Tsrange) ([]*models.Orders, error) {
	var orders []*models.Orders
	err := OR.DB.Select(&orders, `SELECT * FROM orders WHERE FromTo && $1 and roomid=$2`, fromTo, roomID)
	return orders, err
}

func (OR OrderRepo) GetOrder(id int) (*models.Orders, error) {
	var order *models.Orders
	err := OR.DB.Get(&order, `SELECT * FROM orders WHERE id=$1 LIMIT 1`, id)
	return order, err
}

func (OR OrderRepo) CreateOrder(order *models.Orders) (int, error) {
	var id int
	err := OR.DB.QueryRow(`INSERT INTO orders (userid, roomid, checkindate, checkoutdate, totalprice, state, fromto, note) VALUES ($1, $2, now(), now(), $3, $4, $5, $6)  returning id;`, order.UserID, order.RoomID, order.TotalPrice, order.State, order.FromTo, order.Note).Scan(&id)
	return id, err
}

func (OR OrderRepo) UpdateOrder(order *models.Orders) error {
	_, err := OR.DB.Exec(`UPDATE orders SET userid=$1, roomid=$2, totalprice=$3, state=$4, fromto=$5, note=$6 WHERE id=$7`, order.UserID, order.RoomID, order.TotalPrice, order.State, order.FromTo, order.Note, order.ID)
	return err
}

func (OR OrderRepo) DeleteOrder(id int) error {
	_, err := OR.DB.Exec(`DELETE FROM orders WHERE id=$1`, id)
	return err
}

func (OR OrderRepo) GetAllOrderFromUser(userID int) ([]*models.Orders, error) {
	var orders []*models.Orders
	err := OR.DB.Select(&orders, `SELECT * FROM orders WHERE userid=$1`, userID)
	return orders, err
}

func (OR OrderRepo) GetAllOrderFromRoom(roomID int) ([]*models.Orders, error) {
	var orders []*models.Orders
	err := OR.DB.Select(&orders, `SELECT * FROM orders WHERE roomid=$1`, roomID)
	return orders, err
}

// need edit
func (OR OrderRepo) CheckOrder(order *models.Orders) bool {
	return true
}
