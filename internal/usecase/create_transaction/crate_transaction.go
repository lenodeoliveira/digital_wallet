package createtransaction

import (
	"github.com.br/lenodeoliveira/fc-ms-wallet/internal/entity"
	"github.com.br/lenodeoliveira/fc-ms-wallet/internal/gateway"
)

type CreateTransactionInputDto struct {
	AccountIDFrom string
	AccountIDTo   string
	Amount        float64
}

type CreateTransactionOutPutDTO struct {
	ID string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
}

func NewTransactionUseCase(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
	}
}

func (uc *CreateTransactionUseCase) Execute(input CreateTransactionInputDto) (*CreateTransactionOutPutDTO, error) {
	accountFrom, err := uc.AccountGateway.FindById(input.AccountIDFrom)

	if err != nil {
		return nil, err
	}

	accountTo, err := uc.AccountGateway.FindById(input.AccountIDTo)

	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)

	if err != nil {
		return nil, err
	}

	err = uc.TransactionGateway.Create(transaction)

	if err != nil {
		return nil, err
	}

	return &CreateTransactionOutPutDTO{ID: transaction.ID}, nil

}
