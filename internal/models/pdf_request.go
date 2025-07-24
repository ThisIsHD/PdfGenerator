package models

type PDFRequest struct {
	TravelerName     string              `json:"traveler_name" binding:"required"`
	DepartureFrom    string              `json:"departure_from" binding:"required"`
	Destination      string              `json:"destination" binding:"required"`
	DepartureDate    string              `json:"departure_date" binding:"required"`
	ArrivalDate      string              `json:"arrival_date" binding:"required"`
	TravellerCount   int                 `json:"traveller_count" binding:"required"`
	DailyItinerary   map[string]DayPlan  `json:"daily_itinerary" binding:"required"`
	FlightSummary    []FlightInfo        `json:"flight_summary" binding:"required"`
	HotelSummary     []HotelInfo         `json:"hotel_summary" binding:"required"`
	InclusionSummary []InclusionItem     `json:"inclusion_summary" binding:"required"`
	ActivitySummary  []ActivityInfo      `json:"activity_summary" binding:"required"`
	PaymentPlan      PaymentDetail       `json:"payment_plan" binding:"required"`
	VisaDetails      VisaInfo            `json:"visa_details" binding:"required"`
}

type DayPlan struct {
	Morning   string `json:"morning"`
	Afternoon string `json:"afternoon"`
	Evening   string `json:"evening"`
}
type FlightInfo struct {
	Date         string `json:"date"`
	FlightDetail string `json:"flight_detail"`
}
type HotelInfo struct {
	City      string `json:"city"`
	CheckIn   string `json:"checkin"`
	CheckOut  string `json:"checkout"`
	Nights    int    `json:"nights"`
	HotelName string `json:"hotel_name"`
}
type InclusionItem struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
	Details  string `json:"details"`
	Status   string `json:"status"`
}
type ActivityInfo struct {
	City         string `json:"city"`
	Activity     string `json:"activity"`
	Type         string `json:"type"`
	TimeRequired string `json:"time_required"`
}
type PaymentDetail struct {
	TotalAmount  string        `json:"total_amount"`
	TCS          string        `json:"tcs"`
	Installments []Installment `json:"installments"`
}
type Installment struct {
	Installment string `json:"installment"`
	Amount      string `json:"amount"`
	DueDate     string `json:"due_date"`
}
type VisaInfo struct {
	VisaType       string `json:"visa_type"`
	Validity       string `json:"validity"`
	ProcessingDate string `json:"processing_date"`
}
