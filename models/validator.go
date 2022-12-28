package models

import (
	validator "github.com/go-playground/validator/v10"
)

// Every request struct needs to implement this interface

type RequestValidation interface {
	Validate() error
}

var validate *validator.Validate
