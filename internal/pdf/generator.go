package pdf

import (
	"fmt"
	"os"

	"Pdf_Generator/internal/models"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

func GeneratePDFWithOverlay(req *models.PDFRequest, outputPath string) error {
	templatePath := "assets/Template.pdf"

	// Copy template to destination
	if err := copyFile(templatePath, outputPath); err != nil {
		return fmt.Errorf("failed to copy template: %w", err)
	}

	conf := api.LoadConfiguration()
	lineSpacing := 18
	var overlays []string

	// Basic Info
	overlays = append(overlays,
		fmt.Sprintf("pos:100 720, font:Helvetica, points:12, fillcol:#000000, Hi, %s!", req.TravelerName),
		fmt.Sprintf("pos:90 685, font:Helvetica, points:12, fillcol:#000000, %s", req.DepartureFrom),
		fmt.Sprintf("pos:325 685, font:Helvetica, points:12, fillcol:#000000, %s", req.Destination),
		fmt.Sprintf("pos:90 660, font:Helvetica, points:12, fillcol:#000000, %s", req.DepartureDate),
		fmt.Sprintf("pos:325 660, font:Helvetica, points:12, fillcol:#000000, %s", req.ArrivalDate),
		fmt.Sprintf("pos:325 635, font:Helvetica, points:12, fillcol:#000000, %d", req.TravellerCount),
	)

	// Flight Summary
	y := 515
	for _, flight := range req.FlightSummary {
		overlays = append(overlays, fmt.Sprintf("pos:70 %d, font:Helvetica, points:11, fillcol:#000000, %s - %s", y, flight.Date, flight.FlightDetail))
		y -= lineSpacing
	}

	// Hotel Summary
	y = 430
	for _, h := range req.HotelSummary {
		overlays = append(overlays, fmt.Sprintf("pos:70 %d, font:Helvetica, points:11, fillcol:#000000, %s: %s to %s (%d nights) - %s",
			y, h.City, h.CheckIn, h.CheckOut, h.Nights, h.HotelName))
		y -= lineSpacing
	}

	// Inclusion Summary
	y = 300
	for _, inc := range req.InclusionSummary {
		overlays = append(overlays, fmt.Sprintf("pos:70 %d, font:Helvetica, points:11, fillcol:#000000, %s | %d | %s | %s",
			y, inc.Category, inc.Count, inc.Details, inc.Status))
		y -= lineSpacing
	}

	// Activity Summary
	y = 200
	for _, act := range req.ActivitySummary {
		overlays = append(overlays, fmt.Sprintf("pos:70 %d, font:Helvetica, points:11, fillcol:#000000, %s | %s | %s | %s",
			y, act.City, act.Activity, act.Type, act.TimeRequired))
		y -= lineSpacing
	}

	// Payment Plan
	y = 120
	overlays = append(overlays, fmt.Sprintf("pos:70 %d, font:Helvetica, points:11, fillcol:#000000, Total: %s | TCS: %s",
		y, req.PaymentPlan.TotalAmount, req.PaymentPlan.TCS))
	y -= lineSpacing
	for _, inst := range req.PaymentPlan.Installments {
		overlays = append(overlays, fmt.Sprintf("pos:70 %d, font:Helvetica, points:11, fillcol:#000000, %s - %s due %s",
			y, inst.Installment, inst.Amount, inst.DueDate))
		y -= lineSpacing
	}

	// Visa Info
	y = 70
	overlays = append(overlays,
		fmt.Sprintf("pos:70 %d, font:Helvetica, points:11, fillcol:#000000, Visa Type: %s", y, req.VisaDetails.VisaType),
		fmt.Sprintf("pos:70 %d, font:Helvetica, points:11, fillcol:#000000, Validity: %s", y-lineSpacing, req.VisaDetails.Validity),
		fmt.Sprintf("pos:70 %d, font:Helvetica, points:11, fillcol:#000000, Processing Date: %s", y-2*lineSpacing, req.VisaDetails.ProcessingDate),
	)

	// Apply each overlay as a watermark
	for _, desc := range overlays {
		wm, err := pdfcpu.ParseTextWatermarkDetails(desc, "", true, types.DisplayUnit(0))
		if err != nil {
			return fmt.Errorf("failed to parse watermark details: %w", err)
		}
		err = api.AddWatermarksFile(outputPath, outputPath, []string{"1"}, wm, conf)
		if err != nil {
			return fmt.Errorf("failed to add watermark: %w", err)
		}
	}

	return nil
}

// Utility function to copy file
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.ReadFrom(in)
	return err
}
