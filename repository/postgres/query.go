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

	proto "github.com/ta04/payment-service/model/proto"
)

// Postgres is the implementor of Postgres interface
type Postgres struct {
	DB *sql.DB
}

// NewPostgres will create a new postgres instance
func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{
		DB: db,
	}
}

// GetAll will get all payments
func (postgres *Postgres) GetAll(request *proto.GetAllPaymentsRequest) (payments *[]*proto.Payment, err error) {
	var id, orderID, paymentMethodID int32
	var status string

	query := fmt.Sprintf("SELECT * FROM payments WHERE status = '%s'", request.Status)
	rows, err := postgres.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&id, &orderID, &paymentMethodID, &status)
		if err != nil {
			return nil, err
		}

		payments := &proto.Payment{
			Id:              id,
			OrderId:         orderID,
			PaymentMethodId: paymentMethodID,
			Status:          status,
		}
		payments = append(payments, payment)
	}

	return &payments, err
}

// GetOneByOrderID will get a payment by order id
func (postgres *Postgres) GetOneByOrderID(request *proto.GetOnePaymentRequest) (*proto.Payment, error) {
	var id, orderID, paymentMethodID int32
	var status string
	query := fmt.Sprintf("SELECT * FROM payments WHERE order_id = %d", request.Id)
	err := postgres.DB.QueryRow(query).Scan(&id, &orderID, &paymentMethodID, &status)
	if err != nil {
		return nil, err
	}

	return &proto.Payment{
		Id:              id,
		OrderId:         orderID,
		PaymentMethodId: paymentMethodID,
		Status:          status,
	}, err
}

// GetOne will get a payment by id
func (postgres *Postgres) GetOne(request *proto.GetOnePaymentRequest) (*proto.Payment, error) {
	var id, orderID, paymentMethodID int32
	var status string
	query := fmt.Sprintf("SELECT * FROM payments WHERE id = %d", request.Id)
	err := postgres.DB.QueryRow(query).Scan(&id, &orderID, &paymentMethodID, &status)
	if err != nil {
		return nil, err
	}

	return &proto.Payment{
		Id:              id,
		OrderId:         orderID,
		PaymentMethodId: paymentMethodID,
		Status:          status,
	}, err
}
