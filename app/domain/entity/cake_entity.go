package entity

import (
	"github.com/go-playground/validator/v10"
	formater "github.com/vandyahmad24/validator-formater"
	"time"
)

type CakeInput struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Rating      float32   `json:"rating" validate:"required"`
	Image       string    `json:"image" validate:"required"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func ValidateInputCake(input CakeInput) interface{} {
	var errors interface{}
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		errors = formater.FormatErrorValidation(err, "You must complete input")
	}
	return errors
}
