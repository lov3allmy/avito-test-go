package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/lov3allmy/avito-test-go/internal/domain"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateP2PInput(input domain.P2PInput) []*ErrorResponse {
	validate := validator.New()
	var errors []*ErrorResponse
	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func ValidateGetBalanceInput(input domain.GetBalanceInput) []*ErrorResponse {
	validate := validator.New()
	var errors []*ErrorResponse
	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func ValidateBalanceOperationInput(input domain.BalanceOperationInput) []*ErrorResponse {
	validate := validator.New()
	var errors []*ErrorResponse
	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
