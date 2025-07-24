package main

import (
	"fmt"
	"net/http"

	"github.com/omatheusq/rinha-backend-2025/internal/payments"
	"github.com/omatheusq/rinha-backend-2025/internal/summary"
	"github.com/redis/go-redis/v9"
)

func main() {
	opt, err := redis.ParseURL("redis://redis:6379")
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	//payment
	paymentService := payments.NewPaymentService(client)
	paymentHandler := payments.NewPaymentHandler(paymentService)
	http.Handle("/payments", paymentHandler)

	//summary
	summaryService := summary.NewSummaryService()
	summaryHandler := summary.NewSummaryHandler(summaryService)
	http.Handle("/payments-summary", summaryHandler)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
