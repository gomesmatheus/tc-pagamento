package usecase

import (
	"github.com/gomesmatheus/tc-pagamento/infraestructure/external"
)

type pagamentoUseCases struct {
	client external.PedidoClient
}

func NewPagamentoUseCases(PedidoClient external.PedidoClient) *pagamentoUseCases {
	return &pagamentoUseCases{
		client: PedidoClient,
	}
}

func (usecase *pagamentoUseCases) AtualizarPagamento(id int, status bool) error {
	return usecase.client.AtualizarPagamento(id, status)
}
