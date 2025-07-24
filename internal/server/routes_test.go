package server;

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"Pdf_Generator/internal/handlers"
)

// Minimal valid JSON input
var testPayload = `{
	"traveler_name": "Test User",
	"departure_from": "Delhi",
	"destination": "Paris",
	"departure_date": "2025-08-01",
	"arrival_date": "2025-08-10",
	"traveller_count": 2,
	"daily_itinerary": {
		"Day 1": {
			"morning": "Check-in",
			"afternoon": "City Tour",
			"evening": "Dinner Cruise"
		}
	},
	"flight_summary": [],
	"hotel_summary": [],
	"inclusion_summary": [],
	"activity_summary": [],
	"payment_plan": {
		"total_amount": "50000",
		"tcs": "500",
		"installments": []
	},
	"visa_details": {
		"visa_type": "Tourist",
		"validity": "6 months",
		"processing_date": "2025-07-20"
	}
}`

func TestGeneratePDFHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/generate", handlers.GeneratePDFHandler)

	req, _ := http.NewRequest("POST", "/generate", bytes.NewBuffer([]byte(testPayload)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", w.Code)
	}
}
