package entrypoints

import (
	"audiience_challenge/entities"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

// GetEstimate endpoint handler
func (s Server) GetEstimate(w http.ResponseWriter, r *http.Request) {

	log.Println("GetEstimate called...")

	distance, _ := strconv.ParseFloat(r.URL.Query().Get("distance"), 32)
	amount, _ := strconv.ParseFloat(r.URL.Query().Get("base_amount"), 32)

	inquiry := entities.InquiryModel{
		State:      r.URL.Query().Get("state"),
		Type:       r.URL.Query().Get("type"),
		Distance:   float32(distance),
		BaseAmount: float32(amount),
	}

	estimationAmount, err := s.estimate.Estimate(inquiry)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Estimation error: " + err.Error())
		return
	}

	response := entities.InquiryResponse{
		EstimatedAmount: estimationAmount,
		Date:            time.Now().Format(time.RFC3339),
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
	return

}
