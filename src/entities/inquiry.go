package entities

// InquiryModel core model for estimation inquiries
type InquiryModel struct {
	State      string  `json:"state" validate:"required"`
	Type       string  `json:"type" validate:"required"`
	Distance   float32 `json:"distance" validate:"required"`
	BaseAmount float32 `json:"baseAmount" validate:"required"`
}

// InquiryResponse response model for estimate endpoint
type InquiryResponse struct {
	EstimatedAmount float32 `json:"estimatedAmount"`
	Date            string  `json:"date"`
}
