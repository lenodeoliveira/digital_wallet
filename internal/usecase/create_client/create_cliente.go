package createclient

import (
	"time"

	"github.com.br/lenodeoliveira/fc-ms-wallet/internal/entity"
	"github.com.br/lenodeoliveira/fc-ms-wallet/internal/gateway"
)

type CreateClientInputDTO struct {
	Name  string
	Email string
}

type CreateClientOutPutDTO struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateClientUseCase struct {
	ClientGateway gateway.ClientGateway
}

func NewCreateClientUseCase(clientGateway gateway.ClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientGateway: clientGateway,
	}
}

func (uc *CreateClientUseCase) Execute(input CreateClientInputDTO) (*CreateClientOutPutDTO, error) {
	client, err := entity.NewClient(input.Name, input.Email)

	if err != nil {
		return nil, err
	}

	err = uc.ClientGateway.Save(client)

	if err != nil {
		return nil, err
	}

	return &CreateClientOutPutDTO{
		ID:    client.ID,
		Name:  client.Name,
		Email: client.Email,
	}, nil

}
