// payment-service/repository/repository.go

package repository

import (
	paymentPB "github.com/SleepingNext/payment-service/proto"
)

type Repository interface {
	Index() ([]*paymentPB.Payment, error)
	IndexByUserID(*paymentPB.User) ([]*paymentPB.Payment, error)
	Show(*paymentPB.Payment) (*paymentPB.Payment, error)
	Store(*paymentPB.Payment) (*paymentPB.Payment, error)
	Update(*paymentPB.Payment) (*paymentPB.Payment, error)
	Destroy(*paymentPB.Payment) (*paymentPB.Payment, error)
}
