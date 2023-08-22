package gateway

import "github.com.br/lenodeoliveira/fc-ms-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
