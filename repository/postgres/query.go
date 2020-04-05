// payment-service/repository/postgres/query.go

package postgres

import (
	"database/sql"
	"fmt"

	paymentPB "github.com/SleepingNext/payment-service/proto"
)

type Repository struct {
	DB *sql.DB
}

// Index will index all active payments
func (repo *Repository) Index() (payments []*paymentPB.Payment, err error) {
	var id, orderID int32
	var paymentType, picture, status string

	query := "SELECT * FROM payments WHERE status = 'active'"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &orderID, &paymentType, &picture, &status)
		if err != nil {
			return nil, err
		}
		payment := &paymentPB.Payment{
			Id:      id,
			OrderId: orderID,
			Type:    paymentType,
			Picture: picture,
			Status:  status,
		}
		payments = append(payments, payment)
	}

	return payments, err
}

// IndexByUserID will index all active payments by it's userID
func (repo *Repository) IndexByUserID(user *paymentPB.User) (payments []*paymentPB.Payment, err error) {
	var id, orderID int32
	var paymentType, picture, status string

	query := fmt.Sprintf("SELECT * FROM payments WHERE user_id = %d AND status = 'active'", user.Id)
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &orderID, &paymentType, &picture, &status)
		if err != nil {
			return nil, err
		}
		payment := &paymentPB.Payment{
			Id:      id,
			OrderId: orderID,
			Type:    paymentType,
			Picture: picture,
			Status:  status,
		}
		payments = append(payments, payment)
	}

	return payments, err
}

// Show will show an active payment by it's id
func (repo *Repository) Show(payment *paymentPB.Payment) (*paymentPB.Payment, error) {
	var id, orderID int32
	var paymentType, picture, status string

	query := fmt.Sprintf("SELECT * FROM payments WHERE id = %d AND status = 'active'", payment.Id)
	err := repo.DB.QueryRow(query).Scan(&id, &orderID, &paymentType, &picture, &status)
	if err != nil {
		return nil, err
	}

	return &paymentPB.Payment{
		Id:      id,
		OrderId: orderID,
		Type:    paymentType,
		Picture: picture,
		Status:  status,
	}, err
}
