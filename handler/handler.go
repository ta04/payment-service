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

	paymentPB "github.com/ta04/payment-service/proto"
	paymentRepo "github.com/ta04/payment-service/repository"
)

// Handler is the handler of payment service
type Handler struct {
	repository paymentRepo.Repository
}

// NewHandler creates a new payment service handler
func NewHandler(repo paymentRepo.Repository) *Handler {
	return &Handler{
		repository: repo,
	}
}

// IndexPayments indexes the payment
func (h *Handler) IndexPayments(ctx context.Context, req *paymentPB.IndexPaymentsRequest, res *paymentPB.Response) error {
	payments, err := h.repository.Index(req)
	if err != nil {
		return err
	}

	res.Payments = payments
	res.Error = nil

	return err
}

// ShowPayment shows a payment by ID
func (h *Handler) ShowPayment(ctx context.Context, req *paymentPB.Payment, res *paymentPB.Response) error {
	payment, err := h.repository.Show(req)
	if err != nil {
		return err
	}

	res.Payment = payment
	res.Error = nil

	return nil
}

// ShowPaymentByOrderID shows a payment by order ID
func (h *Handler) ShowPaymentByOrderID(ctx context.Context, req *paymentPB.Order, res *paymentPB.Response) error {
	payment, err := h.repository.ShowByOrderID(req)
	if err != nil {
		return err
	}

	res.Payment = payment
	res.Error = nil

	return nil
}

// StorePayment stores a new payment
func (h *Handler) StorePayment(ctx context.Context, req *paymentPB.Payment, res *paymentPB.Response) error {
	payment, err := h.repository.Store(req)
	if err != nil {
		return err
	}

	res.Payment = payment
	res.Error = nil

	return err
}

// UpdatePayment updates a payment
func (h *Handler) UpdatePayment(ctx context.Context, req *paymentPB.Payment, res *paymentPB.Response) error {
	payment, err := h.repository.Update(req)
	if err != nil {
		return err
	}

	res.Payment = payment
	res.Error = nil

	return nil
}
