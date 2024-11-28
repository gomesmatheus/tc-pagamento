package external

import (
	"fmt"
	"net/http"
	"time"
)

type PedidoHttpClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewPedidoHttpClient() *PedidoHttpClient {
	return &PedidoHttpClient{
		BaseURL: "http://localhost:8080",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *PedidoHttpClient) AtualizarPagamento(idPedido int, pagamentoAprovado bool) error {
	// TODO adicionar chamada para API de pedidos
	fmt.Println(fmt.Sprintf("Pedido %d atualizado com sucesso!", idPedido))
	return nil
}
