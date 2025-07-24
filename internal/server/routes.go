package server

import (
	"net/http"

	"Pdf_Generator/internal/handlers" 
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Adjust as needed
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// ✅ Register your real routes here
	r.POST("/generate", handlers.GeneratePDFHandler)

	// (Optional) Static file serving for generated PDFs
	r.Static("/pdfs", "./pdf")

	return r
}
