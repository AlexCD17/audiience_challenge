package entities

type InquiryModel struct {
	State      string  `json:"state" validate:"required"`
	Type       string  `json:"type" validate:"required"`
	Distance   float32 `json:"distance" validate:"required"`
	BaseAmount float32 `json:"baseAmount" validate:"required"`
}

type InquiryResponse struct {
	EstimatedAmount float32 `json:"estimatedAmount"`
	Date            string  `json:"date"`
}
