package createaccount

import (
	"github.com.br/lenodeoliveira/fc-ms-wallet/internal/entity"
	"github.com.br/lenodeoliveira/fc-ms-wallet/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string
}

type CreateAccountOutPutDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(a gateway.AccountGateway, c gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: a,
		ClientGateway:  c,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutPutDTO, error) {
	client, err := uc.ClientGateway.Get(input.ClientID)

	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(client)
	err = uc.AccountGateway.Save(account)

	if err != nil {
		return nil, err
	}

	return &CreateAccountOutPutDTO{
		ID: account.ID,
	}, nil

}
