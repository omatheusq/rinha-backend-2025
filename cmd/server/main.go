package main

import (
	"fmt"
	"github.com/omatheusq/rinha-backend-2025/internal/payments"
	"github.com/omatheusq/rinha-backend-2025/internal/summary"
	"net/http"
)

func main() {

	//payment
	paymentService := payments.NewPaymentService()
	paymentHandler := payments.NewPaymentHandler(paymentService)
	http.Handle("/payments", paymentHandler)

	//summary
	summaryService := summary.NewSummaryService()
	summaryHandler := summary.NewSummaryHandler(summaryService)
	http.Handle("/payments-summary", summaryHandler)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
