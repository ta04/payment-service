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

package repository

import (
	proto "github.com/ta04/payment-service/model/proto"
)

// Repository is the interface of repositories.
// As there are number of repositories can be used.
type Repository interface {
	GetAll(request *proto.GetAllPaymentsRequest) (*[]*proto.Payment, error)
	GetOneByOrderID(request *proto.GetOnePaymentRequest) (*proto.Payment, error)
	GetOne(request *proto.GetOnePaymentRequest) (*proto.Payment, error)
	CreateOne(payment *proto.Payment) (*proto.Payment, error)
	UpdateOne(payment *proto.Payment) (*proto.Payment, error)
}
