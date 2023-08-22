package gateway

import "github.com.br/lenodeoliveira/fc-ms-wallet/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(cliente *entity.Client) error
}
