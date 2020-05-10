/*
Dear Programmers,

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*                                                 *
*	This file belongs to Kevin Veros Hamonangan   *
*	and	Fandi Fladimir Dachi and is a part of     *
*	our	last project as the student of Del        *
*	Institute of Technology, Sitoluama.           *
*	Please contact us via Instagram:              *
*	sleepingnext and fandi_dachi                  *
*	before copying this file.                     *
*	Thank you, buddy. 😊                          *
*                                                 *
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*/

package postgres

import (
	"database/sql"
	"fmt"

	paymentPB "github.com/ta04/payment-service/proto"
)

// Postgres is the implementor of Postgres interface
type Postgres struct {
	DB *sql.DB
}

// Index indexes all payments
func (repo *Postgres) Index(req *paymentPB.IndexPaymentsRequest) (payments []*paymentPB.Payment, err error) {
	var id, orderID, paymentMethodID int32
	var status string

	query := "SELECT * FROM payments"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &orderID, &paymentMethodID, &status)
		if err != nil {
			return nil, err
		}
		payment := &paymentPB.Payment{
			Id:              id,
			OrderId:         orderID,
			PaymentMethodId: paymentMethodID,
			Status:          status,
		}
		payments = append(payments, payment)
	}

	return payments, err
}

// Show shows a payment by id
func (repo *Postgres) Show(payment *paymentPB.Payment) (*paymentPB.Payment, error) {
	var id, orderID, paymentMethodID int32
	var status string
	query := fmt.Sprintf("SELECT * FROM payments WHERE id = %d", payment.Id)
	err := repo.DB.QueryRow(query).Scan(&id, &orderID, &paymentMethodID, &status)
	if err != nil {
		return nil, err
	}

	return &paymentPB.Payment{
		Id:              id,
		OrderId:         orderID,
		PaymentMethodId: paymentMethodID,
		Status:          status,
	}, err
}

// ShowByOrderID shows a payment by orderID
func (repo *Postgres) ShowByOrderID(order *paymentPB.Order) (*paymentPB.Payment, error) {
	var id, orderID, paymentMethodID int32
	var status string
	query := fmt.Sprintf("SELECT * FROM payments WHERE order_id = %d", order.Id)
	err := repo.DB.QueryRow(query).Scan(&id, &orderID, &paymentMethodID, &status)
	if err != nil {
		return nil, err
	}

	return &paymentPB.Payment{
		Id:              id,
		OrderId:         orderID,
		PaymentMethodId: paymentMethodID,
		Status:          status,
	}, err
}
