package models

import (
	"fmt"
	"regexp"
)

type SampleRequest struct {
	Name     string `json:"name" binding:"required"`
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// binding:"required" enforces on biding time that the value must exist

func (req *SampleRequest) Validate() error {
	if !isEmailValid(req.Email) {
		return StandardError{
			Code:        INVALID_EMAIL_FORMAT,
			ActualError: fmt.Errorf("invalid email"),
			Line:        "Sample Route Modles: Validate():18",
			Message:     INVALID_EMAIL_FORMAT_MSG,
		}
	}
	return nil
}

// isEmailValid checks if the email provided is valid by regex.
func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}
