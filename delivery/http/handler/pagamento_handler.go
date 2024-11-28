package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gomesmatheus/tc-pagamento/usecase"
)

type PagamentoHandler struct {
	pagamentoUseCases usecase.PagamentoUseCases
}

type Pagamento struct {
	Id int `json:"id"`
}

func NewPagamentoHandler(pagamentoUseCases usecase.PagamentoUseCases) *PagamentoHandler {
	return &PagamentoHandler{
		pagamentoUseCases: pagamentoUseCases,
	}
}

func (c *PagamentoHandler) AtualizarPagamentoRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var pagamento Pagamento
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			fmt.Println("Error parsing request body")
			w.WriteHeader(400)
			w.Write([]byte("400 bad request"))
			return
		}
		err = json.Unmarshal(body, &pagamento)
		if err != nil {
			fmt.Println("Error parsing request body")
			w.WriteHeader(400)
			w.Write([]byte("400 bad request"))
			return
		}

		// TEST Coverage fail
		if pagamento.Id == 123 {
			fmt.Println("Coverage fail test")
		}

		c.pagamentoUseCases.AtualizarPagamento(pagamento.Id, true)

		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("Pagamento %d confirmado!", pagamento.Id)))
	}

	return
}
