
# 🧾 Vigovia PDF Generator

This project is a **PDF generation microservice** built with **Go** that accepts JSON input (an itinerary) and overlays it onto a pre-designed PDF template. The output PDF is saved and served via an HTTP endpoint.

---

## ✨ Features

- Accepts a detailed itinerary in JSON format.
- Overlays text on a pre-designed PDF using precise coordinates.
- Returns a downloadable **PDF URL** after generation.
- Built using:
  - `Gin` web framework
  - `pdfcpu` for watermark-based text overlays
  - JSON models for input
- Easily extensible to switch to HTML-to-PDF in the future.

---

## 📁 Project Structure

```
Pdf_Generator/
│
├── assets/                 # Contains the Template.pdf
│   └── Template.pdf
│
├── cmd/
│   └── api/
│       └── main.go         # Entry point of the server
│
├── internal/
│   ├── handlers/           # HTTP handler logic
│   │   └── pdf_handler.go
│   │
│   ├── models/             # Request model for PDF generation
│   │   └── pdf_request.go
│   │
│   ├── pdf/                # Core PDF generation logic using pdfcpu
│   │   └── generator.go
│   │
│   └── server/             # Route registration and CORS setup
│       └── router.go
│
├── pdf/                    # Output PDFs get stored here
│   └── itinerary_*.pdf
│
└── go.mod / go.sum         # Go dependencies
```

---

## 🔧 Setup Instructions

### ✅ Prerequisites

- Go 1.18+
- Git

---

### 🚀 Run the Server

1. **Clone the repo**

```bash
git clone https://github.com/your-username/Pdf_Generator.git
cd Pdf_Generator
```

2. **Install dependencies**

```bash
go mod tidy
```

3. **Run the server**

```bash
go run cmd/api/main.go
```

By default, the server runs at: `http://localhost:8080`

---

## 📤 API Endpoint

### `POST /generate`

Generates a customized itinerary PDF.

#### 📄 Request Body (JSON)

```json
{
  "TravelerName": "John Doe",
  "DepartureFrom": "New York",
  "Destination": "Paris",
  "DepartureDate": "2025-08-01",
  "ArrivalDate": "2025-08-10",
  "TravellerCount": 2,
  "FlightSummary": [
    { "Date": "2025-08-01", "FlightDetail": "Flight ABC123" }
  ],
  "HotelSummary": [
    { "City": "Paris", "CheckIn": "2025-08-01", "CheckOut": "2025-08-05", "Nights": 4, "HotelName": "Hotel Parisienne" }
  ],
  "InclusionSummary": [
    { "Category": "Meals", "Count": 2, "Details": "Breakfast + Dinner", "Status": "Included" }
  ],
  "ActivitySummary": [
    { "City": "Paris", "Activity": "Eiffel Tower Visit", "Type": "Sightseeing", "TimeRequired": "2 hours" }
  ],
  "PaymentPlan": {
    "TotalAmount": "$2000",
    "TCS": "$200",
    "Installments": [
      { "Installment": "1st", "Amount": "$1000", "DueDate": "2025-07-01" },
      { "Installment": "2nd", "Amount": "$1000", "DueDate": "2025-08-01" }
    ]
  },
  "VisaDetails": {
    "VisaType": "Tourist",
    "Validity": "30 Days",
    "ProcessingDate": "2025-07-15"
  }
}
```

---

### ✅ Response

```json
{
  "message": "PDF successfully generated",
  "file_url": "http://localhost:8080/pdfs/itinerary_20250724_221918.pdf"
}
```

---

## 🧠 How It Works

1. A POST request hits the `/generate` endpoint with itinerary details.
2. `handlers.GeneratePDFHandler` parses the request.
3. `pdf.GeneratePDFWithOverlay`:
   - Copies `Template.pdf` from `assets/`
   - Creates text overlays with positions
   - Uses `pdfcpu` to apply text overlays as watermarks
4. The PDF is saved to `pdf/` and served via static URL.

---

## 📌 Notes

- Coordinates for overlay are manually adjusted using a PDF viewer with measure tools (e.g., PDF XChange Editor).
- The service currently uses only page 1 of the template.
- `pdfcpu.ParseTextWatermarkDetails()` handles the styling, position, and font.

---

## 💡 Future Improvements

- Replace PDF coordinate overlay with HTML-to-PDF rendering using headless Chrome or wkhtmltopdf.
- Add support for multi-page itineraries.
- Allow customization of fonts/colors via API.
- Deploy on cloud and generate temporary download URLs.

---

## 🛠 Tools & Libraries

| Tool         | Purpose                           |
|--------------|-----------------------------------|
| Gin          | HTTP routing and middleware       |
| pdfcpu       | PDF manipulation & watermarking   |
| Go Templates | Dynamic HTML/PDF (if used later)  |

---

## 🤝 License

MIT License © 2025 Himadri Dey
