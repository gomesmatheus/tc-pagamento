package external

import (
	"testing"
	"time"
)

func TestNewPedidoHttpClient(t *testing.T) {
	client := NewPedidoHttpClient()

	if client.BaseURL != "http://localhost:8080" {
		t.Errorf("expected BaseURL to be 'http://localhost:8080', got '%s'", client.BaseURL)
	}

	if client.HTTPClient == nil {
		t.Fatalf("expected HTTPClient to be initialized, got nil")
	}

	if client.HTTPClient.Timeout != 10*time.Second {
		t.Errorf("expected HTTPClient timeout to be 10 seconds, got %v", client.HTTPClient.Timeout)
	}
}

func TestAtualizarPagamento(t *testing.T) {
	client := NewPedidoHttpClient()

	err := client.AtualizarPagamento(123, true)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

}
