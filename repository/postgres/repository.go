// payment-service/repository/postgres/repository.go

package postgres

import (
	"fmt"

	paymentPB "github.com/SleepingNext/payment-service/proto"
)

func (repo *Repository) Store(payment *paymentPB.Payment) (*paymentPB.Payment, error) {
	query := fmt.Sprintf("INSERT INTO payments (order_id, type, picture, status)"+
		" VALUES ('%d', '%s', '%s', '%s')", payment.OrderId, payment.Type, payment.Picture, payment.Status)
	_, err := repo.DB.Exec(query)

	return payment, err
}

func (repo *Repository) Update(payment *paymentPB.Payment) (*paymentPB.Payment, error) {
	query := fmt.Sprintf("UPDATE payments SET order_id = '%d', type = '%s', picture = %s, status = '%s'"+
		" WHERE id = %d", payment.OrderId, payment.Type, payment.Picture, payment.Status, payment.Id)
	_, err := repo.DB.Exec(query)

	return payment, err
}

func (repo *Repository) Destroy(payment *paymentPB.Payment) (*paymentPB.Payment, error) {
	query := fmt.Sprintf("DELETE FROM payments WHERE id=%d", payment.Id)
	_, err := repo.DB.Exec(query)

	return payment, err
}
