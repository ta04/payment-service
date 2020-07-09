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

package handler

import (
	"context"
	"errors"

	proto "github.com/ta04/payment-service/model/proto"
	"github.com/ta04/payment-service/usecase"
)

// Handler is the handler of payment service
type Handler struct {
	Usecase usecase.Usecase
}

// NewHandler will create a new payment handler
func NewHandler(Usecase usecase.Usecase) *Handler {
	return &Handler{
		Usecase: Usecase,
	}
}

// GetAllPayments will get all payments
func (handler *Handler) GetAllPayments(ctx context.Context, req *proto.GetAllPaymentsRequest, res *proto.Response) error {
	payments, err := handler.Usecase.GetAll(req)
	if err != nil {
		res.Payments = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.Payments = payments
	res.Error = nil

	return nil
}

// GetOnePayment will get a payment
func (handler *Handler) GetOnePayment(ctx context.Context, req *proto.GetOnePaymentRequest, res *proto.Response) error {
	payment, err := handler.Usecase.GetOne(req)
	if err != nil {
		res.Payment = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.Payment = payment
	res.Error = nil

	return nil
}

// CreateOnePayment will create a new payment
func (handler *Handler) CreateOnePayment(ctx context.Context, req *proto.Payment, res *proto.Response) error {
	payment, err := handler.Usecase.CreateOne(req)
	if err != nil {
		res.Payment = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.Payment = payment
	res.Error = nil

	return nil
}

// UpdateOnePayment will update a payment
func (handler *Handler) UpdateOnePayment(ctx context.Context, req *proto.Payment, res *proto.Response) error {
	payment, err := handler.Usecase.UpdateOne(req)
	if err != nil {
		res.Payment = nil
		res.Error = err

		return errors.New(err.Message)
	}

	res.Payment = payment
	res.Error = nil

	return nil
}
