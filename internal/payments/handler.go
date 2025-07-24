package payments

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Payment struct {
	CorrelationID string    `json:"correlationId"`
	Amount        float64   `json:"amount"`
	Timestamp     time.Time `json:"timestamp"`
}

type PaymentHandler struct {
	Service *PaymentService
}

func NewPaymentHandler(service *PaymentService) *PaymentHandler {
	return &PaymentHandler{Service: service}
}

func (h *PaymentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var p Payment
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	p.Timestamp = time.Now()

	ctx := r.Context()

	if err := h.Service.QueuePayment(ctx, p); err != nil {
		log.Println("Payment failed:", err)
		http.Error(w, "Failed to process payment", http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}
