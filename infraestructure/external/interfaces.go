package external

type PedidoClient interface {
	AtualizarPagamento(idPedido int, pagamentoAprovado bool) error
}
