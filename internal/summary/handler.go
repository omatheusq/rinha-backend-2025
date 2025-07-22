package summary

import (
	"encoding/json"
	"net/http"
)

type SummaryHandler struct {
	Service *SummaryService
}

func NewSummaryHandler(service *SummaryService) *SummaryHandler {
	return &SummaryHandler{Service: service}
}

func (h *SummaryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": h.Service.GetSummary(),
	})
}
