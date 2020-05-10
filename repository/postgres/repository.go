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
	"fmt"

	paymentPB "github.com/ta04/payment-service/proto"
)

// Store stores a new payment
func (repo *Postgres) Store(payment *paymentPB.Payment) (*paymentPB.Payment, error) {
	query := fmt.Sprintf("INSERT INTO payments (order_id, payment_method_id, status)"+
		" VALUES (%d, %d, 'unpaid')", payment.OrderId, payment.PaymentMethodId)
	_, err := repo.DB.Exec(query)

	return payment, err
}

// Update updates a payment
func (repo *Postgres) Update(payment *paymentPB.Payment) (*paymentPB.Payment, error) {
	query := fmt.Sprintf("UPDATE payments SET order_id = %d, payment_method_id = %d, status = '%s'"+
		" WHERE id = %d", payment.OrderId, payment.PaymentMethodId, payment.Status, payment.Id)
	_, err := repo.DB.Exec(query)

	return payment, err
}
