// payment-service/handler/handler.go

package handler

import (
	"context"

	paymentPB "github.com/SleepingNext/payment-service/proto"
	paymentRepo "github.com/SleepingNext/payment-service/repository"
)

type handler struct {
	repository paymentRepo.Repository
}

func NewHandler(repo paymentRepo.Repository) *handler {
	return &handler{
		repository: repo,
	}
}

func (h *handler) IndexPayments(ctx context.Context, req *paymentPB.IndexPaymentsRequest, res *paymentPB.Response) error {
	payments, err := h.repository.Index()
	if err != nil {
		return err
	}

	res.Payments = payments
	res.Error = nil

	return err
}

func (h *handler) IndexPaymentsByUserID(ctx context.Context, req *paymentPB.User, res *paymentPB.Response) error {
	payments, err := h.repository.IndexByUserID(req)
	if err != nil {
		return err
	}

	res.Payments = payments
	res.Error = nil

	return err
}

func (h *handler) ShowPayment(ctx context.Context, req *paymentPB.Payment, res *paymentPB.Response) error {
	payment, err := h.repository.Show(req)
	if err != nil {
		return err
	}

	res.Payment = payment
	res.Error = nil

	return nil
}

func (h *handler) StorePayment(ctx context.Context, req *paymentPB.Payment, res *paymentPB.Response) error {
	payment, err := h.repository.Store(req)
	if err != nil {
		return err
	}

	res.Payment = payment
	res.Error = nil

	return err
}

func (h *handler) UpdatePayment(ctx context.Context, req *paymentPB.Payment, res *paymentPB.Response) error {
	payment, err := h.repository.Update(req)
	if err != nil {
		return err
	}

	res.Payment = payment
	res.Error = nil

	return nil
}


func (h *handler) DestroyPayment(ctx context.Context, req *paymentPB.Payment, res *paymentPB.Response) error {
	payment, err := h.repository.Destroy(req)
	if err != nil {
		return err
	}

	res.Payment = payment
	res.Error = nil

	return err
}
