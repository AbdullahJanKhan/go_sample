# Models

Define your data types and validation if any on them here.
Using the golang validator we can implement a validate function with each data modle we define and use it to perform server side validations on incomming requests

## Sample Code

```go

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
	Key dataType `tag:"tag_key"`
}

func (data *SampleDataModle) Validate() error {
	if data.Key == "" {
		return fmt.Errorf("Error %v", 1001)
	}
	return nil
}

```

## Constants

In the file named consts state all the constants you need. Divide the file in to multiple files as per your requirements.

## Error

I have included a standard error in models. It wraps the traditional error to a custom error.
