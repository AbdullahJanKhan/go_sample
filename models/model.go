package models

import (
	"fmt"

	validator "github.com/go-playground/validator/v10"
)

// Every request struct needs to implement this interface

type RequestValidation interface {
	Validate() error
}

var validate *validator.Validate

type SampleDataModle struct {
	Key string `json:"key"`
}

func (data *SampleDataModle) Validate() error {
	if data.Key == "" {
		return fmt.Errorf("Error %v", 1001)
	}
	return nil
}
