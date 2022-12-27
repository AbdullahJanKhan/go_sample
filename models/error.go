package models

import (
	"fmt"
)

type StandardError struct {
	Code        uint
	ActualError error
	Line        string
	Message     string
}

func (s StandardError) Error() string {
	errStr := fmt.Sprintf("Code : %v Line:%v \n Error:%v \n Message: %v", s.Code, s.Line, s.ActualError, s.Message)
	return errStr
}

type APILimiterDto struct {
	UserIp string
	Api    string
	Tries  int64
}
