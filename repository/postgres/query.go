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

func (repo *Repository) Index() (payments []*paymentPB.Payment, err error) {
	var id, orderId int32
	var paymentType, picture, status string

	query := "SELECT * FROM payments"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &orderId, &paymentType, &picture, &status)
		if err != nil {
			return nil, err
		}
		payment := &paymentPB.Payment{
			Id:      id,
			OrderId: orderId,
			Type:    paymentType,
			Picture: picture,
			Status:  status,
		}
		payments = append(payments, payment)
	}

	return payments, err
}

func (repo *Repository) IndexByUserID(user *paymentPB.User) (payments []*paymentPB.Payment, err error) {
	var id, orderId int32
	var paymentType, picture, status string

	query := fmt.Sprintf("SELECT * FROM payments WHERE user_id = %d", user.Id)
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &orderId, &paymentType, &picture, &status)
		if err != nil {
			return nil, err
		}
		payment := &paymentPB.Payment{
			Id:      id,
			OrderId: orderId,
			Type:    paymentType,
			Picture: picture,
			Status:  status,
		}
		payments = append(payments, payment)
	}

	return payments, err
}

func (repo *Repository) Show(payment *paymentPB.Payment) (*paymentPB.Payment, error) {
	var id, orderId int32
	var paymentType, picture, status string

	query := fmt.Sprintf("SELECT * FROM payments WHERE id=%d", payment.Id)
	err := repo.DB.QueryRow(query).Scan(&id, &orderId, &paymentType, &picture, &status)
	if err != nil {
		return nil, err
	}

	return &paymentPB.Payment{
		Id:      id,
		OrderId: orderId,
		Type:    paymentType,
		Picture: picture,
		Status:  status,
	}, err
}
