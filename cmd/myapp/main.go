package main

import (
	"fmt"
	"log"
	"net/http"

	handlers "github.com/gomesmatheus/tc-pagamento/delivery/http/handler"
	"github.com/gomesmatheus/tc-pagamento/infraestructure/external"
	usecase "github.com/gomesmatheus/tc-pagamento/usecase/pagamento"
)

func main() {
	pedidoClient := external.NewPedidoHttpClient()
	pagamentoUseCases := usecase.NewPagamentoUseCases(pedidoClient)
	handler := handlers.NewPagamentoHandler(pagamentoUseCases)

	http.HandleFunc("/pagamento/webhooks/", handler.AtualizarPagamentoRoute)
	fmt.Println("Pagamento ms running!")

	log.Fatal(http.ListenAndServe(":3333", nil))
}
