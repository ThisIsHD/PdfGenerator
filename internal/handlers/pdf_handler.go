package handlers

import (
	"fmt"
	"net/http"
	"time"

	"Pdf_Generator/internal/models"
	"Pdf_Generator/internal/pdf"

	"github.com/gin-gonic/gin"
)

func GeneratePDFHandler(c *gin.Context) {
	var req models.PDFRequest

	// Bind incoming JSON to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Generate unique filename with timestamp
	timestamp := time.Now().Format("20060102_150405")
	fileName := fmt.Sprintf("itinerary_%s.pdf", timestamp)
	outputPath := "pdf/" + fileName

	// Generate PDF
	err := pdf.GeneratePDFWithOverlay(&req, outputPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate PDF", "details": err.Error()})
		return
	}

	// Create PDF URL (served via /pdfs)
	pdfURL := fmt.Sprintf("http://localhost:8080/pdfs/%s", fileName)

	// Return URL to client
	c.JSON(http.StatusOK, gin.H{
		"message": "PDF successfully generated",
		"url":     pdfURL,
	})
}
