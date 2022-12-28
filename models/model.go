package models

import (
	"fmt"
)

type SampleDataModle struct {
	Key string `json:"key"`
}

func (data *SampleDataModle) Validate() error {
	if data.Key == "" {
		return fmt.Errorf("Error %v", 1001)
	}
	return nil
}
