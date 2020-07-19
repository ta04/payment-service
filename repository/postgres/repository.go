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
	"errors"
	"fmt"

	paymentPB "github.com/ta04/payment-service/model/proto"
)

// CreateOne will create a new payment
func (postgres *Postgres) CreateOne(payment *paymentPB.Payment) (*paymentPB.Payment, error) {
	query := fmt.Sprintf("INSERT INTO payments (order_id, payment_method_id, status)"+
		" VALUES (%d, %d, 'unpaid')", payment.OrderId, payment.PaymentMethodId)
	_, err := postgres.DB.Exec(query)

	return payment, err
}

// UpdateOne will update a payment
func (postgres *Postgres) UpdateOne(payment *paymentPB.Payment) (*paymentPB.Payment, error) {
	query := fmt.Sprintf("UPDATE payments SET order_id = %d, payment_method_id = %d, status = '%s'"+
		" WHERE id = %d", payment.OrderId, payment.PaymentMethodId, payment.Status, payment.Id)
	res, err := postgres.DB.Exec(query)

	count, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if count <= 0 {
		return nil, errors.New("sql: no rows found")
	}

	return payment, err
}
