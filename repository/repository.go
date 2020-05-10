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
	paymentPB "github.com/ta04/payment-service/proto"
)

// Repository is the interface of repositories.
// As there are number of repositories can be used.
type Repository interface {
	Index(*paymentPB.IndexPaymentsRequest) ([]*paymentPB.Payment, error)
	Show(*paymentPB.Payment) (*paymentPB.Payment, error)
	ShowByOrderID(*paymentPB.Order) (*paymentPB.Payment, error)
	Store(*paymentPB.Payment) (*paymentPB.Payment, error)
	Update(*paymentPB.Payment) (*paymentPB.Payment, error)
}
