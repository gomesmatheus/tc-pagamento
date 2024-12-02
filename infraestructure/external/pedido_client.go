package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type PedidoHttpClient struct {
	BaseURL string
}

func NewPedidoHttpClient() *PedidoHttpClient {
	return &PedidoHttpClient{
		BaseURL: "http://svc-pedido-app",
	}
}

func (c *PedidoHttpClient) AtualizarPagamento(idPedido int, pagamentoAprovado bool) error {
	port := 80
	endpoint := fmt.Sprintf("%s:%d/pedido/atualizar/%d", c.BaseURL, port, idPedido)

	requestBody := map[string]interface{}{
		"status": "Pagamento Aprovado",
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Failed to marshal request body", err)
		return err
	}

	req, err := http.NewRequest("PATCH", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Failed to create HTTP request", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to call pedido-api", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Println("Failed to update pedido. Status:", resp.StatusCode)
		return fmt.Errorf("failed to update pedido with status: %d", resp.StatusCode)
	}

	fmt.Printf("Pedido %d atualizado com sucesso!\n", idPedido)
	return nil
}
