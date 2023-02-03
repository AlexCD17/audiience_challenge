package services

import (
	"audiience_challenge/entities"
	"audiience_challenge/repositories/rates"
	"strings"
)

type IService interface {
	Estimate(inquiry entities.InquiryModel) (float32, error)
}

type Service struct {
	ratesRepo rates.IRepository
}

func NewEstimateService(ratesRepo rates.IRepository) *Service {
	return &Service{
		ratesRepo: ratesRepo,
	}
}

// Estimate calculates the final estimation
func (e *Service) Estimate(inquiry entities.InquiryModel) (finalEstimation float32, err error) {

	state := strings.ToUpper(inquiry.State)
	estimationType := strings.ToLower(inquiry.Type)

	rate, err := e.ratesRepo.GetRates(state, estimationType)
	if err != nil {
		return 0, err
	}

	baseEstimation := inquiry.BaseAmount + (inquiry.BaseAmount * rate)

	tax := getTax(state, baseEstimation)

	taxedEstimation := baseEstimation + tax

	// apply discounts based on estimation type
	if estimationType == "premium" {
		if inquiry.Distance > 25.0 {
			finalEstimation = taxedEstimation * 0.95
		} else {
			finalEstimation = taxedEstimation
		}
	} else {
		discount := getNormalDiscount(state, taxedEstimation, inquiry)
		finalEstimation = taxedEstimation - discount
	}

	err = nil

	return
}

// getNormalDiscount gets discount for normal estimation based on business rules
func getNormalDiscount(state string, baseEstimation float32, inquiry entities.InquiryModel) (discount float32) {
	distance := inquiry.Distance
	amount := inquiry.BaseAmount

	switch state {
	case "TX", "OH":
		if distance >= 20.0 && distance <= 30.0 {
			discount = amount * 0.03
		} else if distance > 30.0 {
			discount = baseEstimation * 0.05
		}
	case "CA", "AZ":
		if distance > 26.0 {
			discount = baseEstimation * 0.05
		}
	case "NY":
		discount = 0.0
	default:
		discount = 0.0
	}

	return
}

// getTax calculates tax accordingly
func getTax(state string, baseEstimation float32) (tax float32) {

	if state == "NY" {
		tax = baseEstimation * 0.21
	} else {
		tax = 0.0
	}
	return
}
