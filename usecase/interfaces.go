package usecase

type PagamentoUseCases interface {
	AtualizarPagamento(id int, status bool) error
}
