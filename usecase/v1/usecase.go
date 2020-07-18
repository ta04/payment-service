package v1

import (
	"net/http"

	proto "github.com/ta04/payment-service/model/proto"
	"github.com/ta04/payment-service/repository"
)

// usecase is the struct of payment usecase
type Usecase struct {
	Repository repository.Repository
}

// NewUsecase will create a new payment usecase
func NewUsecase(repository repository.Repository) *Usecase {
	return &Usecase{
		Repository: repository,
	}
}

func (usecase *Usecase) GetAll(request *proto.GetAllPaymentsRequest) (*[]*proto.Payment, *proto.Error) {
	if request == nil {
		return nil, &proto.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	if request.Status == "" {
		request.Status = "waiting for payment"
	}

	var err error
	payment, err := usecase.Repository.GetAll(request)

	if payment == nil || len(*payment) <= 0 {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: "no payment found",
		}
	}

	if err != nil {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return payment, nil
}

func (usecase *Usecase) GetOne(request *proto.GetOnePaymentRequest) (*proto.Payment, *proto.Error) {
	if request == nil || (request.Id == 0 && request.OrderId == 0) {
		return nil, &proto.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	var payment *proto.Payment
	var err error
	if request.OrderId != 0 {
		payment, err = usecase.Repository.GetOneByOrderID(request)
		if err != nil {
			return nil, &proto.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	} else {
		payment, err = usecase.Repository.GetOne(request)
		if err != nil {
			return nil, &proto.Error{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			}
		}
	}

	return payment, nil
}

func (usecase *Usecase) CreateOne(payment *proto.Payment) (*proto.Payment, *proto.Error) {
	if payment == nil {
		return nil, &proto.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	payment, err := usecase.Repository.CreateOne(payment)
	if err != nil {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return payment, nil
}

func (usecase *Usecase) UpdateOne(payment *proto.Payment) (*proto.Payment, *proto.Error) {
	if payment == nil {
		return nil, &proto.Error{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}
	}

	payment, err := usecase.Repository.UpdateOne(payment)
	if err != nil {
		return nil, &proto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return payment, nil
}
