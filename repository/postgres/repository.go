// payment-service/repository/postgres/repository.go

package postgres

import (
	"fmt"

	paymentPB "github.com/SleepingNext/payment-service/proto"
)

// Store will store a new payment
func (repo *Repository) Store(payment *paymentPB.Payment) (*paymentPB.Payment, error) {
	query := fmt.Sprintf("INSERT INTO payments (order_id, type, picture, status)"+
		" VALUES ('%d', '%s', '%s', 'active')", payment.OrderId, payment.Type, payment.Picture)
	_, err := repo.DB.Exec(query)

	return payment, err
}

// Update will update an existing payment's data
func (repo *Repository) Update(payment *paymentPB.Payment) (*paymentPB.Payment, error) {
	query := fmt.Sprintf("UPDATE payments SET order_id = '%d', type = '%s', picture = %s, status = 'active'"+
		" WHERE id = %d", payment.OrderId, payment.Type, payment.Picture, payment.Id)
	_, err := repo.DB.Exec(query)

	return payment, err
}

// Destroy will update an existing payment's status to inactive
func (repo *Repository) Destroy(payment *paymentPB.Payment) (*paymentPB.Payment, error) {
	query := fmt.Sprintf("UPDATE payments SET status='inactive' WHERE id=%d", payment.Id)
	_, err := repo.DB.Exec(query)

	return payment, err
}
