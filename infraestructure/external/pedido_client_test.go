package external

import (
	"testing"
)

func TestNewPedidoHttpClient(t *testing.T) {
	client := NewPedidoHttpClient()

	if client.BaseURL != "http://localhost:8080" {
		t.Errorf("expected BaseURL to be 'http://localhost:8080', got '%s'", client.BaseURL)
	}
}

func TestAtualizarPagamento(t *testing.T) {
	client := NewPedidoHttpClient()

	err := client.AtualizarPagamento(123, true)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

}
