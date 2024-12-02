package usecase

import (
	"errors"
	"testing"
)

// MockPedidoClient is a mock implementation of the PedidoClient interface
type MockPedidoClient struct {
	CalledWithID     int
	CalledWithStatus bool
	ReturnError      error
}

func (m *MockPedidoClient) AtualizarPagamento(id int, status bool) error {
	m.CalledWithID = id
	m.CalledWithStatus = status
	return m.ReturnError
}

func TestNewPagamentoUseCases(t *testing.T) {
	mockClient := &MockPedidoClient{}
	useCase := NewPagamentoUseCases(mockClient)

	if useCase == nil {
		t.Fatal("expected NewPagamentoUseCases to return a non-nil instance")
	}

	if useCase.client != mockClient {
		t.Errorf("expected client to be set to the provided mockClient")
	}
}

func TestAtualizarPagamento(t *testing.T) {
	t.Run("Successful call to AtualizarPagamento", func(t *testing.T) {
		mockClient := &MockPedidoClient{}
		useCase := NewPagamentoUseCases(mockClient)

		err := useCase.AtualizarPagamento(123, true)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		if mockClient.CalledWithID != 123 || mockClient.CalledWithStatus != true {
			t.Errorf("expected AtualizarPagamento to be called with id=123 and status=true, got id=%d and status=%v",
				mockClient.CalledWithID, mockClient.CalledWithStatus)
		}
	})

	t.Run("Error from AtualizarPagamento", func(t *testing.T) {
		mockClient := &MockPedidoClient{
			ReturnError: errors.New("mock error"),
		}
		useCase := NewPagamentoUseCases(mockClient)

		err := useCase.AtualizarPagamento(456, false)
		if err == nil {
			t.Fatal("expected an error, got nil")
		}

		if err.Error() != "mock error" {
			t.Errorf("expected error to be 'mock error', got '%v'", err)
		}

		if mockClient.CalledWithID != 456 || mockClient.CalledWithStatus != false {
			t.Errorf("expected AtualizarPagamento to be called with id=456 and status=false, got id=%d and status=%v",
				mockClient.CalledWithID, mockClient.CalledWithStatus)
		}
	})
}
