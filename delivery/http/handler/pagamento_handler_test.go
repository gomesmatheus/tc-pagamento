package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock implementation of the PagamentoUseCases interface
type mockPagamentoUseCases struct {
	calledWithID     int
	calledWithStatus bool
}

func (m *mockPagamentoUseCases) AtualizarPagamento(id int, status bool) error {
	m.calledWithID = id
	m.calledWithStatus = status
	return nil
}

func TestAtualizarPagamentoRoute(t *testing.T) {
	mockUseCases := &mockPagamentoUseCases{}
	handler := NewPagamentoHandler(mockUseCases)

	tests := []struct {
		name           string
		method         string
		body           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Successful POST request",
			method:         http.MethodPost,
			body:           `{"id":123}`,
			expectedStatus: http.StatusOK,
			expectedBody:   "Pagamento 123 confirmado!",
		},
		{
			name:           "Invalid body in POST request",
			method:         http.MethodPost,
			body:           `{"id":"abc"}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "400 bad request",
		},
		{
			name:           "Non-POST request",
			method:         http.MethodGet,
			body:           "",
			expectedStatus: http.StatusOK,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/pagamento", bytes.NewBuffer([]byte(tt.body)))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()

			handler.AtualizarPagamentoRoute(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			if rec.Body.String() != tt.expectedBody {
				t.Errorf("expected body %q, got %q", tt.expectedBody, rec.Body.String())
			}

			if tt.name == "Successful POST request" {
				if mockUseCases.calledWithID != 123 || mockUseCases.calledWithStatus != true {
					t.Errorf("expected AtualizarPagamento to be called with id=123 and status=true, got id=%d, status=%v", mockUseCases.calledWithID, mockUseCases.calledWithStatus)
				}
			}
		})
	}
}
